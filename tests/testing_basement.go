// this  package provides testing tools in this project
package tests

import (
	"bytes"
	"fmt"
	"github.com/spf13/viper"
	"log"
	"net/http"
	"net/http/httptest"
	"os/exec"
	"strings"
)

// RefreshDatabase
// TODO: modify to pipe , prevent concurrent issue
func RefreshDatabase() {
	cmd, args := prepareCommandString()

	cmdDown := exec.Command(cmd, append(args, "down", "-all")...)
	var out bytes.Buffer
	var stderr bytes.Buffer
	cmdDown.Stdout = &out
	cmdDown.Stderr = &stderr
	cmdDown.Dir = "/go/src/app"
	err := cmdDown.Run()
	if err != nil {
		log.Printf("\nout:%s\nerr:%s\n", out.String(), stderr.String())
		log.Print(cmdDown.Args)
		log.Fatal(err)
	}

	cmdUp := exec.Command(cmd, append(args, "up")...)
	cmdUp.Stdout = &out
	cmdUp.Stderr = &stderr
	cmdUp.Dir = "/go/src/app"
	err = cmdUp.Run()
	if err != nil {
		log.Printf("out:%s\nerr:%s\n", out.String(), stderr.String())
		log.Panic(err)
	}
}

func Call(r http.Handler, method, path, body string) *httptest.ResponseRecorder {
	req, _ := http.NewRequest(method, path, strings.NewReader(body))
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	return w
}

// prepareCommandString
// migration command cheat sheet:
// ./cmd/migrate.linux-amd64 -database "postgres://postgres:fortestpwd@localhost:45487/auth?sslmode=disable" -verbose -path db/migrations up
func prepareCommandString() (string, []string) {
	connection := fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s?sslmode=disable",
		viper.GetString("db_user"),
		viper.GetString("db_password"),
		viper.GetString("db_host"),
		viper.GetString("db_port"),
		viper.GetString("db_name"))

	return "./db/migrate.linux-amd64", []string{"-database", connection, "-verbose", "-path", "./db/migrations/"}
}
