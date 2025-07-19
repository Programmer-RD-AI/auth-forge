package session

import "github.com/Programmer-RD-AI/auth-forge/internal/model"

type SessionRepository interface {
	GetByUserId() *model.Session
	GetBySessionId() *model.SessionValue
	DeleteSessionId() bool
}
