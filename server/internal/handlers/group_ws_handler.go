package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/SwipEats/SwipEats/server/internal/constants"
	"github.com/SwipEats/SwipEats/server/internal/dtos"
	"github.com/SwipEats/SwipEats/server/internal/errors"
	"github.com/SwipEats/SwipEats/server/internal/middlewares"
	"github.com/SwipEats/SwipEats/server/internal/repositories"
	"github.com/SwipEats/SwipEats/server/internal/services"
	"github.com/SwipEats/SwipEats/server/internal/types"
	"github.com/SwipEats/SwipEats/server/internal/utils"
	"github.com/gorilla/websocket"
)

// Helper Function to handle sending member updates
func handleSendMemberUpdate(groupCode string, userID uint, session *types.GroupSession, conn *websocket.Conn, groupStatus types.GroupStatusEnum, deleteMember bool) {
	var groupRestaurants []dtos.GroupRestaurantResponseDto = nil

	// Get group members
	members, err := services.GetGroupMembers(groupCode, userID)
	if err != nil {
		conn.WriteJSON(map[string]string{"error": err.Error()})
		return
	}

	// If the connection is closed, we delete the member
	if deleteMember {
		for i, v := range members {
			if v.UserID == userID {
				members = append(members[:i], members[i+1:]...)
				break
			}
		}
	}

	// If the group status is not waiting, we fetch the group restaurants
	if groupStatus != types.GroupStatusWaiting {
		groupRestaurants, err = services.GetGroupRestaurantsByGroupCode(groupCode)
		if err != nil {
			conn.WriteJSON(map[string]string{"error": err.Error()})
			return
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


// MakeGroupWsHandler creates a WebSocket handler for group sessions
func MakeGroupWsHandler(gss *types.GroupSessionService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var errorResponse dtos.APIErrorResponse
		encoder := json.NewEncoder(w)

		userID := r.Context().Value(middlewares.UserIDKey).(uint)
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

		// Check if the user is already in the group
		_, err = services.JoinGroup(groupCode, userID)
		if err != nil && err != errors.ErrUserAlreadyInGroup {
			errorResponse.Message = "Failed to join group: " + err.Error()
			if err := client.Conn.WriteJSON(errorResponse); err != nil {
				log.Printf("Failed to send error response to user %d: %v", userID, err)
			}
			conn.Close()
			return
		}

		// Defer cleanup when the connection is closed
		defer func() {
			status := getGroupStatus(groupCode)
			if status == nil {
				log.Printf("Failed to get group status for group %s", groupCode)
				return
			}

			log.Printf("User %d disconnected from group %s", userID, groupCode)
			delete(session.Clients, userID)
			handleSendMemberUpdate(groupCode, userID, session, conn, *status, true)
			services.LeaveGroup(userID, groupCode)
			conn.Close()
		}()

		// Send the initial group status and members to the user
		status := getGroupStatus(groupCode)
		if status == nil {
			log.Printf("Failed to get group status for group %s", groupCode)
			return
		}
		handleSendMemberUpdate(groupCode, userID, session, conn, *status, false)

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

					*status = types.GroupStatusActive

					// Broadcast the group session start message to all clients
					services.GroupBroadcast(*session, map[string]interface{}{
						"message":         "Group session started",
						"type":            "group_session_started",
						"group_status": *status,
						"group_code":     groupCode,
						"group_restaurants": groupRestaurants,
					})

				// Handle group session end messages
				case types.GroupSessionEndMessage:
					err := services.EndGroupSession(groupCode, userID)
					if err != nil {
						conn.WriteJSON(map[string]string{"error": err.Error()})
						break
					}

					*status = types.GroupStatusClosed

					// Broadcast the group session end message to all clients
					services.GroupBroadcast(*session, map[string]interface{}{
						"message":     "Group session ended",
						"type":        "group_session_ended",
						"group_status": *status,
						"group_code": groupCode,
						// Add Matched Restaurant here as well as statistics
						// match: Restaurant,
						// statistics: map[string]interface{}{"votes": 10, "comments": 5},
					})
			}
		}
	}
}