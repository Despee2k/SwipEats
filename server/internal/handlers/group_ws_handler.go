package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	// "time"

	"github.com/SwipEats/SwipEats/server/internal/constants"
	"github.com/SwipEats/SwipEats/server/internal/dtos"
	"github.com/SwipEats/SwipEats/server/internal/errors"
	"github.com/SwipEats/SwipEats/server/internal/repositories"
	"github.com/SwipEats/SwipEats/server/internal/services"
	"github.com/SwipEats/SwipEats/server/internal/types"
	"github.com/SwipEats/SwipEats/server/internal/utils"
	"github.com/gorilla/websocket"
)

// MakeGroupWsHandler creates a WebSocket handler for group sessions
func MakeGroupWsHandler(gss *types.GroupSessionService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var errorResponse dtos.APIErrorResponse
		var status *types.GroupStatusEnum
		encoder := json.NewEncoder(w)


		// Get token
		token := r.URL.Query().Get("token")
		if token == "" {
			errorResponse.Message = "Token is required"
			w.WriteHeader(http.StatusBadRequest)
			encoder.Encode(&errorResponse)
			return
		}

		userID, err := getUserIDFromToken(token)
		if err != nil {
			errorResponse.Message = "Unauthorized: Invalid token"
			w.WriteHeader(http.StatusUnauthorized)
			encoder.Encode(&errorResponse)
			return
		}
		
		groupCode := r.URL.Query().Get("group_code")

		if groupCode == "" {
			errorResponse.Message = "Group code is required"
			w.WriteHeader(http.StatusBadRequest)
			encoder.Encode(&errorResponse)
			return
		}

		// Upgrade the HTTP connection to a WebSocket connection
		conn, err := utils.Upgrade(w, r, nil)
		if err != nil {
			errorResponse.Message = "Failed to upgrade connection"
			w.WriteHeader(http.StatusInternalServerError)
			encoder.Encode(&errorResponse)
			return
		}

		client := &types.Client{
			ID:   userID,
			Conn: conn,
			IsFinished: false, // Initially, the client has not finished their swiping session
		}

		// Get or create the group session
		session := services.GetOrCreateGroupSession(gss, groupCode)
		session.Clients[userID] = client

		// Check if the user is already in the group
		_, err = services.JoinGroup(groupCode, userID)
		if err != nil && err != errors.ErrUserAlreadyInGroup {
			errorResponse.Message = "Failed to join group: " + err.Error()
			if err := client.Conn.WriteJSON(errorResponse); err != nil {
				log.Printf("Failed to send error response to user %d: %v", userID, err)
			}
			conn.Close()
			return
		} else {
			status = getGroupStatus(groupCode)
			if status == nil {
				log.Printf("Failed to get group status for group %s", groupCode)
				return
			}

			if err == nil {
				// Check if user has already voted
				group, err := repositories.GetGroupByCode(groupCode)
				if err != nil {
					log.Printf("Failed to get group by code %s: %v", groupCode, err)
					return 
				}
				count, err := repositories.GetSwipeCountByUserAndGroup(userID, group.ID)
				if err != nil {
					log.Printf("Failed to get swipe count for user %d in group %d: %v", userID, group.ID, err)
					return
				}

				if count > 0 {
					client.IsFinished = true
				}

				handleSendMemberUpdate(groupCode, userID, session, conn, *status, false)
			}
		}



//
//
// 		MAIN LOGIC
// 		Handle incoming messages from the WebSocket connection
//
//
		for {
			var msg map[string]interface{}

			// Read the message from the WebSocket connection
			if err := conn.ReadJSON(&msg); err != nil {
				break
			}

			// Check message type and handle accordingly
			switch types.GroupSessionMessage(msg["type"].(string)) {


				// START
				// Handle group session start messages
				case types.GroupSessionStartMessage:
					err := services.StartGroupSession(groupCode, userID)
					if err != nil {
						conn.WriteJSON(map[string]string{"error": err.Error()})
						break
					}
				
					// Generate group restaurants and send them to all clients
					groupRestaurants, err := services.GenerateGroupRestaurants(groupCode, constants.SEARCH_RADIUS, constants.MAX_NUM_OF_RESTAURANTS)
					if err != nil {
						conn.WriteJSON(map[string]string{"error": err.Error()})
						break
					}

					*status = types.GroupStatusActive

					// Broadcast the group session start message to all clients
					services.GroupBroadcast(*session, map[string]any{
						"message":         "Group session started",
						"type":            "group_session_started",
						"group_status": 	types.GroupStatusActive,
						"group_code":     	groupCode,
						"group_restaurants": groupRestaurants,
					})



				// END
				// Handle group session end messages
				case types.GroupSessionEndMessage:
					if success := endGroupSession(groupCode, userID, conn, session, status); !success {
						conn.WriteJSON(map[string]string{"error": "Failed to end group session"})
					}
					return



				// LEAVE
				// Handle group session leave messages
				case types.GroupSessionLeaveMessage:
					handleSendMemberUpdate(groupCode, userID, session, conn, *status, true)
					return



				// SUBMIT SWIPES
				// Handle group session submit swipes messages
				case types.GroupSessionSubmitSwipes:
					
					votesRaw, ok := msg["votes"].(map[string]interface{})
					if !ok {
						log.Println("votes field is not a valid map[string]interface{}")
						return
					}

					for groupRestaurantIDString, voteRaw := range votesRaw {
						groupRestaurantID, err := strconv.ParseUint(groupRestaurantIDString, 10, 32)
						if err != nil {
							log.Printf("Invalid restaurant ID: %s", groupRestaurantIDString)
							continue
						}

						// Convert vote to string, then compare
						vote, ok := voteRaw.(bool)
						if !ok {
							log.Println("Vote value is not a string:", voteRaw)
							continue
						}

						// Add Swipe for this group restaurant
						services.AddSwipe(dtos.AddSwipeDto{
							IsLiked:          vote,
							GroupRestaurantID: uint(groupRestaurantID),
						}, userID)
					}

					client.IsFinished = true // Mark the client as finished swiping
					client.Conn.WriteJSON(map[string]interface{}{
						"type":            "group_session_swipe_finished",
					})

					// Check if all clients have finished swiping
					allFinished := true
					for _, c := range session.Clients {
						if !c.IsFinished {
							log.Printf("Client %d has not finished swiping", c.ID)
							allFinished = false
							break
						}
					}

					// If finished, end the group session
					if allFinished {
						endGroupSession(groupCode, userID, conn, session, status)
					}
			}
		}
	}
}







