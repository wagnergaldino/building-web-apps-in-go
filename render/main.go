package main

import (
	"net/http"
	"gopkg.in/unrolled/render.v1"
)

func main() {
	r := render.New(render.Options{})
	mux := http.NewServeMux()
	
	mux.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		w.Write([]byte("Welcomw, visit subpages now"))
	})
	
	mux.HandleFunc("/data", func(w http.ResponseWriter, req *http.Request) {
		r.Data(w, http.StatusOK, []byte("Some binary data here"))
	})
	
	mux.HandleFunc("/json", func(w http.ResponseWriter, req *http.Request) {
		r.JSON(w, http.StatusOK, map[string]string{"hello": "json",})
	})
	
	mux.HandleFunc("/html", func(w http.ResponseWriter, req *http.Request) {
		r.HTML(w, http.StatusOK, "example", nil)
	})
	
	http.ListenAndServe(":8080", mux)
}
