package lib

import (
	"auth/lib"
	"auth/lib/asseration"
	"auth/models"
	"auth/tests"
	"testing"
)

func Test_DatabaseHas(t *testing.T) {
	lib.InitialConfigurations()
	tests.RefreshDatabase()
	db := lib.InitialDatabase()

	mockT := new(testing.T)

	if asseration.DatabaseHas(mockT, &models.User{}, map[string]string{},db) {
		t.Error("should return false")
	}

	// todo: success case
	// todo: not found case and assert if get additional information
}
