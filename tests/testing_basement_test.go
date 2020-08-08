package tests

import (
	"github.com/smartystreets/assertions"
	"testing"
)

func Test_InitInstance(t *testing.T) {
	app := InitialTestingApplication()
	assertions.ShouldHaveSameTypeAs(app, "Instance")
}

func Test_UseAbsolutePath(t *testing.T) {
	app := InitialTestingApplication()
	assertions.ShouldNotBeNil(app.Configs.Server.Port)
}
