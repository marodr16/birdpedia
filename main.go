package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

// create router and return it
func newRouter() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/hello", handler).Methods("GET")

	// declare static file directory and point it to assets
	// directory
	staticFileDirectory := http.Dir("/assets/")

	// declare the handler that routs requests to filenames
	// while stripping prefix
	staticFileHandler := http.StripPrefix("/assets/", http.FileServer(staticFileDirectory))

	// mach all routes that start with the assets prefix
	r.PathPrefix("/assets/").Handler(staticFileHandler).Methods("GET")

	return r

}

func main() {
	// declare new router
	r := newRouter()

	// listen on port 8080
	http.ListenAndServe(":8080", r)

}

func handler(w http.ResponseWriter, r *http.Request) {
	// pipe hello world into the response
	fmt.Fprintf(w, "Hello World!")
}
