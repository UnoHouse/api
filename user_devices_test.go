package main

import (
	"net/http"
	"strings"
	"testing"

	"github.com/UnoHouse/api/utils/logger"
	_ "github.com/joho/godotenv/autoload"
)

func Test_getUserDevicesById(t *testing.T) {
	req, err := http.NewRequest("GET", "/users/1/devices", nil)
	if err != nil {
		logger.LogErr(err)
	}
	response := executeRequest(req)
	CheckResponseCode(t, http.StatusOK, response.Code)
}

func Test_getUserDevicesByIdResponse(t *testing.T) {
	req, err := http.NewRequest("GET", "/users/1/devices", nil)
	if err != nil {
		logger.LogErr(err)
	}
	response := executeRequest(req)
	CheckResponseType(t, contentJson, response.Header())
	expected := `{"data":{"items"`
	CheckExpectedResponseBody(t, expected, response.Body.String())
}

func Test_getUserDevicesByStringId(t *testing.T) {
	req, err := http.NewRequest("GET", "/users/asd/devices", nil)
	if err != nil {
		logger.LogErr(err)
	}

	response := executeRequest(req)
	CheckResponseCode(t, http.StatusBadRequest, response.Code)
}

func Test_getUserDevicesByNonExistentId(t *testing.T) {
	req, err := http.NewRequest("GET", "/users/485687/devices", nil)
	if err != nil {
		logger.LogErr(err)
	}

	response := executeRequest(req)
	CheckResponseCode(t, http.StatusNotFound, response.Code)
}

func Test_postUserDevicesByIdMissingBody(t *testing.T) {
	req, err := http.NewRequest("POST", "/users/1/devices", nil)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	if err != nil {
		logger.LogErr(err)
	}
	response := executeRequest(req)
	CheckResponseCode(t, http.StatusBadRequest, response.Code)
}

func Test_postUserDevicesById(t *testing.T) {
	req, err := http.NewRequest("POST", "/users/1/devices", strings.NewReader("firebaseToken=token&osType=2&sdkVersion=29&model=s20&brand=Samsung"))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	if err != nil {
		logger.LogErr(err)
	}
	response := executeRequest(req)
	CheckResponseCode(t, http.StatusOK, response.Code)
	expected := `{"data":{"id"`
	CheckExpectedResponseBody(t, expected, response.Body.String())
}

func Test_postUserDevicesByStringId(t *testing.T) {
	req, err := http.NewRequest("POST", "/users/asd/devices", strings.NewReader("firstname=test&lastname=test2"))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	if err != nil {
		logger.LogErr(err)
	}
	response := executeRequest(req)
	CheckResponseCode(t, http.StatusBadRequest, response.Code)
}
