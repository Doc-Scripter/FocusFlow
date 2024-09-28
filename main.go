package main

import (
	"Focus/handler"
	"net/http"
)

func main() {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case "/":
			handler.Homehandler(w, r)
		case "/NewEvent":
			handler.NewEventhandler(w, r)
		case "/AddEvent":
			handler.AddEventHandler(w, r)
		default:
			http.Error(w, "404 Not Found", http.StatusNotFound)
		}
	})
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./web/static/"))))
	http.ListenAndServe(":8080", nil)
}
