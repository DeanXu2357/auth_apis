package integration

import (
	"auth/tests"
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"log"
	"net/http"
	"testing"
)

func Test_Health(t *testing.T) {
	router := tests.PrepareServer()

	w := tests.Call(router, "GET", "/api/v1/health", "")

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Contains(t, w.Body.String(), "OK")
}

func Test_RegisterByEmailSuccess(t *testing.T) {
	router := tests.PrepareServer()

	name := "poyu"
	email := "dean.dh@gmail.com"
	password := "password"

	body, _ := json.Marshal(map[string]interface{}{"name": name, "email": email, "password": password})
	w := tests.Call(router, "POST", "/api/v1/email/register", string(body))

	log.Printf("response body : " + w.Body.String())
	assert.Contains(t, w.Body.String(), "success")
}
