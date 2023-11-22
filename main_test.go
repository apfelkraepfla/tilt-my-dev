package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
)

func TestIndexHandler(t *testing.T) {
	request, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Fatal(err)
	}

	resultRecorder := httptest.NewRecorder()

	router := mux.NewRouter()
	router.HandleFunc("/", indexHandler)
	router.ServeHTTP(resultRecorder, request)

	expectedStatusCode := http.StatusOK
	if recvStatusCode := resultRecorder.Result().StatusCode; recvStatusCode != expectedStatusCode {
		t.Errorf("Handler returned wrong status code: got %v want %v", recvStatusCode, expectedStatusCode)
	}

	expectedText := "<h1>Hello World!!!</h1>"
	if resultRecorder.Body.String() != expectedText {
		t.Errorf("Handler returned unexpected body: got %v want %v", resultRecorder.Body.String(), expectedText)
	}
}
