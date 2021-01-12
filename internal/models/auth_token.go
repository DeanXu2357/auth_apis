package models

import (
	"github.com/satori/uuid"
	"time"
)

const (
	RevokedFalse bool = false
	RevokedTrue bool = true

	LoginByEmail = "email"
	LoginByRefresh = "refresh"
)

type AuthToken struct {
	ID        uuid.UUID `gorm:"type:uuid;primary_key;" fake:"{uuid}"`
	UserID    uuid.UUID `fake:"{uuid}"`
	LoginWay  string    `fake:"randomstring:[email]"`
	Revoked   bool     `fake:"{boolean}"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (m AuthToken) TableName() string {
	return "auth_tokens"
}
