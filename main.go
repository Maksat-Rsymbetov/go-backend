package main

import (
	"log"
	"net/http"
)

func home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	w.Write([]byte("Hello from webiste"))
}

func page(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("You go to another page"))
}

func main() {
	var mux = http.NewServeMux()
	mux.HandleFunc("/", home)
	mux.HandleFunc("/page", page)

	log.Println("Starting server on :4000")
	var err error = http.ListenAndServe(":4000", mux)
	log.Fatal(err)
}
