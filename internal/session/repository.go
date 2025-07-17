package session

import "github.com/Programmer-RD-AI/auth-forge/internal/model"

type SessionRepository struct {
	store model.BaseStore
}

func NewSessionRepository(store model.BaseStore) SessionRepository {
	return SessionRepository{
		store: nil,
	}
}
