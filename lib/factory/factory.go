package factory

import (
	"gorm.io/gorm"
	"reflect"
)

func Create(db *gorm.DB, model interface{}, custom map[string]string, number int) []interface{} {
	models := Generate(model, custom, number)

	db.Create(&models)

	return models
}

func Generate(model interface{}, data map[string]string, number int) []interface{} {
	if number < 1 {
		panic("number cannot less than 1")
	}

	res := make([]interface{}, 0, number)
	for i := 0;i <= number;i++ {
		f := fakeOne(model)
		res = append(res, f)
	}

	return res
}

func fakeOne(model interface{}) interface{} {
	modelType := reflect.TypeOf(model).Elem()
	modelPtr := reflect.New(modelType)
	gofakeit.Struct(&modelPtr)

	return modelPtr
}
