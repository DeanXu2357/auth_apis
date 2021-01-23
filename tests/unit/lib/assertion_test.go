package lib

import (
	"auth/internal/config"
	"auth/internal/models"
	"auth/lib/assertion"
	"auth/lib/database"
	"auth/tests"
	"testing"
)

func Test_DatabaseHas(t *testing.T) {
	config.InitialConfigurations()
	tests.RefreshDatabase()
	db := database.NewDBEngine()

	mockT := new(testing.T)

	if assertion.DatabaseHas(mockT, &models.User{}, map[string]string{}, db) {
		t.Error("should return false")
	}

	// todo: success case
	// todo: not found case and assert if get additional information
}
