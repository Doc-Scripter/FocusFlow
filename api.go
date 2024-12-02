package handler
import (
	"fmt"
	"net/http"

	"Focus/handler"
	"Focus/service"
)

func init() {
	go service.StartAlarmLoop()
	fmt.Println("Alarm loop started")
}

func Handler(w http.ResponseWriter, r *http.Request) {
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
}

func StaticHandler(w http.ResponseWriter, r *http.Request) {
	http.StripPrefix("/static/", http.FileServer(http.Dir("./web/static/"))).ServeHTTP(w, r)
}