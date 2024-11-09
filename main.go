// npx nodemon --exec go run main.go --signal SIGTERM
package main

import (
	"log"
	"net/http"
)

func home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello from Snippetbox"))
}
func main() {
	//mux is the part of the app that guides requests
	//to the url that matches their path
	http.HandleFunc("/", home)
	err := http.ListenAndServe(":4000", nil)
	log.Fatal(err)
}
