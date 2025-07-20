package model

import (
	"time"
)

type Metadata = map[string]any

type Session struct {
	SessionId string    `bson:"sessionId,omitempty"`
	ExpireAt  time.Time `bson:"expireAt,omitempty"`
	UserId    string    `bson:"userId,omitempty"`
	Metadata  Metadata  `bson:"metaData,omitempty"`
}
