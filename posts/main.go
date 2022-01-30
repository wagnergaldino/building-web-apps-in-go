package main

import (
	"net/http"
	"fmt"
	"github.com/julienschmidt/httprouter"
)

func main() {
	
	r := httprouter.New()
	r.GET("/", HomeHandler)
	
	//posts collection
	r.GET("/posts", PostsIndexHandler)
	r.POST("/posts", PostsCreateHandler)
	
	//posts singular
	r.GET("/posts/:id", PostShowHandler)
	r.PUT("/posts/:id", PostUpdateHandler)
	r.GET("/posts/:id/edit", PostEditHandler)
	
	fmt.Println("Starting server on :8080")
	http.ListenAndServe(":8080", r)
	
}

func HomeHandler(rw http.ResponseWriter, r *http.Request, p httprouter.Params) {
	fmt.Fprintln(rw, "Home")
}

func PostsIndexHandler(rw http.ResponseWriter, r *http.Request, p httprouter.Params) {
	fmt.Fprintln(rw, "Posts Index")
}

func PostsCreateHandler(rw http.ResponseWriter, r *http.Request, p httprouter.Params) {
	fmt.Fprintln(rw, "Posts Create")
}

func PostShowHandler(rw http.ResponseWriter, r *http.Request, p httprouter.Params) {
	id := p.ByName("id")
	fmt.Fprintln(rw, "Showing Post", id)
}

func PostUpdateHandler(rw http.ResponseWriter, r *http.Request, p httprouter.Params) {
	fmt.Fprintln(rw, "Post Update")
}

func PostDeleteHandler(rw http.ResponseWriter, r *http.Request, p httprouter.Params) {
	fmt.Fprintln(rw, "Post Delete")
}

func PostEditHandler(rw http.ResponseWriter, r *http.Request, p httprouter.Params) {
	fmt.Fprintln(rw, "Post Edit")
}
