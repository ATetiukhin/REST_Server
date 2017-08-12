package main_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"."
)

var a main.App

func executeRequest(req *http.Request) *httptest.ResponseRecorder {
	rr := httptest.NewRecorder()
	a.Initialize()
	a.Router.ServeHTTP(rr, req)
	return rr
}

func checkResponseCode(t *testing.T, expected, actual int) {
	if expected != actual {
		t.Errorf("Expected response code %d. Got %d\n", expected, actual)
	}
}

func TestStartPage(t *testing.T) {
	req, _ := http.NewRequest("GET", "/", nil)
	response := executeRequest(req)

	checkResponseCode(t, http.StatusOK, response.Code)

	if body := response.Body.String(); body == "[]" {
		t.Errorf("Expected don't an empty array. Got %s", body)
	}
}

func TestDevice(t *testing.T) {
	req, _ := http.NewRequest("GET", "/devices/eth0", nil)
	response := executeRequest(req)

	checkResponseCode(t, http.StatusOK, response.Code)
}

func TestPathError(t *testing.T) {
	req, _ := http.NewRequest("GET", "/devices/", nil)
	response := executeRequest(req)

	checkResponseCode(t, http.StatusNotFound, response.Code)
}

func TestDevicePageError(t *testing.T) {
	req, _ := http.NewRequest("GET", "/devices/1", nil)
	response := executeRequest(req)

	checkResponseCode(t, http.StatusInternalServerError, response.Code)
}
