package services_test

import (
	a "auth/app"
	"github.com/smartystreets/assertions"
	"testing"
)

func TestInitInstance(t *testing.T) {
	app := InitApplicationInTest()
	assertions.ShouldHaveSameTypeAs(app, "Instance")
}

func TestUseAbsolutePath(t *testing.T) {
	app := InitApplicationInTest()
	assertions.ShouldNotBeNil(app.Configs.Server.Port)
}

func InitApplicationInTest() *a.Instance {
	app := a.New()
	a.SetConfigAbsolutePath()
	configs := a.InitConfigs()
	app.SetConfigs(configs)
	return app
}
