// npx nodemon --exec go run ./cmd/web --signal SIGTERM -e go
package main

import (
	"log"
	"net/http"
)

func main() {
	//mux is the part of the app that guides requests
	//to the url that matches their path
	mux := http.NewServeMux()
	// the file server with assets comes from a specific folder
	fileServer := http.FileServer(http.Dir("./ui/static/"))
	// StripPrefix gets rid of /static from the URL so we aren't searching
	// for /static/static/path-to-asset
	// create a get route for all assets
	mux.Handle("GET /static/", http.StripPrefix("/static", neuter(fileServer)))

	mux.HandleFunc("GET /{$}", home)
	mux.HandleFunc("GET /snippet/view/{id}", snippetView)
	mux.HandleFunc("GET /snippet/create", snippetCreate)
	mux.HandleFunc("POST /snippet/create", snippetCreatePost)

	log.Print("starting server on :4000")

	err := http.ListenAndServe(":4000", mux)
	log.Fatal(err)
}
