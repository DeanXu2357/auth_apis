package pkg

import "github.com/jinzhu/gorm"

type SeederFactory struct {
	definitions map[string]func()
	db *gorm.DB
}

type seeder struct {
	data map[string]interface{}
	db *gorm.DB
	model interface{}
}

var seedFactory *SeederFactory

func NewSeederFactory(db *gorm.DB) *SeederFactory {
	once.Do(func() {
		seedFactory = &SeederFactory{db: db}
	})
	return seedFactory
}

func (f *SeederFactory)define(alias string, c func()) {
	f.definitions[alias] = c
}

func Seed(alias string, number int) (res []*seeder) {
	for i:=number;i>0;i-- {
		res = append(res, &seeder{db: seedFactory.db})
	}
	return
}

func (s *seeder)Create(custom map[string]interface{}) {
	//
}
