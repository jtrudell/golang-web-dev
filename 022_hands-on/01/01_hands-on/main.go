// ListenAndServe on port ":8080" using the default ServeMux.
//
// Use HandleFunc to add the following routes to the default ServeMux:
// "/"
// "/dog/"
// "/me/
//
// Add a func for each of the routes.
// Have the "/me/" route print out your name.

package main

import (
	"fmt"
	"net/http"
)

func home(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintln(w, "<h1>You are home</h1>")
}

func dog(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintln(w, "<h1>Woof</h1>")
}

func me(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintln(w, "<h1>Hey Jen</h1>")
}

func main() {
	http.HandleFunc("/", home)
	http.HandleFunc("/dog/", dog)
	http.HandleFunc("/me/", me)

	http.ListenAndServe(":8080", nil)
}
