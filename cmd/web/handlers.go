package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"
)

func home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	var templates []string = []string{
		"./ui/html/base.html",
		"./ui/html/pages/home.html",
	}
	// ts, err := template.ParseFiles("./ui/html/pages/home.html") // initialize the html page
	ts, err := template.ParseFiles(templates...) // initialize the html page
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Internal server error", 500)
		return
	}

	// err = ts.Execute(w, nil) // show html page and handle an error
	err = ts.ExecuteTemplate(w, "base", nil)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Internal server error", 500)
	}
	// w.Write([]byte("Home page of out site"))
}

func pageView(w http.ResponseWriter, r *http.Request) {
	var id, err = strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil || id < 1 {
		http.NotFound(w, r)
		return
	}
	fmt.Fprintf(w, "Viewing page #%v", id)
}

func pageCreate(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.Header().Set("Allow", http.MethodPost)
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	w.Write([]byte("Creating a new page"))
}
