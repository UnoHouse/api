package main

import (
	"net/http"
	"strings"
	"testing"

	"github.com/UnoHouse/api/utils/logger"
	_ "github.com/joho/godotenv/autoload"
)

func Test_getUserById(t *testing.T) {
	req, err := http.NewRequest("GET", "/users/1", nil)
	if err != nil {
		logger.LogErr(err)
	}
	response := executeRequest(req)
	CheckResponseCode(t, http.StatusOK, response.Code)
}

func Test_getUserByIdResponse(t *testing.T) {
	req, err := http.NewRequest("GET", "/users/1", nil)
	if err != nil {
		logger.LogErr(err)
	}
	response := executeRequest(req)
	CheckResponseType(t, contentJson, response.Header())
	expected := `{"data":{"id":1,"username":`
	CheckExpectedResponseBody(t, expected, response.Body.String())
}
func Test_getUserWithNonExistentId(t *testing.T) {
	req, err := http.NewRequest("GET", "/users/1235654", nil)
	if err != nil {
		logger.LogErr(err)
	}

	response := executeRequest(req)
	CheckResponseCode(t, http.StatusNotFound, response.Code)
}

func Test_getUserWithMissingId(t *testing.T) {
	req, err := http.NewRequest("GET", "/users", nil)
	if err != nil {
		logger.LogErr(err)
	}

	response := executeRequest(req)
	CheckResponseCode(t, http.StatusNotFound, response.Code)
}

func Test_getUserByStringId(t *testing.T) {
	req, err := http.NewRequest("GET", "/users/asd", nil)
	if err != nil {
		logger.LogErr(err)
	}

	response := executeRequest(req)
	CheckResponseCode(t, http.StatusBadRequest, response.Code)
}
func Test_postUserByIdMissingBody(t *testing.T) {
	req, err := http.NewRequest("POST", "/users/1", nil)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	if err != nil {
		logger.LogErr(err)
	}
	response := executeRequest(req)
	CheckResponseCode(t, http.StatusBadRequest, response.Code)
}

func Test_postUserById(t *testing.T) {
	req, err := http.NewRequest("POST", "/users/1", strings.NewReader("firstname=test&lastname=test2"))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	if err != nil {
		logger.LogErr(err)
	}
	response := executeRequest(req)
	CheckResponseCode(t, http.StatusOK, response.Code)
	expected := `{"data":{"id":1,"username":"user1","firstname":"test","lastname":"test2"`
	CheckExpectedResponseBody(t, expected, response.Body.String())
}

func Test_postUserByStringId(t *testing.T) {
	req, err := http.NewRequest("POST", "/users/asd", strings.NewReader("firstname=test&lastname=test2"))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	if err != nil {
		logger.LogErr(err)
	}
	response := executeRequest(req)
	CheckResponseCode(t, http.StatusBadRequest, response.Code)
}
