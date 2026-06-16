package main

import (
	"fmt"
	"net/http"
	"strconv"

	"snippetbox.innolabs.ai/components"
)

func (conf *Config) homeHandler(w http.ResponseWriter, r *http.Request) {
	//* priority: first
	w.Header().Add("Server", "Go")
	//* priority: second
	// w.WriteHeader(http.StatusOK)
	//* priority: third
	// w.Write([]byte("Hello from Snippetbox"))

	err := components.Home().Render(r.Context(), w)
	if err != nil {
		conf.serverError(w, r, err)
	}
}

func (conf *Config) snippetView(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		http.NotFound(w, r)
		return
	}

	fmt.Fprintf(w, "Display a specific snippet with ID %d...", id)
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
