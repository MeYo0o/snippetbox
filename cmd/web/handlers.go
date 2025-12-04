package main

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"github.com/MeYo0o/snippetbox/internal/config"
	"github.com/MeYo0o/snippetbox/internal/models"
)

func home(app *config.Application) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Server", "Go")

		snippets, err := app.Snippets.Latest()
		if err != nil {
			serverError(app, w, r, err)
			return
		}

		//* initialize the TemplateData with Current Year
		data := newTemplateData()
		//* update the "Snippets" field
		data.Snippets = snippets

		render(app, w, r, http.StatusOK, "home.tmpl", data)

	}

}

func snippetView(app *config.Application) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id, err := strconv.Atoi(r.PathValue("id"))
		if err != nil || id < 1 {
			clientError(w, http.StatusNotFound)
			return
		}

		snippet, err := app.Snippets.Get(id)
		if err != nil {
			if errors.Is(err, models.ErrNoRecord) {
				clientError(w, http.StatusNotFound)

			} else {
				serverError(app, w, r, err)
			}
			return
		}

		//* initialize the TemplateData with Current Year
		data := newTemplateData()
		//* update the "Snippet" field
		data.Snippet = snippet

		render(app, w, r, http.StatusOK, "view.tmpl", data)
	}
}

func snippetCreate(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Display a form for creating a new snippet..."))
}

func snippetCreatePost(app *config.Application) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		title := "O snail"
		content := "O snail\nClimb Mount Fuji,\nBut slowly, slowly!\n\n– Kobayashi Issa"
		expires := 7

		id, err := app.Snippets.Insert(title, content, expires)
		if err != nil {
			serverError(app, w, r, err)
			return
		}

		http.Redirect(w, r, fmt.Sprintf("/snippet/view/%d", id), http.StatusSeeOther)
	}
}
