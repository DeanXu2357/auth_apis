package integration

import (
	"auth/lib"
	"auth/lib/asseration"
	"auth/lib/event_listener"
	"auth/models"
	"auth/routes"
	"auth/tests"
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"log"
	"net/http"
	"testing"
)

func Test_Health(t *testing.T) {
	lib.InitialConfigurations()
	tests.RefreshDatabase()
	db := lib.InitialDatabase()
	dispatcher := event_listener.NewDispatcher()
	router := routes.InitRouter(db, dispatcher)

	w := tests.Call(router, "GET", "/api/v1/health", "")

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Contains(t, w.Body.String(), "OK")
}

func Test_RegisterByEmailSuccess(t *testing.T) {
	lib.InitialConfigurations()
	tests.RefreshDatabase()
	db := lib.InitialDatabase()
	dispatcher := event_listener.NewDispatcher()
	router := routes.InitRouter(db, dispatcher)

	name := "poyu"
	email := "dean.dh@gmail.com"
	password := "password"

	body, _ := json.Marshal(map[string]interface{}{"name": name, "email": email, "password": password})
	w := tests.Call(router, "POST", "/api/v1/email/register", string(body))

	log.Printf("response body : " + w.Body.String())
	assert.Contains(t, w.Body.String(), "success")
	asseration.DatabaseHas(t, &models.User{}, map[string]string{"name":name}, db)
	asseration.DatabaseHas(t, &models.EmailLogin{}, map[string]string{"email":email}, db)
}

// todo: assert handler be triggered
