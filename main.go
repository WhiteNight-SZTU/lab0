package main

import (
	"fmt"
	"html"
	"log"
	"net/http"
)

func main() {
	// TODO: some code goes here
	// Fill out the HomeHandler function in handlers/handlers.go which handles the user's GET request.
	// Start an http server using http.ListenAndServe that handles requests using HomeHandler.
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello,%q", html.EscapeString(r.URL.Path))
	})
	log.Fatal(http.ListenAndServe(":8085", nil))
}
