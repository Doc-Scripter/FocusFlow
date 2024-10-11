package main

import (
	"fmt"
	"net/http"

	"Focus/handler"
	"Focus/service"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case "/":
			handler.RegisterPageHandler(w, r)

		case "/register":
			handler.RegisterHandler(w, r)
		case "/AddEvent":
			handler.AddEventHandler(w, r)
		case "/Delete":
			handler.DeleteEventHandler(w, r)
		case "/login":
			handler.LoginPageHandler(w, r)
		case "/auth":
			handler.LoginHandler(w, r)
		case "/contact":
			handler.ContactPageHandler(w, r)
		case "/logout":
			handler.LogoutHandler(w, r)
		default:
			http.Error(w, "404 Not Found", http.StatusNotFound)
		}
	})
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./web/static/"))))

	go service.StartAlarmLoop()
	fmt.Println("Server started at :5050")
	http.ListenAndServe(":5050", nil)
}
