package main

import "net/http"

func (conf *Config) serverError(w http.ResponseWriter, r *http.Request, err error) {
	var (
		method = r.Method
		uri    = r.URL.RequestURI()
	)

	conf.Logger.Error(err.Error(), "method", method, "uri", uri)
	http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
}

func (conf *Config) clientError(w http.ResponseWriter, status int) {
	http.Error(w, http.StatusText(status), status)
}
