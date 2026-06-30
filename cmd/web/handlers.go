package main

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"unicode/utf8"

	"snippetbox.innolabs.ai/components"
	"snippetbox.innolabs.ai/internal/models"
)

func (app *application) homeHandler(w http.ResponseWriter, r *http.Request) {
	//* priority: first
	// w.Header().Add("Server", "Go")
	//* priority: second
	// w.WriteHeader(http.StatusOK)
	//* priority: third
	// w.Write([]byte("Hello from Snippetbox"))

	snippets, err := app.Snippets.Latest()
	if err != nil {
		app.serverError(w, r, err)
		return
	}

	err = components.Home(snippets).Render(r.Context(), w)
	if err != nil {
		app.serverError(w, r, err)
	}
}

func (app *application) snippetView(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		http.NotFound(w, r)
		return
	}

	snippet, err := app.Snippets.Get(int32(id))
	if err != nil {
		if errors.Is(err, models.ErrNoRecord) {
			http.NotFound(w, r)
		} else {
			app.serverError(w, r, err)
		}
		return
	}

	err = components.ViewSnippet(snippet).Render(r.Context(), w)
	if err != nil {
		app.serverError(w, r, err)
	}
}

func (app *application) snippetCreate(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)

	components.CreateSnippet(components.SnippetCreateForm{
		Expires: 365,
	}).Render(r.Context(), w)
}

func (app *application) snippetCreatePost(w http.ResponseWriter, r *http.Request) {

	err := r.ParseForm()
	if err != nil {
		app.clientError(w, http.StatusBadRequest)
		return
	}

	title := r.PostForm.Get("title")
	content := r.PostForm.Get("content")

	expires, err := strconv.Atoi(r.PostForm.Get("expires"))
	if err != nil {
		app.clientError(w, http.StatusBadRequest)
		return
	}

	form := components.SnippetCreateForm{
		Title:       title,
		Content:     content,
		Expires:     expires,
		FieldErrors: map[string]string{},
	}

	if strings.TrimSpace(title) == "" {
		form.FieldErrors["title"] = "This field cannot be blank"
	} else if utf8.RuneCountInString(title) > 100 {
		form.FieldErrors["title"] = "This field cannot be more than 100 characters long"
	}

	if strings.TrimSpace(content) == "" {
		form.FieldErrors["content"] = "This field cannot be blank"
	}

	if expires != 1 && expires != 7 && expires != 365 {
		form.FieldErrors["expires"] = "This field must be equal to 1, 7 or 365"
	}

	if len(form.FieldErrors) > 0 {
		w.WriteHeader(http.StatusUnprocessableEntity)
		components.CreateSnippet(form).Render(r.Context(), w)
		return
	}

	id, err := app.Snippets.Insert(form.Title, form.Content, form.Expires)
	if err != nil {
		app.serverError(w, r, err)
		return
	}

	http.Redirect(w, r, fmt.Sprintf("/snippet/view/%d", id), http.StatusSeeOther)
}
