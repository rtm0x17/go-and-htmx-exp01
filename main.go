package main

import (
	"html/template"
	"log"
	"net/http"
)

func main() {

	http.HandleFunc("/", httpHandlerIndex)
	// Catch error and logs it and exit the application
	log.Fatal(http.ListenAndServe(":8000", nil))
}

func httpHandlerIndex(resposeWriter http.ResponseWriter, request *http.Request) {
	tpl := template.Must(template.ParseFiles("templates/index.html"))
	tpl.Execute(resposeWriter, nil)
}
