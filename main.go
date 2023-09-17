package main

import (
	"io"
	"log"
	"net/http"
)

func main() {

	http.HandleFunc("/", httpHandlerIndex)
	// Catch error and logs it and exit the application
	log.Fatal(http.ListenAndServe(":8000", nil))
}

func httpHandlerIndex(resposeWriter http.ResponseWriter, request *http.Request) {
	io.WriteString(resposeWriter, "Hello World\n")
	io.WriteString(resposeWriter, request.Method)
}
