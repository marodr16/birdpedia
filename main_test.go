package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHandler(t *testing.T) {
	// create new http request that will be passed to handler
	req, err := http.NewRequest("GET", "", nil)

	if err != nil {
		t.Fatal(err)
	}

	//create http recorder
	recorder := httptest.NewRecorder()

	//create an http handler
	hf := http.HandlerFunc(handler)

	// execute handler
	hf.ServeHTTP(recorder, req)

	// check the status
	if status := recorder.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	// check response body
	expected := `Hello World!`
	actual := recorder.Body.String()
	if actual != expected {
		t.Errorf("handler returned unexpected body: got %v want %v", actual, expected)
	}
}
