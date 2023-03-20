package main

import (
	"fmt"
	"net/http"
	"runtime/debug"
)

// Write an error message and stack trace to the errorLog, then send error 500
func (app *application) serverError(w http.ResponseWriter, err error) {
	trace := fmt.Sprintf("%s\n%s", err.Error(), debug.Stack())
	app.errorLog.Output(2, trace)

	http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
}

// Send error with corresponding description to the user
func (app *application) clientError(w http.ResponseWriter, status int) {
	http.Error(w, http.StatusText(status), status)
}

// NotFound helper made for consistency
func (app *application) notFound(w http.ResponseWriter) {
	app.clientError(w, http.StatusNotFound)
}

// Rendering the page
func (app *application) render(w http.ResponseWriter, status int, page string, data *BaseTemplate) {
	tmpl, ok := app.templateCache[page]
	if !ok {
		err := fmt.Errorf("the template %s does not exist", page)
		app.serverError(w, err)
		return
	}
	w.WriteHeader(status)
	err := tmpl.ExecuteTemplate(w, "base", data)
	if err != nil {
		app.serverError(w, err)
	}
}
