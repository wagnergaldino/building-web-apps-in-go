package main

import (
	"net/http"
	"encoding/json"
)

type Book struct {
	Title string `json:"title"`
	Author string `json:"author"`
}

func main() {
	http.HandleFunc("/", ShowBooks)
	http.ListenAndServe(":8080", nil)
}

func ShowBooks(w http.ResponseWriter, r *http.Request) {
	book := Book{"Building Web Apps with GO", "Jeremy Saenz"}
	
	js, err := json.Marshal(book)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	
	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}
