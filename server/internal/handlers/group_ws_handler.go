package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/SwipEats/SwipEats/server/internal/dtos"
	"github.com/SwipEats/SwipEats/server/internal/middlewares"
	"github.com/SwipEats/SwipEats/server/internal/services"
	"github.com/SwipEats/SwipEats/server/internal/types"
	"github.com/SwipEats/SwipEats/server/internal/utils"
)

func MakeGroupWsHandler(gss *types.GroupSessionService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var errorResponse dtos.APIErrorResponse
		encoder := json.NewEncoder(w)

		userID := r.Context().Value(middlewares.UserIDKey).(uint)
		groupCode := r.URL.Query().Get("group_code")

		if groupCode == "" {
			errorResponse.Message = "Group code is required"
			encoder.Encode(&errorResponse)
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		conn, err := utils.Upgrade(w, r, nil)
		if err != nil {
			errorResponse.Message = "Failed to upgrade connection"
			encoder.Encode(&errorResponse)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		defer conn.Close()

		client := &types.Client{
			ID:   userID,
			Conn: conn,
		}

		session := services.GetOrCreateGroupSession(gss, groupCode)
		session.Clients[userID] = client

		for {
			var msg map[string]string

			if err := conn.ReadJSON(&msg); err != nil {
				conn.Close()
				delete(session.Clients, userID)
				return
			}

			switch types.GroupSessionMessage(msg["type"]) {
				case types.GroupSessionStartMessage:
					err := services.StartGroupSession(groupCode, userID)
					if err != nil {
						conn.WriteJSON(map[string]string{"error": err.Error()})
						continue
					}
					services.GroupBroadcast(*session, msg)
				case types.GroupSessionEndMessage:
					err := services.EndGroupSession(groupCode, userID)
					if err != nil {
						conn.WriteJSON(map[string]string{"error": err.Error()})
						continue
					}
					services.GroupBroadcast(*session, msg)
			}
		}
	}
}