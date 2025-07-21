package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/SwipEats/SwipEats/server/internal/constants"
	"github.com/SwipEats/SwipEats/server/internal/dtos"
	"github.com/SwipEats/SwipEats/server/internal/errors"
	"github.com/SwipEats/SwipEats/server/internal/repositories"
	"github.com/SwipEats/SwipEats/server/internal/services"
	"github.com/SwipEats/SwipEats/server/internal/types"
	"github.com/SwipEats/SwipEats/server/internal/utils"
	"github.com/gorilla/websocket"
)

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
		if err = services.SaveMostLikedGroupRestaurant(mostLikedGroupRestaurant.ID); err != nil {
			conn.WriteJSON(map[string]string{"error": err.Error()})
			return false
		}

		// Add the most liked group restaurant to the response
		response["most_liked_group_restaurant"] = mostLikedGroupRestaurant
	}

	// Broadcast the group session end message to all clients
	services.GroupBroadcast(*session, response)

	return true
}


// MakeGroupWsHandler creates a WebSocket handler for group sessions
func MakeGroupWsHandler(gss *types.GroupSessionService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var errorResponse dtos.APIErrorResponse
		var status *types.GroupStatusEnum
		encoder := json.NewEncoder(w)

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
		}

		// Get or create the group session
		session := services.GetOrCreateGroupSession(gss, groupCode)
		session.Clients[userID] = client

		// Check group status
		 

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
				handleSendMemberUpdate(groupCode, userID, session, conn, *status, false)
			}
		}

		done := make(chan struct{})

		// Goroutine to check if group is done
		go func() {
			// every 5 seconds
			ticker := time.NewTicker(5 * time.Second)
			defer ticker.Stop()

			for {
				select {
				case <-ticker.C:
					isDone, err := services.CheckIfGroupIsDone(groupCode)
					if err != nil {
						continue
					}
					if isDone {
						success := endGroupSession(groupCode, userID, conn, session, status)
						if !success {
							log.Printf("Failed to end group session for group %s", groupCode)
							conn.WriteJSON(map[string]string{"error": "Failed to end group session"})
						}
						return
					}
					if status != nil {
						handleSendMemberUpdate(groupCode, userID, session, conn, *status, false)
					}
				case <-done:
					return
				}
			}
		}()

		// Defer cleanup when the connection is closed
		defer func() {
			status := getGroupStatus(groupCode)
			if status == nil {
				log.Printf("Failed to get group status for group %s", groupCode)
				return
			}

			close(done)
			delete(session.Clients, userID)
			handleSendMemberUpdate(groupCode, userID, session, conn, *status, false)
			conn.Close()
		}()

		// Handle incoming messages from the WebSocket connection
		for {
			var msg map[string]string

			// Read the message from the WebSocket connection
			if err := conn.ReadJSON(&msg); err != nil {
				break
			}

			// Check message type and handle accordingly
			switch types.GroupSessionMessage(msg["type"]) {
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

					if status != nil {
						*status = types.GroupStatusActive
					}

					// Broadcast the group session start message to all clients
					services.GroupBroadcast(*session, map[string]interface{}{
						"message":         "Group session started",
						"type":            "group_session_started",
						"group_status": 	types.GroupStatusActive,
						"group_code":     	groupCode,
						"group_restaurants": groupRestaurants,
					})

				// Handle group session end messages
				case types.GroupSessionEndMessage:
					if success := endGroupSession(groupCode, userID, conn, session, status); !success {
						conn.WriteJSON(map[string]string{"error": "Failed to end group session"})
					}
					return

				// Handle group session leave messages
				case types.GroupSessionLeaveMessage:
					handleSendMemberUpdate(groupCode, userID, session, conn, *status, true)
					return
			}
		}
	}
}