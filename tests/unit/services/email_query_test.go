package service_tests

import (
	"auth/internal/config"
	"auth/internal/models"
	"auth/internal/services"
	"auth/lib/database"
	"auth/lib/factory"
	"auth/tests"
	"errors"
	"github.com/brianvoe/gofakeit/v5"
	"github.com/stretchr/testify/assert"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"testing"
	"time"
)

func Test_LoginNotFound(t *testing.T) {
	// Setup
	config.InitialConfigurations()
	tests.RefreshDatabase()

	// Initial
	db := database.InitialDatabase()

	// act
	_, err := services.FindEmailLogin("test@gmail.com", db.Session(&gorm.Session{NewDB: true}))

	// assert
	assert.NotNil(t, err)
	assert.True(t, errors.Is(err, services.ErrLoginNotExist))
}

// test not verify yet
func Test_LoginNotVerifiedYet(t *testing.T) {
	// Setup
	config.InitialConfigurations()
	tests.RefreshDatabase()

	// Initial
	db := database.InitialDatabase()

	// arrange
	u, l, _ := createFakeUserWithEmailLogin(
		map[string]interface{}{},
		map[string]interface{}{},
		db.Session(&gorm.Session{NewDB: true}),
	)

	// act
	actual, err := services.FindEmailLogin(u.Email, db.Session(&gorm.Session{NewDB: true}))

	// assert
	assert.NotNil(t, err)
	assert.True(t, errors.Is(err, services.ErrEmailNotVerify))
	assert.Equal(t, l.Email, actual.Email)
	assert.Equal(t, l.Password, actual.Password)
}

// test get
func Test_GetLoginSuccess(t *testing.T) {
	// Setup
	config.InitialConfigurations()
	tests.RefreshDatabase()

	// Initial
	db := database.InitialDatabase()

	// arrange
	u, l, _ := createFakeUserWithEmailLogin(
		map[string]interface{}{},
		map[string]interface{}{"VerifiedAt":time.Now()},
		db.Session(&gorm.Session{NewDB: true}),
	)

	// act
	actual, err := services.FindEmailLogin(u.Email, db.Session(&gorm.Session{NewDB: true}))

	// assert
	assert.Nil(t, err)
	assert.Equal(t, l.Email, actual.Email)
	assert.Equal(t, l.Password, actual.Password)
}

func createFakeUserWithEmailLogin(
	userCustom map[string]interface{},
	emailLoginCustom map[string]interface{},
	db *gorm.DB,
) (*models.User, *models.EmailLogin, string) {
	fakeUsers := factory.Create(db, &models.User{}, userCustom, 1)
	fakeUser := fakeUsers[0]

	pwd := gofakeit.Color()
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(pwd), 8)

	emailLoginCustom["Password"] = string(hashedPassword)
	emailLoginCustom["Email"] = fakeUser.(*models.User).Email

	fakeLogins := factory.Create(db, &models.EmailLogin{}, emailLoginCustom, 1)
	fakeLogin := fakeLogins[0]

	return fakeUser.(*models.User), fakeLogin.(*models.EmailLogin), pwd
}
