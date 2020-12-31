package handler_tests

import (
	"auth/config"
	"auth/handlers/v1"
	"auth/lib/assertion"
	"auth/lib/database"
	"auth/models"
	"auth/tests"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_RegisterSuccess(t *testing.T) {
	config.InitialConfigurations()
	tests.RefreshDatabase()
	db := database.InitialDatabase()

	name := "poyu"
	email := "dean.dh@gmail.com"
	password := "password"

	user, err := handlers_v1.Register(name, email, password, db)

	assert.Nil(t, err)
	assert.IsType(t, &models.User{}, user)
	assertion.DatabaseHas(t, &models.User{}, map[string]string{"name": name}, db)
	assertion.DatabaseHas(t, &models.EmailLogin{}, map[string]string{"email": email}, db)
}

// todo: test email has already registered
