package authforge

import (
	"time"

	"github.com/Programmer-RD-AI/auth-forge/internal/session"
)

func CreateOrExpireToken(userId string) session.Session {
	return session.Session{SessionID: "test", ExpireTime: time.Duration(60) * time.Second}
}

func SessionValidation(sessionId string) bool { return true }

func RevokeSession() (bool, error) { return true, nil }
