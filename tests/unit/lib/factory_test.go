package lib

import (
	"auth/lib"
	"auth/lib/factory"
	"auth/models"
	"auth/tests"
	"github.com/smartystreets/assertions"
	"reflect"
	"testing"
)

func Test_FactoryGenerate(t *testing.T) {
	name := "dean"
	fakeUsers := factory.Generate(&models.User{}, map[string]interface{}{"name":name}, 1)

	fakeUser := fakeUsers[0]
	n := reflect.ValueOf(fakeUser).Elem()
	assertions.ShouldEqual(n.FieldByName("name").Interface().(string), name)

}

func Test_FactoryCreate(t *testing.T) {
	lib.InitialConfigurations()
	tests.RefreshDatabase()
	db := lib.InitialDatabase()

	name := "dean"
	fakeUsers := factory.Create(db, &models.User{}, map[string]interface{}{"name":name}, 1)

	fakeUser := fakeUsers[0]
	n := reflect.ValueOf(fakeUser).Elem()
	assertions.ShouldEqual(n.FieldByName("name").Interface().(string), name)
}
