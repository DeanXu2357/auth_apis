package handler_tests

import (
	"auth/handlers/v1"
	"auth/lib"
	"auth/lib/asseration"
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

	user, err := handlers_v1.Register(name, email, password, db)

	assert.Nil(t, err)
	assert.IsType(t, &models.User{}, user)
	asseration.DatabaseHas(t, &models.User{}, map[string]string{"name": name}, db)
	asseration.DatabaseHas(t, &models.EmailLogin{}, map[string]string{"email": email}, db)
}
