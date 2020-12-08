package asseration

import (
	"auth/models"
	"fmt"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
	"reflect"
)

type tHelper interface {
	Helper()
}

func DatabaseHas(t assert.TestingT, model interface{}, condition interface{}, db *gorm.DB, msgAndArgs ...interface{}) bool {
	if h, ok := t.(tHelper); ok {
		h.Helper()
	}

	immutable := reflect.ValueOf(model).Elem()
	sliceType := reflect.SliceOf(immutable.Type())
	slice := reflect.New(sliceType).Interface()

	result := db.Where(condition).Find(slice)
	if result.RowsAffected <= 0 {
		return assert.Fail(
			t,
			fmt.Sprintf(
				"Raws could not be found in %s\nRaws Found : \n%s",
				reflect.TypeOf(model).Kind(),
				getAdditionalInfo(model, db)),
			msgAndArgs...)
	}

	return true
}

func getAdditionalInfo(model interface{}, db *gorm.DB) string {
	immutable := reflect.ValueOf(model).Elem()
	sliceType := reflect.SliceOf(immutable.Type())
	slice := reflect.New(sliceType).Interface()

	if h, ok := model.(models.CustomModel); ok {
		db.Table(h.TableName()).Find(slice)
	}

	output := ""
	switch reflect.TypeOf(slice).Kind() {
	case reflect.Slice:
		for _, r := range slice.([]interface{}) {
			output += "\n" + fmt.Sprintln(r)
		}
	}

	return output
}
