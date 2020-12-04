package handler_tests

import (
	handlers_v1 "auth/handlers/v1"
	"auth/tests"
	"github.com/smartystreets/assertions"
	"testing"
)

func Test_RegisterSuccess(t *testing.T) {
	app := tests.InitialTestingApplication()
	tests.RefreshDatabase(app)
	h := handlers_v1.EmailLoginHandler{app}

	name := "poyu"
	email := "dean.dh@gmail.com"
	password := "password"

	user, err := h.Register(name, email, password)

	assertions.ShouldNotBeNil(err)
	assertions.ShouldHaveSameTypeAs(user, "User")
	// todo: assert database has ...
}
