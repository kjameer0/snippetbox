package main

import (
	"log/slog"
	"net/http"
	"runtime/debug"
)


// this helper method on the app struct creates a serverError
func (app *application) serverError(w http.ResponseWriter, r *http.Request, err error) {
	var (
		method = r.Method
		uri    = r.URL.RequestURI()
		trace  = string(debug.Stack())
	)
	// send error to configured logger with key value paies
	app.logger.Error(err.Error(), "method", method, "uri", uri, slog.String("trace", trace))

	//get correct http error status code
	http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
}

func (app *application) clientError(w http.ResponseWriter, status int) {
	http.Error(w, http.StatusText(status), status)
}
