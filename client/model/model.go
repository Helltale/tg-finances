package model

type AuthState int

const (
	StateStart AuthState = iota
	StateAwaitingName
	StateAwaitingGroup
)

type SessionContext struct {
	State          AuthState
	UserID         int64
	ProfileName    string
	LastMessageIDs []int64
}
