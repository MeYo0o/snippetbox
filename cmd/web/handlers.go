package main

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"snippetbox.innolabs.ai/internal/models"
)

func (conf *Config) homeHandler(w http.ResponseWriter, r *http.Request) {
	//* priority: first
	w.Header().Add("Server", "Go")
	//* priority: second
	// w.WriteHeader(http.StatusOK)
	//* priority: third
	// w.Write([]byte("Hello from Snippetbox"))

	// err := components.Home().Render(r.Context(), w)
	// if err != nil {
	// 	conf.serverError(w, r, err)
	// }

	snippets, err := conf.Snippets.Latest()
	if err != nil {
		conf.serverError(w, r, err)
		return
	}

	for _, snippet := range snippets {
		fmt.Fprintf(w, "%+v\n", snippet)
	}
}

func (conf *Config) snippetView(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		http.NotFound(w, r)
		return
	}

	snippet, err := conf.Snippets.Get(int32(id))
	if err != nil {
		if errors.Is(err, models.ErrNoRecord) {
			http.NotFound(w, r)
		} else {
			conf.serverError(w, r, err)
		}
		return
	}

	fmt.Fprintf(w, "%+v", snippet)
}

func (conf *Config) snippetCreate(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Display a form for creating a new snippet..."))
}

func (conf *Config) snipperCreatePost(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusCreated)

	title := "O snail"
	content := "O snail\nClimb Mount Fuji,\nBut slowly, slowly!\n\n– Kobayashi Issa"
	expires := 7

	id, err := conf.Snippets.Insert(title, content, expires)
	if err != nil {
		conf.serverError(w, r, err)
		return
	}

	http.Redirect(w, r, fmt.Sprintf("/snippet/view/%d", id), http.StatusSeeOther)
}
