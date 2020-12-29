package models

import (
	"github.com/satori/uuid"
	"gorm.io/gorm"
	"reflect"
	"time"
)

const (
	MailTypeVerifyAccount = "verify"
	MailTypeResetPassword = "reset"

	VerifyFalse int8 = iota
	VerifyTrue
)

type EmailVerify struct {
	ID           uuid.UUID `gorm:"type:uuid;primary_key;" fake:"{uuid}"`
	Email        string    `gorm:"type:string;size:128;not null" fake:"{email}"`
	MailType     string    `gorm:"type:string;size:64:not null" fake:"{randomstring:[verify,reset]}"`
	Verification int8      `gorm:"type:int;"`
	UserID       uuid.UUID `gorm:"type:uuid" fake:"{uuid}"`
	CreatedAt    time.Time
	UpdatedAt    time.Time
	User         User `gorm:"foreignKey:UserID" fake:"skip"`
}

func (u *EmailVerify) BeforeCreate(tx *gorm.DB) (err error) {
	if u.ID != reflect.Zero(reflect.TypeOf(u.ID)).Interface() {
		return
	}

	u.ID = uuid.NewV4()
	return
}

func (u EmailVerify) TableName() string {
	return "email_verify"
}
