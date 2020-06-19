package models

import (
	"github.com/satori/uuid"
	"log"
	"time"
)

type User struct {
	CustomModel

	ID        uuid.UUID `gorm:"type:uuid;primary_key;default:uuid_generate_v4()"`
	Name      string    `gorm:"type:string;not null"`
	Email     string    `gorm:"type:string;size:128;not null"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time `gorm:index`
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


