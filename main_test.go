package main

import (
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

var a App

func TestMain(m *testing.M) {
	a = App{}
	a.initialize()

	code := m.Run()
	os.Exit(code)
}

func TestApi(t *testing.T) {
	// req, err := http.NewRequest("GET", "127.0.0.1/api/ipinfo/8.8.8.8", nil)
	// // resp, err := http.Get("http://127.0.0.1:7000/api/ipinfo/8.8.8.8")
	// if err != nil {
	// 	log.Fatal(err)
	// }

	req, _ := http.NewRequest("GET", "/api/ipinfo/8.8.8.8", nil)
	response := executeRequest(req)

	checkResponseCode(t, http.StatusOK, response.Code)
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
