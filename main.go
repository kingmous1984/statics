package main

import (
	"fmt"
	"html/template"
	"net/http"
)

// IndexHandler index page
func IndexHandler(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseFiles("template/views/index.html"))
	t.Execute(w, "")
}

func main() {
	fmt.Println("Hello world!")
	http.Handle("/statics/", http.StripPrefix("/statics/", http.FileServer(http.Dir("template/statics/"))))
	http.HandleFunc("/", IndexHandler)

	http.ListenAndServe(":8080", nil)
}
