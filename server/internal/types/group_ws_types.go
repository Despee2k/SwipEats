package types

import (
	"github.com/gorilla/websocket"
)

type Client struct {
	ID      uint
	Conn    *websocket.Conn
	IsFinished bool // Indicates if the client has finished their swiping session
}

type GroupSession struct {
	ID      string
	Clients map[uint]*Client
}

type GroupSessionService struct {
	Sessions map[string]*GroupSession
}

type GroupSessionMessage string

var (
	GroupSessionStartMessage GroupSessionMessage = "start"
	GroupSessionEndMessage   GroupSessionMessage = "end"
	GroupSessionLeaveMessage  GroupSessionMessage = "leave"
	GroupSessionSubmitSwipes GroupSessionMessage = "submit_swipes"
)