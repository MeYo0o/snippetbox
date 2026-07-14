package main

import (
	"net/http"

	"github.com/justinas/nosurf"
	"snippetbox.innolabs.ai/components"
)

func (app *application) newTemplateData(r *http.Request) components.TemplateData {
	return components.TemplateData{
		Flash:           app.popFlash(r),
		IsAuthenticated: app.isAuthenticated(r),
		CSRFToken:       nosurf.Token(r),
	}
}
