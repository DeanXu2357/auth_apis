// this  package provides testing tools in this project
package tests

import (
	"auth/internal/cmd/migration"
	"net/http"
	"net/http/httptest"
	"strings"
)

// RefreshDatabase
func RefreshDatabase() {
	migration.RefreshDatabase()
}

func Call(r http.Handler, method, path, body string) *httptest.ResponseRecorder {
	req, _ := http.NewRequest(method, path, strings.NewReader(body))
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	return w
}
