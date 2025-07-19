package model

import (
	"time"
)

type Metadata = map[string]any

type Session struct {
	SessionID  string
	ExpireTime time.Duration
	SessionValue
}

type SessionValue struct {
	UserID   string
	Metadata Metadata
}
