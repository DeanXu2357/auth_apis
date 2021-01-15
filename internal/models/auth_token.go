package models

import (
	"github.com/satori/uuid"
	"gorm.io/gorm"
	"reflect"
	"time"
)

const (
	RevokedFalse bool = false
	RevokedTrue  bool = true

	LoginByEmail   = "email"
	LoginByRefresh = "refresh"
)

type AuthToken struct {
	ID        uuid.UUID `gorm:"type:uuid;primary_key;" fake:"{uuid}"`
	UserID    uuid.UUID `fake:"{uuid}"`
	LoginWay  string    `fake:"randomstring:[email]"`
	Revoked   bool      `fake:"{boolean}"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (m AuthToken) TableName() string {
	return "auth_tokens"
}

func (m *AuthToken) BeforeCreate(tx *gorm.DB) (err error) {
	if m.ID != reflect.Zero(reflect.TypeOf(m.ID)).Interface() {
		return
	}

	m.ID = uuid.NewV4()
	return
}

func (m *AuthToken) IsRevoked() bool {
	return m.Revoked == true
}

func (m *AuthToken) DoRevoked(db *gorm.DB) error {
	return db.Model(m).Updates(AuthToken{Revoked: RevokedTrue}).Error
}
