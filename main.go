package main

import (
	"Focus/handler"
	"net/http"
)

func main() {

	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./web/static/"))))
	http.HandleFunc("/", handler.Homehandler)
	http.ListenAndServe(":8080", nil)
}
