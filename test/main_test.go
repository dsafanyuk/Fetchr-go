package test

import (
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/dsafanyuk/fetchr-go/app"
	"github.com/dsafanyuk/fetchr-go/config"
)

var a app.App

func TestMain(m *testing.M) {
	config := config.GetConfig()
	a = app.App{}
	a.Initialize(config)

	ensureUserTableExists()
	ensureProductTableExists()

	code := m.Run()

	clearProductTable()
	clearUserTable()

	os.Exit(code)
}

func executeRequest(req *http.Request) *httptest.ResponseRecorder {
	rr := httptest.NewRecorder()
	a.Router.ServeHTTP(rr, req)

	return rr
}

func checkResponseCode(t *testing.T, expected, actual int) {
	if expected != actual {
		t.Errorf("Expected response code %d. Got %d\n", expected, actual)
	}
}
