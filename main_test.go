package main

import (
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"

	unoHttp "github.com/UnoHouse/api/http"
	"github.com/UnoHouse/api/service"
	"github.com/UnoHouse/api/sql"
	"github.com/UnoHouse/api/utils/logger"
	"github.com/gorilla/mux"
	_ "github.com/joho/godotenv/autoload"
)

const (
	contentJson string = "application/json"
	contentText string = "text/plain"
)

var TestRouter http.Handler

func TestMain(m *testing.M) {
	db := sql.NewDB()
	usersDb := sql.NewUsersDB()

	s := service.NewMySqlService(db)
	us := service.NewMySqlService(usersDb)
	r := mux.NewRouter()
	h := unoHttp.NewHandler(r, s, us)
	TestRouter = unoHttp.NewRouter(h)
	code := m.Run()
	os.Exit(code)
}

func Test_Indexpage(t *testing.T) {
	req, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		logger.LogErr(err)
	}
	response := executeRequest(req)
	CheckResponseCode(t, http.StatusOK, response.Code)
	CheckResponseType(t, contentJson, response.Header())
	expected := `{
        "path": "/",
        "methods": [
            "GET"
        ]
    },`
	CheckExpectedResponseBody(t, expected, response.Body.String())
}

func Test_Health(t *testing.T) {
	req, err := http.NewRequest("GET", "/health", nil)
	if err != nil {
		logger.LogErr(err)
	}
	response := executeRequest(req)
	CheckResponseCode(t, http.StatusOK, response.Code)
	expected := "Up"
	CheckExpectedResponseBody(t, expected, response.Body.String())
}

func Test_NotFoundResponse(t *testing.T) {
	req, err := http.NewRequest("GET", "/route-that-not-exists", nil)
	if err != nil {
		logger.LogErr(err)
	}
	response := executeRequest(req)
	CheckResponseCode(t, http.StatusNotFound, response.Code)
}

func executeRequest(req *http.Request) *httptest.ResponseRecorder {

	rw := httptest.NewRecorder()
	TestRouter.ServeHTTP(rw, req)
	return rw
}

func CheckResponseCode(t *testing.T, expected, actual int) {
	if expected != actual {
		t.Errorf("Expected response code %d. Got %d\n", expected, actual)
	}
}

func CheckResponseType(t *testing.T, expected string, h http.Header) {
	if strings.Contains(h.Get("Content-Type"), expected) == false {
		t.Errorf("Expected response header Content-Type %s. Got %s\n", expected, h.Get("Content-Type"))
	}
}

func CheckExpectedResponseBody(t *testing.T, expected string, b string) {
	if strings.Contains(b, expected) == false {
		t.Errorf("Expected response  %s. Got %s\n", expected, b)
	}
}
