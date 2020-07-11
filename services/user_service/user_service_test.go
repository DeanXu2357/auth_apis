package user_service

import (
	a "auth/app"
	"github.com/smartystreets/assertions"
	"testing"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func TestInitUserService(t *testing.T) {
	app := a.New()
	actual := New(app)
	assertions.ShouldHaveSameTypeAs(actual, "UserService")
}

func TestUserService_Create(t *testing.T) {
}

func TestUserService_GetUserByUUID(t *testing.T) {
}

