package types

type GroupStatusEnum string

const (
	GroupStatusWaiting GroupStatusEnum = "waiting"
	GroupStatusActive  GroupStatusEnum = "active"
	GroupStatusClosed  GroupStatusEnum = "closed"
)