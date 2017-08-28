// # ListenAndServe on port 8080 of localhost
//
// For the default route "/" Have a func called "foo"
// which writes to the response "foo ran"
//
// For the route "/dog/" Have a func called "dog"
// which parses a template called "dog.gohtml"
// and writes to the response "<h1>This is from dog</h1>"
// and also shows a picture of a dog when the template is executed.
//
// Use "http.ServeFile" to serve the file "dog.jpeg"

package main

import (
	"html/template"
	"io"
	"log"
	"net/http"
)

func foo(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "foo ran")
}

func dog(w http.ResponseWriter, req *http.Request) {
	tpl := template.Must(template.ParseFiles("dog.gohtml"))
	err := tpl.ExecuteTemplate(w, "dog.gohtml", nil)
	if err != nil {
		log.Fatal(err)
	}
}

func dogJpg(w http.ResponseWriter, req *http.Request) {
	http.ServeFile(w, req, "dog.jpg")
}

func main() {
	http.HandleFunc("/", foo)
	http.HandleFunc("/dog/", dog)
	http.HandleFunc("/dog.jpg", dogJpg)
	http.ListenAndServe(":8080", nil)
}
