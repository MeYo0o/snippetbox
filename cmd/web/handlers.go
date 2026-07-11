package main

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"snippetbox.innolabs.ai/components"
	"snippetbox.innolabs.ai/internal/models"
	"snippetbox.innolabs.ai/internal/validator"
)

func (app *application) homeHandler(w http.ResponseWriter, r *http.Request) {
	//* priority: first
	// w.Header().Add("Server", "Go")
	//* priority: second
	// w.WriteHeader(http.StatusOK)
	//* priority: third
	// w.Write([]byte("Hello from Snippetbox"))

	snippets, err := app.snippets.Latest()
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

	snippet, err := app.snippets.Get(int32(id))
	if err != nil {
		if errors.Is(err, models.ErrNoRecord) {
			http.NotFound(w, r)
		} else {
			app.serverError(w, r, err)
		}
		return
	}

	flash := app.sessionManager.PopString(r.Context(), "flash")

	err = components.ViewSnippet(snippet, flash).Render(r.Context(), w)
	if err != nil {
		app.serverError(w, r, err)
	}
}

func (app *application) snippetCreate(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)

	err := components.CreateSnippet(components.SnippetCreateForm{
		Expires: 365,
	}).Render(r.Context(), w)
	if err != nil {
		app.serverError(w, r, err)
	}
}

func (app *application) snippetCreatePost(w http.ResponseWriter, r *http.Request) {

	// new form entity
	var form components.SnippetCreateForm

	// this will run a decoder check on the already assigned fields on the SnippetCreateForm struct and will map all values to the struct fields upon successful decode.
	err := app.decodePostForm(r, &form)
	if err != nil {
		app.clientError(w, http.StatusBadRequest)
		return
	}

	form.CheckField(validator.NotBlank(form.Title), "title", "This field cannot be blank")
	form.CheckField(validator.MaxChars(form.Title, 100), "title", "This field cannot be more than 100 characters long")
	form.CheckField(validator.NotBlank(form.Content), "content", "This field cannot be blank")
	form.CheckField(validator.PermittedValue(form.Expires, 1, 7, 365), "expires", "This field must equal 1, 7 or 365")

	if !form.Valid() {
		components.CreateSnippet(form).Render(r.Context(), w)
		return
	}

	id, err := app.snippets.Insert(form.Title, form.Content, form.Expires)
	if err != nil {
		app.serverError(w, r, err)
		return
	}

	app.sessionManager.Put(r.Context(), "flash", "Snippet successfully created!")

	http.Redirect(w, r, fmt.Sprintf("/snippet/view/%d", id), http.StatusSeeOther)
}

func (app *application) userSignup(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	components.SignUp(components.UserSignUpForm{}).Render(r.Context(), w)
}

func (app *application) userSignupPost(w http.ResponseWriter, r *http.Request) {
	var form components.UserSignUpForm

	err := app.decodePostForm(r, &form)
	if err != nil {
		app.clientError(w, http.StatusBadRequest)
		return
	}

	form.CheckField(validator.NotBlank(form.Name), "name", "This field cannot be blank")
	form.CheckField(validator.NotBlank(form.Email), "email", "This field cannot be blank")
	form.CheckField(validator.NotBlank(form.Password), "password", "This field cannot be blank")
	form.CheckField(validator.Matches(form.Email, validator.EmailRX), "email", "This field must be a valid email address")
	form.CheckField(validator.MinChars(form.Password, 8), "password", "This field must be at least 8 characters long")
	form.CheckField(validator.MaxBytes(form.Password, 72), "password", "This field must not be more than 72 bytes long")

	if !form.Valid() {
		err := components.SignUp(form).Render(r.Context(), w)
		if err != nil {
			app.serverError(w, r, err)
		}
		return
	}

	err = app.users.Insert(form.Name, form.Email, form.Password)
	if err != nil {
		if errors.Is(err, models.ErrDuplicateEmail) {
			form.AddFieldErrorKey("email", "Email address is already in use")
			err := components.SignUp(form).Render(r.Context(), w)
			if err != nil {
				app.serverError(w, r, err)
			}
			return
		} else {
			app.serverError(w, r, err)
		}

		return
	}

	app.sessionManager.Put(r.Context(), "flash", "your signup was successful. please login.")

	http.Redirect(w, r, "/user/login", http.StatusSeeOther)
}

func (app *application) userLogin(w http.ResponseWriter, r *http.Request) {
	err := components.Login(components.UserLoginForm{}).Render(r.Context(), w)
	if err != nil {
		app.serverError(w, r, err)
	}
}

func (app *application) userLoginPost(w http.ResponseWriter, r *http.Request) {
	var form components.UserLoginForm

	err := app.decodePostForm(r, &form)
	if err != nil {
		app.clientError(w, http.StatusBadRequest)
		return
	}

	form.CheckField(validator.NotBlank(form.Email), "email", "this field cannot be blank")
	form.CheckField(validator.Matches(form.Email, validator.EmailRX), "email", "This field must be a valid email address")
	form.CheckField(validator.NotBlank(form.Password), "password", "This field cannot be blank")
	form.CheckField(validator.MaxBytes(form.Password, 72), "password", "This field must not be more than 72 bytes long")

	if !form.Valid() {
		err := components.Login(form).Render(r.Context(), w)
		if err != nil {
			app.serverError(w, r, err)
		}
		return
	}

	id, err := app.users.Authenticate(form.Email, form.Password)
	if err != nil {
		if errors.Is(err, models.ErrInvalidCredentials) {
			form.AddNonFieldError("Email or password is incorrect")
			err = components.Login(form).Render(r.Context(), w)
			if err != nil {
				app.serverError(w, r, err)
			}
			return
		}
		app.serverError(w, r, err)
		return
	}

	err = app.sessionManager.RenewToken(r.Context())
	if err != nil {
		app.serverError(w, r, err)
		return
	}

	app.sessionManager.Put(r.Context(), "authenticatedUserID", id)

	http.Redirect(w, r, "/snippet/create", http.StatusSeeOther)
}

func (app *application) userLogoutPost(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Logout the user...")
}
