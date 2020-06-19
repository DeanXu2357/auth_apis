package user_service

import (
	a "auth/app"
	c "auth/configs"
	"fmt"
	"github.com/jinzhu/gorm"
	"github.com/smartystreets/assertions"
	"github.com/stretchr/testify/assert"
	"log"
	"testing"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func TestNew(t *testing.T) {
	app, err := initApplication()
	if err != nil {
		assert.Fail(t, err.Error())
	}
	actual := New(app)

	assertions.ShouldHaveSameTypeAs(actual, "UserService")
}

func TestUserService_Create(t *testing.T) {
}

func TestUserService_GetUserByUUID(t *testing.T) {
}

func initApplication() (application *a.Instance, err error) {
	c.InitConfig()
	config := c.GetConfigs()


	dbInfo := fmt.Sprintf(
		"host=%s port=%s user=%s dbname=%s sslmode=disable password=%s",
		config.Database.DBHost,
		config.Database.DBPort,
		config.Database.DBUser,
		config.Database.DBName,
		config.Database.DBPassword)
	db, err := gorm.Open("postgres", dbInfo)
	if err != nil {
		log.Printf("Database Connection failed : %s", err)
		return
	}
	defer db.Close()

	application = a.New(config, db)

	return
}