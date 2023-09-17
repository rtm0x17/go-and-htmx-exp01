package main

import (
	"html/template"
	"log"
	"net/http"
	"time"
)

type Film struct {
	Title    string
	Director string
}

func main() {

	http.HandleFunc("/", httpHandlerIndex)
	http.HandleFunc("/add-film", httpHandlerAddFilm)

	// Catch error and logs it and exit the application
	log.Fatal(http.ListenAndServe(":8000", nil))
}

func httpHandlerAddFilm(resposeWriter http.ResponseWriter, request *http.Request) {
	log.Print("HTMX request")
	log.Print(request.Header.Get("HX-Request"))

	title := request.PostFormValue("title")
	director := request.PostFormValue("director")
	time.Sleep(1 * time.Second)
	tpl := template.Must(template.ParseFiles("templates/index.html"))
	tpl.ExecuteTemplate(resposeWriter, "film-list-item", Film{Title: title, Director: director})
}

//func httpHandlerAddFilm(resposeWriter http.ResponseWriter, request *http.Request) {
//	log.Print("HTMX request")
//	log.Print(request.Header.Get("HX-Request"))
//
//	title := request.PostFormValue("title")
//	director := request.PostFormValue("director")
//
//	htmlStr := fmt.Sprintf("<p>%s - %s</p>", title, director)
//
//	tmpl, _ := template.New("t").Parse(htmlStr)
//	time.Sleep(1 * time.Second)
//	err := tmpl.Execute(resposeWriter, nil)
//
//	if err != nil {
//		return
//	}
//}

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
