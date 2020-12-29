package factory

import (
	"github.com/brianvoe/gofakeit/v5"
	"github.com/fatih/structs"
	"gorm.io/gorm"
	"log"
	"reflect"
	//m "auth/models"
)

// Create generate fake data raw and insert into database
func Create(db *gorm.DB, model interface{}, custom map[string]interface{}, number int) []interface{} {
	models := Generate(model, custom, number)

	for _, m := range models {
		m2 := structs.Map(m)
		log.Println(m2)
		db.Model(model).Create(m2)
	}
	// TODO: refactor this
	// because of using orm batch insert , hooks won’t be invoked
	// use log insert for temporary

	return models
}

// Generate only generate fake data raw , but no insert
func Generate(model interface{}, data map[string]interface{}, number int) []interface{} {
	if number < 1 {
		panic("number cannot less than 1")
	}

	res := make([]interface{}, 0, number)
	for i := 0;i < number;i++ {
		gofakeit.Struct(model)
		if len(data) != 0 {
			for name, value := range data {
				setFakeField(model, name, value)
			}
		}

		res = append(res, reflect.ValueOf(model).Interface())
	}

	return res
}

func setFakeField(fakeModel interface{}, fieldName string, value interface{}) {
	v := reflect.ValueOf(fakeModel).Elem().FieldByName(fieldName)
	if v.IsValid() {
		vValue := reflect.New(reflect.TypeOf(value))
		vValue.Elem().Set(reflect.ValueOf(value))
		v.Set(vValue.Elem())
	}
}
