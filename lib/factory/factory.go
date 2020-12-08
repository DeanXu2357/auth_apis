package factory

import "gorm.io/gorm"

func Create(db *gorm.DB, model interface{}, custom map[string]string, number int) []interface{} {
	models := Generate(model, custom, number)

	db.Create(&models)

	return models
}

func Generate(model interface{}, custom map[string]string, number int) []interface{} {
	//
}
