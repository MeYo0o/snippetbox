package components

import "snippetbox.innolabs.ai/internal/validator"

type SnippetCreateForm struct {
	Title   string
	Content string
	Expires int
	validator.Validator
}
