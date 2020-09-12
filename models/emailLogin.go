package models

import (
	"time"
)

type EmailLogin struct {
	Email string `gorm:"primary_key;"`
	Password string
	CreatedAt time.Time
	UpdatedAt time.Time
	VerifiedAt time.Time
}

func (EmailLogin) TableName() string {
	return "email_login"
}
