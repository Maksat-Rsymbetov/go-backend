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
		// searchVal := r.FormValue("search")
		// lowerPrice := r.FormValue("lowerPrice")
		// upperPrice := r.FormValue("upperPrice")
		// searchAddress := "/search?sv="+searchVal+"lp="lowerPrice+"up="+upperPrice
		searchVal := r.FormValue("search")
		lowerPrice := r.FormValue("lowerPrice")
		upperPrice := r.FormValue("upperPrice")
		searchAddress := fmt.Sprintf("/search?sv=%v&lp=%v&up=%v", searchVal, lowerPrice, upperPrice)
		http.Redirect(w, r, searchAddress, http.StatusSeeOther)
	}

	plist, err := app.products.GetList()
	if err != nil {
		app.serverError(w, err)
	}

	app.render(w, 200, "home.html", &BaseTemplate{Products: plist})
}

func (app *application) search(w http.ResponseWriter, r *http.Request) {

	if r.Method == "POST" {
		searchVal := r.FormValue("search")
		lowerPrice := r.FormValue("lowerPrice")
		upperPrice := r.FormValue("upperPrice")
		searchAddress := fmt.Sprintf("/search?sv=%v&lp=%v&up=%v", searchVal, lowerPrice, upperPrice)
		http.Redirect(w, r, searchAddress, http.StatusSeeOther)
	}

	searchVal := r.URL.Query().Get("sv")
	lowerPrice, err := strconv.Atoi(r.URL.Query().Get("lp"))
	if err != nil {
		app.serverError(w, err)
	}
	upperPrice, err := strconv.Atoi(r.URL.Query().Get("up"))
	if err != nil {
		app.serverError(w, err)
	}
	results, err := app.products.SearchResults(searchVal, lowerPrice, upperPrice)
	if err != nil {
		app.serverError(w, err)
	}
	tmp := &BaseTemplate{Products: results}
	app.render(w, 200, "home.html", tmp)
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
	if r.Method == http.MethodGet {
		// http.ServeFile(w, r, "./ui/html/pages/create.html")
		app.render(w, 200, "create.html", &BaseTemplate{})
	} else if r.Method == http.MethodPost {
		err := r.ParseForm()
		if err != nil {
			app.serverError(w, err)
		}
		var name, description string = r.FormValue("name"), r.FormValue("description")
		price, err := strconv.Atoi(r.FormValue("price"))
		// app.infoLog.Println(name, description, price)
		if err != nil {
			app.clientError(w, http.StatusBadRequest)
			return
		}
		id, err := app.products.Insert(name, description, price)
		if err != nil {
			app.serverError(w, err)
			return
		}

		http.Redirect(w, r, fmt.Sprintf("/product/view?id=%v", id), http.StatusSeeOther)
	}
	// if r.Method != http.MethodPost {
	// 	w.Header().Set("Allow", http.MethodPost)
	// 	app.clientError(w, http.StatusMethodNotAllowed)
	// 	return
	// }

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
