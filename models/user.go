package models

import (
	"github.com/jinzhu/gorm"
	"github.com/satori/uuid"
	"time"
)

type User struct {
	ID        uuid.UUID `gorm:"type:uuid;primary_key;"`
	Name      string    `gorm:"type:string;not null"`
	Email     string    `gorm:"type:string;size:128;not null"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time
}

func (u *User)BeforeCreate(scope *gorm.Scope) error {
	return scope.SetColumn("ID", uuid.NewV4().String())
}
