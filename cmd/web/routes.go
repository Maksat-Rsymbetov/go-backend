package main

import "net/http"

func (app *application) routes() *http.ServeMux {
	mux := http.NewServeMux()

	// Fileserver to serve static files (css, js)
	fileserver := http.FileServer(http.Dir("./ui/static"))
	mux.Handle("/static/", http.StripPrefix("/static", fileserver))

	mux.HandleFunc("/", app.home)
	mux.HandleFunc("/product/create", app.productCreate)
	mux.HandleFunc("/product/view", app.productView)
	mux.HandleFunc("/about", app.about)

	mux.HandleFunc("/user/signup/", app.userSignup)
	mux.HandleFunc("/user/login/", app.userLogin)
	mux.HandleFunc("/user/logout/", app.userLogoutPost)

	return mux
}
