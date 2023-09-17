package main

import (
	"html/template"
	"log"
	"net/http"
)

type Film struct {
	Title    string
	Director string
}

func main() {

	http.HandleFunc("/", httpHandlerIndex)
	// Catch error and logs it and exit the application
	log.Fatal(http.ListenAndServe(":8000", nil))
}

func httpHandlerIndex(resposeWriter http.ResponseWriter, request *http.Request) {
	films := map[string][]Film{
		"Films": {
			{Title: "The Green Mile", Director: "X"},
			{Title: "Forest Gump", Director: "Y"},
		},
	}

	tpl := template.Must(template.ParseFiles("templates/index.html"))
	tpl.Execute(resposeWriter, films)
}
