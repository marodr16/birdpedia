package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	// declare new router
	r := mux.NewRouter()

	r.HandleFunc("/hello", handler).Methods("GET")

	// http server
	//http.HandleFunc("/", handler)

	// listen on port 8080
	http.ListenAndServe(":8080", r)

}

func handler(w http.ResponseWriter, r *http.Request) {
	// pipe hello world into the response
	fmt.Fprintf(w, "Hello World!")
}
