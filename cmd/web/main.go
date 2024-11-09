// npx nodemon --exec go run main.go --signal SIGTERM
package main

import (
	"log"
	"net/http"
)

func main() {
	//mux is the part of the app that guides requests
	//to the url that matches their path
	mux := http.NewServeMux()
	mux.HandleFunc("GET /{$}", home)
	mux.HandleFunc("GET /snippet/view/{id}", snippetView)
	mux.HandleFunc("GET /snippet/create", snippetCreate)
	mux.HandleFunc("POST /snippet/create", snippetCreatePost)

	log.Print("starting server on :4000")

	err := http.ListenAndServe(":4000", mux)
	log.Fatal(err)
}
