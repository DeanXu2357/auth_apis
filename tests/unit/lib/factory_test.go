package lib

import (
	"auth/lib"
	"auth/lib/asseration"
	"auth/lib/factory"
	"auth/models"
	"auth/tests"
	"github.com/smartystreets/assertions"
	"reflect"
	"testing"
)

func Test_FactoryGenerate(t *testing.T) {
	name := "dean"
	fakeUsers := factory.Generate(&models.User{}, map[string]interface{}{"Name":name}, 1)

	fakeUser := fakeUsers[0]
	n := reflect.ValueOf(fakeUser).Elem()
	assertions.ShouldEqual(n.FieldByName("Name").Interface().(string), name)
	assertions.ShouldHaveSameTypeAs(fakeUsers, []models.User{})
}

func Test_FactoryCreateOne(t *testing.T) {
	lib.InitialConfigurations()
	tests.RefreshDatabase()
	db := lib.InitialDatabase()

	name := "dean"
	fakeUsers := factory.Create(db, &models.User{}, map[string]interface{}{"Name":name}, 1)

	fakeUser := fakeUsers[0]
	assertions.ShouldEqual(reflect.ValueOf(fakeUser).Elem().FieldByName("Name").Interface().(string), name)
	asseration.DatabaseHas(t, &models.User{}, map[string]string{"name":name}, db)
}

// todo: add Test_FactoryCreateMulti
