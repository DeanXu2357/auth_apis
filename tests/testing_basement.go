// this  package provides testing tools in this project
package tests

import (
	a "auth/app"
	"fmt"
	"log"
	"os/exec"
	"path"
	"runtime"
)

// InitialTestingApplication create Instance type for testing environment
// the application *Instance use {root}/app/app.test.yml as configuration.
func InitialTestingApplication() *a.Instance {
	app := a.New()
	a.SetConfigName("app.test")
	a.SetAbsolutePath()
	configs := a.InitConfigs()
	app.SetConfigs(configs)
	return app
}

// RefreshDatabase
func RefreshDatabase(app *a.Instance) {
	cmd, args := prepareCommandString(app.Configs.Database)

	cmdDown := exec.Command(cmd, append(args, "down")...)
	err := cmdDown.Run()
	if cmdDown.Run() != nil {
		log.Panic(err)
	}

	cmdUp := exec.Command(cmd, append(args, "up")...)
	err = cmdUp.Run()
	if err != nil {
		log.Panic(err)
	}
}

// prepareCommandString
// migration command cheat sheet:
// ./cmd/migrate.linux-amd64 -database "postgres://postgres:fortestpwd@localhost:45487/auth?sslmode=disable" -verbose -path db/migrations up
func prepareCommandString(dbConfig a.DatabaseConfigurations) (string, []string) {
	nowPath := ""
	_, filename, _, ok := runtime.Caller(1)
	if ok {
		nowPath = path.Dir(filename)
	}

	command := fmt.Sprintf("%s/../cmd/migrate.linux-amd64", nowPath)

	connection := fmt.Sprintf(
		"\"postgres://%s:%s@%s:%s/%s?sslmode=disable\"",
		dbConfig.DBUser,
		dbConfig.DBPassword,
		dbConfig.DBHost,
		dbConfig.DBPort,
		dbConfig.DBName)

	migrationsPath := fmt.Sprintf("%s/../db/migrations", nowPath)

	return command, []string{connection, "-verbose", migrationsPath}
}