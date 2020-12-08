package models

import (
	"gorm.io/gorm"
	"github.com/satori/uuid"
	"time"
)

type CustomModel interface {
	TableName() string
}

type User struct {
	ID        uuid.UUID `gorm:"type:uuid;primary_key;"`
	Name      string    `gorm:"type:string;not null"`
	Email     string    `gorm:"type:string;size:128;not null"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time
}

func (u *User)BeforeCreate(tx *gorm.DB) (err error) {
	u.ID = uuid.NewV4()
	return
}

func (u User) TableName() string {
	return "users"
}