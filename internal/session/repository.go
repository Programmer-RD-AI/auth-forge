package session

import "github.com/Programmer-RD-AI/auth-forge/internal/store"


type SessionRepository struct {
	store.RedisStore
}

func CreateSession() {}

func RefreshSession() {}
