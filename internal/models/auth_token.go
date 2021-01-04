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
	ID        uuid.UUID `gorm:"type:uuid;primary_key;" fake:"{uuid}"`
	UserID    uuid.UUID `fake:"{uuid}"`
	LoginWay  string    `fake:"randomstring:[email]"`
	Revoked   uint8     `fake:"{number:0,1}"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (m AuthToken) TableName() string {
	return "auth_tokens"
}
