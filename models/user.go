package models

import (
	"github.com/jinzhu/gorm"
	"github.com/satori/uuid"
	"log"
	"time"
)

type User struct {
	CustomModel

	ID        uuid.UUID `gorm:"type:uuid;primary_key;"`
	Name      string    `gorm:"type:string;not null"`
	Email     string    `gorm:"type:string;size:128;not null"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time `gorm:index`
}

func (u *User)BeforeCreate(scope *gorm.Scope) error {
	return scope.SetColumn("ID", uuid.NewV4().String())
}

func (u *User)Create() (err error) {
	db, err := u.GetDB()
	if err != nil {
		return
	}

	if err = db.Create(&u).Error;err != nil {
		log.Panic("Unable to create user.")

		return
	}

	return
}


