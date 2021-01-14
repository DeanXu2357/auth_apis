package integration

import (
	"auth/internal/config"
	"auth/internal/models"
	"auth/internal/routes"
	"auth/internal/services"
	"auth/lib/database"
	"auth/lib/event_listener"
	"auth/lib/factory"
	"auth/tests"
	"encoding/json"
	"fmt"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

// refresh success
func Test_RefreshLoginTokenSuccess(t *testing.T) {
	// Setup & Initial
	config.InitialConfigurations()
	tests.RefreshDatabase()
	db := database.InitialDatabase()
	router := routes.InitRouter(db, event_listener.NewDispatcher())

	// Arrange
	users := factory.Create(
		db.Session(&gorm.Session{NewDB: true}),
		&models.User{},
		map[string]interface{}{},
		1,
	)
	user := users[0].(*models.User)
	tokenString, err := services.GenerateLoginToken(*user, db.Session(&gorm.Session{NewDB: true}), "for testing")
	if err != nil {
		t.Error(err.Error())
	}

	// Act
	req, _ := http.NewRequest("POST", "/api/v1/refresh", strings.NewReader(""))
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", tokenString))
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	// Assert
	assert.Equal(t, http.StatusOK, w.Code)
	assert.Contains(t, w.Body.String(), "ok")
	var rsp struct{
		Items struct{Token string `json:"token"`} `json:"items"`
		Msg string `json:"msg"`
		Status int `json:"status"`
	}
	if err := json.Unmarshal(w.Body.Bytes(), &rsp) ;err != nil{
		t.Error(err.Error())
	}
	assert.NotNil(t, rsp)
	at, err := services.DecodeLoginToken(rsp.Items.Token, db.Session(&gorm.Session{NewDB: true}))
	if err != nil {
		t.Error(err.Error())
	}
	assert.NotNil(t, at)
}

// refresh failed, due to refresh limit
