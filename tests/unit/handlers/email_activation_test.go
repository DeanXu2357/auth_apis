package handler_tests

import (
	"auth/internal/config"
	handlers_v1 "auth/internal/handlers/v1"
	"auth/internal/models"
	"auth/lib/assertion"
	"auth/lib/database"
	"auth/lib/factory"
	"auth/tests"
	"encoding/hex"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
	"testing"
)

var (
	testPrivateKey = "-----BEGIN RSA PRIVATE KEY-----\n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\n-----END RSA PRIVATE KEY-----\n"
	testPublicKey  = "-----BEGIN PUBLIC KEY-----\nMIICIjANBgkqhkiG9w0BAQEFAAOCAg8AMIICCgKCAgEA9QifyUIaWwnhCbl91STz\nur7wyVNGZBHN4aVeALRHdf1RsAKJJhuY9FJwnbiPvyk7z8m7ogVazUmF0UoKYp4z\nPPJMTLQb97oN45kKNFGa0tYQk68dfird7kvFvZ6JohmW96K/iJGTcMdMelq0DH8f\nKug8TCl/RrTOUYJjSJi5UNcnePpJmvG28Ts/3FL/mGMEQnrrUI3rMdVYd0J9uB/X\n+nvHHqBxWPRLq2gz3/zsOdHSRbp/SVwRlS8123G8KdsSwLlAqyFwYMXBY/lsClPL\nhSfRjFL2GACsFmw22UG8Y8dRD4DNlesNj9bWRW1I2lIO0JCH2sbGreXbQbmwjVoW\n49HGEo9daG38yFUMgM84RsCNoQo2c+BXz7vAx72AizV2nZxeaXuV2KqAlzr3kwoE\n3bI+NcwuuY6d1CQeR9i/xSwxHG+lrIPkOhwjZWcAvpyeBs+WBvhv8N3tw9KF+J+j\nV2Ls2ivFPW3Vt1gWH4pjSfgxFDxOyyFIiR1vG/z9JmCZf825D69vnNdei9u40RcK\njkuWYknJtbv04vTvrvylo1APPXvBuq+nUsImaaCj17A1wSIR1DR0zjubHM7KeY56\nxGEEl9ltqiVXsLQhfb0OcQFYuMRR0C9OQ1G+CQSrBhq+OubQJHRlAYZ3eOOKT7mR\n0dGDIMQWH6+dqCUuzuMgk9MCAwEAAQ==\n-----END PUBLIC KEY-----\n"
)

func Test_GenerateActionToken(t *testing.T) {
	// Setup
	config.InitialConfigurations()
	tests.RefreshDatabase()

	// Initial
	db := database.InitialDatabase()
	config.ActivateAuth = config.ActivateAuthSettings{PrivateKey: testPrivateKey, PublicKey: testPublicKey, Expire: 360}

	// Arrange
	fakeUsers := factory.Create(db, &models.User{}, map[string]interface{}{}, 1)
	fakeUser := fakeUsers[0]
	user, ok := fakeUser.(*models.User)
	if !ok {
		t.Error("Factory Error")
	}

	// Act
	newSession := db.Session(&gorm.Session{NewDB: true})
	tokenString, err := handlers_v1.GenerateActivationToken(user, newSession)

	// Assert
	assert.Nil(t, err)
	assert.NotNil(t, tokenString)
	assertion.DatabaseHas(
		t,
		&models.EmailVerify{},
		map[string]string{"email": user.Email, "user_id": hex.EncodeToString(user.ID.Bytes())},
		db,
	)
}

func Test_ActivateSuccess(t *testing.T) {
	// Setup
	config.InitialConfigurations()
	tests.RefreshDatabase()

	// Initial
	db := database.InitialDatabase()
	config.ActivateAuth = config.ActivateAuthSettings{PrivateKey: testPrivateKey, PublicKey: testPublicKey, Expire: 360}

	// Arrange
	fakeUsers := factory.Create(db, &models.User{}, map[string]interface{}{}, 1)
	fakeUser := fakeUsers[0]
	user, ok := fakeUser.(*models.User)
	_ = factory.Create(db, &models.EmailLogin{}, map[string]interface{}{"Email": user.Email}, 1)
	if !ok {
		t.Error("Factory Error")
	}

	// Act
	tokenString, gErr := handlers_v1.GenerateActivationToken(user, db.Session(&gorm.Session{NewDB: true}))
	aErr := handlers_v1.Activate(tokenString, db.Session(&gorm.Session{NewDB: true}))

	assert.Nil(t, gErr)
	assert.Nil(t, aErr)
	assertion.DatabaseHas(
		t,
		&models.EmailVerify{},
		map[string]interface{}{
			"email":        user.Email,
			"user_id":      hex.EncodeToString(user.ID.Bytes()),
			"verification": models.VerifyTrue},
		db,
	)
}
