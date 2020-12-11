package factory

import (
	"github.com/brianvoe/gofakeit/v5"
	"gorm.io/gorm"
	"log"
	"reflect"
)

func Create(db *gorm.DB, model interface{}, custom map[string]interface{}, number int) []interface{} {
	models := Generate(model, custom, number)

	db.Create(&models)

	return models
}

func Generate(model interface{}, data map[string]interface{}, number int) []interface{} {
	if number < 1 {
		panic("number cannot less than 1")
	}

	res := make([]interface{}, 0, number)
	for i := 0;i < number;i++ {
		f := fakeOne(model)
		log.Println(f)
		if len(data) != 0 {
			for name, value := range data {
				setFakeField(f, name, value)
			}
		}
		log.Println(f)

		res = append(res, f)
	}

	return res
}

func setFakeField(fakeModel interface{}, fieldName string, value interface{}) {
	v := reflect.ValueOf(fakeModel).Elem().FieldByName(fieldName)
	log.Printf("field name and value : %s, %s\n", fieldName, value)
	if v.IsValid() {
		vValue := reflect.New(reflect.TypeOf(value))
		vValue.Elem().Set(reflect.ValueOf(value))
		v.Set(vValue)
	}
}

func fakeOne(model interface{}) interface{} {
	modelV := reflect.TypeOf(model).Elem()
	modelPtr := reflect.New(modelV)
	gofakeit.Struct(&modelPtr)

	return modelPtr.Interface()
}
