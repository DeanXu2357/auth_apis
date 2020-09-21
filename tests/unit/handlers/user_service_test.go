package handlers_test

import (
	"auth/models"
	"auth/services/user_service"
	"auth/tests"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/smartystreets/assertions"
	"testing"
)

func Test_InitUserService(t *testing.T) {
	app := tests.InitialTestingApplication()
	actual := user_service.New(app)
	assertions.ShouldHaveSameTypeAs(actual, "UserService")
}

func Test_UserService_Create(t *testing.T) {
	app := tests.InitialTestingApplication()

	tests.RefreshDatabase(app)

	service := user_service.New(app)
	userName := "poyu"
	userEmail := "poyu@example.com"

	user, err := service.Create(map[string]interface{}{"name": userName, "email": userEmail})
	assertions.ShouldBeNil(err)

	var result models.User
	app.Database.Raw("select * from users where email = ?", userEmail).Scan(&result)
	assertions.ShouldNotBeNil(user)
	assertions.ShouldNotBeNil(result)
	assertions.ShouldEqual(userName, result.Name)
}

func Test_UserService_GetUserByUUID(t *testing.T) {
}

