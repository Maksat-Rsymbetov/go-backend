package main

import (
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()

	fileserver := http.FileServer(http.Dir("./ui/static/"))
	mux.Handle("/static/", http.StripPrefix("/static", fileserver))

	mux.HandleFunc("/", home)
	mux.HandleFunc("/page/view", pageView)
	mux.HandleFunc("/page/create", pageCreate)

	log.Println("Startnig server on port 4000")
	log.Fatal(http.ListenAndServe(":4000", mux))
}
