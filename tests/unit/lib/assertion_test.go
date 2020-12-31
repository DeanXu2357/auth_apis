package lib

import (
	"auth/config"
	"auth/lib/assertion"
	"auth/lib/database"
	"auth/models"
	"auth/tests"
	"testing"
)

func Test_DatabaseHas(t *testing.T) {
	config.InitialConfigurations()
	tests.RefreshDatabase()
	db := database.InitialDatabase()

	mockT := new(testing.T)

	if assertion.DatabaseHas(mockT, &models.User{}, map[string]string{}, db) {
		t.Error("should return false")
	}

	// todo: success case
	// todo: not found case and assert if get additional information
}
