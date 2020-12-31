package models

import (
	"github.com/satori/uuid"
	"time"
)

const (
	LoginByEmail = "email"

	RevokedFalse = iota
	RevokedTrue
)

type AuthToken struct {
	ID        uuid.UUID
	UserID    uuid.UUID
	LoginWay  string
	Revoked   uint8
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (m AuthToken) TableName() string {
	return "auth_tokens"
}
