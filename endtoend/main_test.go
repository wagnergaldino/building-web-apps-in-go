package main

import (
	"testing"
	"net/http"
	"net/http/httptest"
	"io/ioutil"
)

func Test_App(t *testing.T) {
	ts := httptest.NewServer(App())
	defer ts.Close()
	
	res, err := http.GET(ts.URL)
	if err != nil {
		t.Fatal(err)
	}
	
	body, err := ioutil.ReadAll(res.Body)
	res.Body.Close()
	
	if err != nil {
		t.Fatalf(err)
	}
	
	exp := "b4Hello Worldafter"

	if exp != string(body) {
		t.Fatalf("Expected %s got %s", exp, body)
	}
}
