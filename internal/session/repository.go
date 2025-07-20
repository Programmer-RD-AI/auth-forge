package session

import (
	"context"

	"github.com/Programmer-RD-AI/auth-forge/internal/model"
)

type SessionRepository interface {
	GetByUserId(ctx context.Context, userId string) (*model.Session, error)
	GetBySessionId(ctx context.Context, sessionID string) (*model.Session, error)
	DeleteSessionId(ctx context.Context, sessionID string) bool
}
