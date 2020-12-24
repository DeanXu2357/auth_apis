package lib

import (
	"auth/lib/assertion"
	"auth/lib/config"
	"auth/lib/database"
	"auth/lib/factory"
	"auth/models"
	"auth/tests"
	"github.com/stretchr/testify/assert"
	"reflect"
	"testing"
)

func Test_FactoryGenerate(t *testing.T) {
	name := "dean"
	fakeUsers := factory.Generate(&models.User{}, map[string]interface{}{"Name": name}, 1)

	fakeUser := fakeUsers[0]
	n := reflect.ValueOf(fakeUser).Elem()
	assert.Equal(t, name, n.FieldByName("Name").Interface().(string))
}

func Test_FactoryCreateOne(t *testing.T) {
	config.InitialConfigurations()
	tests.RefreshDatabase()
	db := database.InitialDatabase()

	name := "dean"
	fakeUsers := factory.Create(db, &models.User{}, map[string]interface{}{"Name": name}, 1)

	fakeUser := fakeUsers[0]
	assert.Equal(t, name, reflect.ValueOf(fakeUser).Elem().FieldByName("Name").Interface().(string))
	assertion.DatabaseHas(t, &models.User{}, map[string]string{"name": name}, db)
}

// todo: add Test_FactoryCreateMulti
