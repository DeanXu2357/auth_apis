package integration

import (
	"auth/lib"
	"auth/routes"
	"auth/tests"
)

func Test_RegisterSuccess(t *testing.T) {
	tests.RefreshDatabase()
	lib.InitialConfigurations()
	router := routes.InitRouter()

	name := "poyu"
	email := "dean.dh@gmail.com"
	password := "password"

}
