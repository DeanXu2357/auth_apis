package services_test

import (
	a "auth/app"
	"auth/services/user_service"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/smartystreets/assertions"
	"testing"
)

func Test_InitUserService(t *testing.T) {
	app := a.New()
	actual := user_service.New(app)
	assertions.ShouldHaveSameTypeAs(actual, "UserService")
}

func Test_UserService_Create(t *testing.T) {
}

func Test_UserService_GetUserByUUID(t *testing.T) {
}

