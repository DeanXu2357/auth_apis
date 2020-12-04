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
	"path"
	"runtime"
	"strings"
)

// RefreshDatabase
func RefreshDatabase() {
	cmd, args := prepareCommandString()

	cmdDown := exec.Command(cmd, append(args, "down", "-all")...)
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
	nowPath := ""
	_, filename, _, ok := runtime.Caller(1)
	if ok {
		nowPath = path.Dir(filename)
	}

	command := fmt.Sprintf("%s/../cmd/migrate.linux-amd64", nowPath)

	connection := fmt.Sprintf(
		"\"postgres://%s:%s@%s:%s/%s?sslmode=disable\"",
		viper.Get("dbhost"),
		viper.Get("dbport"),
		viper.Get("user"),
		viper.Get("dbname"),
		viper.Get("dbpassword"))

	migrationsPath := fmt.Sprintf("\"%s/../db/migrations/\"", nowPath)

	return command, []string{"-database" ,connection, "-verbose", "-path", migrationsPath}
}
