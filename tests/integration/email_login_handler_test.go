package integration

import (
	"auth/internal"
	"auth/internal/config"
	"auth/internal/models"
	"auth/internal/routes"
	"auth/lib/assertion"
	"auth/lib/database"
	"auth/lib/event_listener"
	"auth/tests"
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"log"
	"net/http"
	"testing"
)

func Test_Health(t *testing.T) {
	config.InitialConfigurations()
	tests.RefreshDatabase()
	db := database.NewDBEngine()
	dispatcher := event_listener.NewDispatcher()
	router := routes.InitRouter(application.Application{DB: db, Dispatcher: dispatcher})

	w := tests.Call(router, "GET", "/api/v1/health", "")

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Contains(t, w.Body.String(), "ok")
}

func Test_RegisterByEmailSuccess(t *testing.T) {
	config.InitialConfigurations()
	tests.RefreshDatabase()
	db := database.NewDBEngine()
	dispatcher := event_listener.NewDispatcher()
	router := routes.InitRouter(application.Application{DB: db, Dispatcher: dispatcher})

	name := "poyu"
	email := "dean.dh@gmail.com"
	password := "password"

	body, _ := json.Marshal(map[string]interface{}{"name": name, "email": email, "password": password})
	w := tests.Call(router, "POST", "/api/v1/email/register", string(body))

	log.Printf("response body : " + w.Body.String())
	assert.Contains(t, w.Body.String(), "ok")
	assertion.DatabaseHas(t, &models.User{}, map[string]string{"name": name}, db)
	assertion.DatabaseHas(t, &models.EmailLogin{}, map[string]string{"email": email}, db)
}

// todo: assert handler be triggered
// todo: test email has already registered
// todo: test validation failed when register
