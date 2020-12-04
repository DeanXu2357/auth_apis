package models

import (
	"errors"
	"gorm.io/gorm"
)

type CustomModel struct {
	db *gorm.DB
}

func (m *CustomModel)SetConnection(db *gorm.DB) {
	m.db = db
}

func (m *CustomModel)GetDB() (db *gorm.DB, err error) {
	if m.db == nil {
		err = errors.New("db connection not set yet")
		return
	}
	db = m.db
	return
}
