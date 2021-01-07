package handler_tests

import (
	"auth/internal/config"
	"auth/internal/models"
	"auth/internal/services"
	"auth/lib/assertion"
	"auth/lib/database"
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

	user, err := services.Register(name, email, password, db)

	assert.Nil(t, err)
	assert.IsType(t, &models.User{}, user)
	assertion.DatabaseHas(t, &models.User{}, map[string]string{"name": name}, db)
	assertion.DatabaseHas(t, &models.EmailLogin{}, map[string]string{"email": email}, db)
}

// todo: test email has already registered
