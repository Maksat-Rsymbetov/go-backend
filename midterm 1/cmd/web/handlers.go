package main

import (
	"errors"
	"fmt"
	// "html/template"

	// "html/template"
	"net/http"
	"strconv"

	"website.maksat.com/internal/models"
)

// Home page --------------------------------------------------------------
func (app *application) home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		app.notFound(w)
		return
	}

	if r.Method == "POST" {

	}

	plist, err := app.products.GetList()
	if err != nil {
		app.serverError(w, err)
	}

	app.render(w, 200, "home.html", &BaseTemplate{Products: plist})
}

// Product view --------------------------------------------------------------
func (app *application) productView(w http.ResponseWriter, r *http.Request) {
	var id, err = strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil || id < 1 {
		// http.NotFound(w, r)
		app.errorLog.Println(err)
		app.notFound(w)
		return
	}

	product, err := app.products.Get(id)
	if err != nil {
		if errors.Is(err, models.ErrNoRecord) {
			app.errorLog.Println(err)
			app.notFound(w)
		} else {
			app.errorLog.Println(err)
			app.serverError(w, err)
		}
		return
	}

	app.render(w, 200, "product.html", &BaseTemplate{Product: product})
}

// Create a new product (not used for now) -------------------------------------
func (app *application) productCreate(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.Header().Set("Allow", http.MethodPost)
		app.clientError(w, http.StatusMethodNotAllowed)
		return
	}

	name := "Amazon kindle"
	description := "buy a kindle book"
	price := 5000

	id, err := app.products.Insert(name, description, price)
	if err != nil {
		app.serverError(w, err)
		return
	}

	http.Redirect(w, r, fmt.Sprintf("/product/view?id=%v", id), http.StatusSeeOther)
}

// User Signup -------------------------------------------------------------------
func (app *application) userSignup(w http.ResponseWriter, r *http.Request) {
	// tmp, err := template.ParseFiles("./ui/html/pages/signup.html")
	// if err != nil { app.serverError(w, err) }

	if r.Method == "GET" {
		http.ServeFile(w, r, "./ui/html/pages/signup.html")
	} else if r.Method == "POST" {
		if err := r.ParseForm(); err != nil {
			app.serverError(w, err)
		}
		username := r.FormValue("username")
		password := r.FormValue("password")
		_, err := app.users.AddUser(username, password)
		if err != nil {
			app.serverError(w, err)
			return
		}

		http.Redirect(w, r, "/", http.StatusSeeOther)
	}
}

// func (app *application) userSignupPost(w http.ResponseWriter, r *http.Request) {
// 	fmt.Fprintln(w, "Create a new user...")
// }

func (app *application) userLogin(w http.ResponseWriter, r *http.Request) {
	// fmt.Fprintln(w, "Display a form to login")
	if r.Method == "GET" {
		http.ServeFile(w, r, "./ui/html/pages/login.html")
	} else if r.Method == "POST" {
		if err := r.ParseForm(); err != nil {
			app.serverError(w, err)
		}
		username := r.FormValue("username")
		password := r.FormValue("password")
		_, err := app.users.GetUser(username, password)
		if err != nil {
			fmt.Fprintln(w, "No such user or wrong password")
			return
		}
		http.Redirect(w, r, "/", http.StatusSeeOther)
	}
}

// func (app *application) userLoginPost(w http.ResponseWriter, r *http.Request) {
// 	fmt.Fprintln(w, "Login the user")
// }

func (app *application) userLogoutPost(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Logout the user ...")
}

// About page (also not used for now) --------------------------------------------
func (app *application) about(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("About this website"))
}