///
/// HELPER FUNCTIONS
///
func getUserIDFromToken(token string) (uint, error) {
	user, err := utils.ValidateJWT(token)
	if err != nil {
		return 0, err
	}

	existingUser, err := repositories.GetUserByEmail(user.Email)
	if err != nil {
		return 0, err
	}

	return existingUser.ID, nil
}

// Helper Function to handle sending member updates
func handleSendMemberUpdate(groupCode string, userID uint, session *types.GroupSession, conn *websocket.Conn, groupStatus types.GroupStatusEnum, isLeave bool) {
	var groupRestaurants []dtos.GroupRestaurantResponseDto = nil

	// Get group members
	members, err := services.GetGroupMembers(groupCode, userID)
	if err != nil {
		conn.WriteJSON(map[string]string{"error": err.Error()})
		return
	}

	if groupStatus != types.GroupStatusWaiting { // If the group status is not waiting, we fetch the group restaurants
		groupRestaurants, err = services.GetGroupRestaurantsByGroupCode(groupCode)
		if err != nil {
			conn.WriteJSON(map[string]string{"error": err.Error()})
			return
		}
	}

	if isLeave {
		// If the user is leaving, we need to remove them from the members list
		for i, member := range members {
			if member.UserID == userID {
				members = append(members[:i], members[i+1:]...)
				services.LeaveGroup(userID, groupCode)
				break
			}
		}
	}

	response := map[string]interface{}{
		"message":     "Group members updated",
		"type":       "members_update",
		"members":     members,
		"group_status": groupStatus,
		"group_code":  groupCode,
	}

	if groupRestaurants != nil {
		response["group_restaurants"] = groupRestaurants
	}

	services.GroupBroadcast(*session, response)
}

// Helper function to get the group status
func getGroupStatus(groupCode string) (*types.GroupStatusEnum) {
	group, err := repositories.GetGroupByCode(groupCode)
	if err != nil || group == nil {
		return nil
	}
	return &group.GroupStatus
}

func endGroupSession(groupCode string, userID uint, conn *websocket.Conn, session *types.GroupSession, status *types.GroupStatusEnum) bool {
	err := services.EndGroupSession(groupCode, userID)
	if err != nil {
		conn.WriteJSON(map[string]string{"error": err.Error()})
		return false
	}

	if status != nil {
		// Update the group status to closed
		*status = types.GroupStatusClosed
	}

	// Get the group's most liked group restaurant
	mostLikedGroupRestaurant, err := services.GetMostLikedGroupRestaurant(groupCode)
	if err != nil {
		conn.WriteJSON(map[string]string{"error": err.Error()})
		return false
	}

	response := map[string]interface{}{
		"message":     "Group session ended",
		"type":        "group_session_ended",
		"group_status": types.GroupStatusClosed,
		"group_code": groupCode,
	}
	
	if mostLikedGroupRestaurant != nil {
		// Save as a match
		id, err := services.SaveMostLikedGroupRestaurant(mostLikedGroupRestaurant.ID)
		if err != nil {
			conn.WriteJSON(map[string]string{"error": err.Error()})
			return false
		}

		match, err := repositories.GetMatchByID(id)
		if err != nil {
			conn.WriteJSON(map[string]string{"error": err.Error()})
			return false
		}

		// Add the most liked group restaurant to the response
		response["most_liked_group_restaurant"] = dtos.GroupRestaurantResponseDto{
			ID:        match.ID,
			GroupID:  match.GroupID,
			Restaurant: match.Restaurant,
			DistanceInKM: utils.DistanceInKM(
				match.Group.LocationLat,
				match.Group.LocationLong,
				match.Restaurant.LocationLat,
				match.Restaurant.LocationLong,
			),
		}
	}

	// Broadcast the group session end message to all clients
	services.GroupBroadcast(*session, response)

	return true
}