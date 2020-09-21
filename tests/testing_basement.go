// this  package provides testing tools in this project
package tests

import (
	a "auth/app"
	"bytes"
	"fmt"
	"log"
	"os/exec"
	"path"
	"runtime"
)

var app *a.Instance

// InitialTestingApplication create Instance type for testing environment
// the application *Instance use {root}/app/app.test.yml as configuration.
func InitialTestingApplication() *a.Instance {
	a.SetConfigName("app")
	a.SetAbsolutePath()
	app = a.NewStatic()
	return app
}

// RefreshDatabase
func RefreshDatabase(app *a.Instance) {
	cmd, args := prepareCommandString(app.Configs.Database)

	cmdDown := exec.Command(cmd, append(args, "down", "-all")...)
	fmt.Printf("cmdDown: %s\n", cmdDown.String())
	var out bytes.Buffer
	var stderr bytes.Buffer
	cmdDown.Stdout = &out
	cmdDown.Stderr = &stderr
	err := cmdDown.Run()
	if cmdDown.Run() != nil {
		fmt.Printf("out:%s\nerr:%s\n", out.String(), stderr.String())
		log.Fatal(err)
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

	migrationsPath := fmt.Sprintf("\"%s/../db/migrations/\"", nowPath)

	return command, []string{"-database" ,connection, "-verbose", "-path", migrationsPath}
}