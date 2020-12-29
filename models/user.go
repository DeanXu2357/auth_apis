package models

import (
	"github.com/satori/uuid"
	"gorm.io/gorm"
	"reflect"
	"time"
)

type CustomModel interface {
	TableName() string
}

type User struct {
	ID        uuid.UUID `gorm:"type:uuid;primary_key;" fake:"{uuid}"`
	Name      string    `gorm:"type:string;not null" fake:"{name}"`
	Email     string    `gorm:"type:string;size:128;not null" fake:"{email}"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time
}

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	if u.ID != reflect.Zero(reflect.TypeOf(u.ID)).Interface() {
		return
	}

	u.ID = uuid.NewV4()
	return
}

func (u User) TableName() string {
	return "users"
}
