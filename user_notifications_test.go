package main

import (
	"net/http"
	"testing"

	"github.com/UnoHouse/api/utils/logger"
	_ "github.com/joho/godotenv/autoload"
)

func Test_getUserNotificationsById(t *testing.T) {
	req, err := http.NewRequest("GET", "/users/1/notifications", nil)
	if err != nil {
		logger.LogErr(err)
	}
	response := executeRequest(req)
	CheckResponseCode(t, http.StatusOK, response.Code)
}

func Test_getUserNotificationsByIdResponse(t *testing.T) {
	req, err := http.NewRequest("GET", "/users/1/notifications", nil)
	if err != nil {
		logger.LogErr(err)
	}
	response := executeRequest(req)
	CheckResponseType(t, contentJson, response.Header())
	expected := `{"data":{"items"`
	CheckExpectedResponseBody(t, expected, response.Body.String())
}

func Test_getUserNotificationsByStringId(t *testing.T) {
	req, err := http.NewRequest("GET", "/users/asd/notifications", nil)
	if err != nil {
		logger.LogErr(err)
	}

	response := executeRequest(req)
	CheckResponseCode(t, http.StatusBadRequest, response.Code)
}

func Test_getUserNotificationsByNonExistentId(t *testing.T) {
	req, err := http.NewRequest("GET", "/users/485687/notifications", nil)
	if err != nil {
		logger.LogErr(err)
	}

	response := executeRequest(req)
	CheckResponseCode(t, http.StatusNotFound, response.Code)
}
