package handler_tests

import (
	handlersv1 "auth/handlers/v1"
	"auth/lib"
	"auth/models"
	"auth/tests"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_RegisterSuccess(t *testing.T) {
	lib.InitialConfigurations()
	tests.RefreshDatabase()
	db := lib.InitialDatabase()

	name := "poyu"
	email := "dean.dh@gmail.com"
	password := "password"

	user, err := handlersv1.Register(name, email, password, db)

	assert.Nil(t, err)
	assert.IsType(t, &models.User{}, user)
	// todo: assert database has ...
}
