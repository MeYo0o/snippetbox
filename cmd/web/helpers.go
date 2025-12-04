package main

import (
	"bytes"
	"fmt"
	"net/http"

	"github.com/MeYo0o/snippetbox/internal/config"
)

func serverError(app *config.Application, w http.ResponseWriter, r *http.Request, err error) {
	var (
		method = r.Method
		uri    = r.URL.RequestURI()
	)

	app.Logger.Error(err.Error(), "method", method, "uri", uri)
	http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
}

func clientError(w http.ResponseWriter, statusCode int) {
	http.Error(w, http.StatusText(statusCode), statusCode)
}

func render(app *config.Application, w http.ResponseWriter, r *http.Request, status int, page string, data templateData) {
	ts, ok := app.TemplateCache[page]
	if !ok {
		err := fmt.Errorf("the template %s does not exist", page)
		serverError(app, w, r, err)
		return
	}

	buffer := new(bytes.Buffer)

	err := ts.ExecuteTemplate(buffer, "base", data)
	if err != nil {
		serverError(app, w, r, err)
	}

	w.WriteHeader(status)

	buffer.WriteTo(w)

}
