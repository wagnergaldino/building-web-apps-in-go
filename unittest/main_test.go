package main

import (
	"testing"
	"net/http"
	"net/http/httptest"
)

func Test_HelloWorld(t *testing.T) {
	req, err := http.NewRequest("GET", "http://localhost:8080", nil)
	if err != nil {
		t.Fatal(err)
	}
	
	res := httptest.NewRecorder()
	HelloWorld(res, req)
	
	exp := "Hello World"
	act := res.Body.String()
	if exp != act {
		t.Fatalf("Expected %s got %s", exp, act)
	}
}
