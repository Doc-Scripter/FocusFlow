package handler

import (
	"encoding/json"
	"html/template"
	"net/http"
	"os"
)

type Event struct {
	AddEvent string
	// Date string
}

type Data struct{
	Events []Event
}

func Homehandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, "Method not Allowed", http.StatusMethodNotAllowed)
		return
	}
	// fs:=http.FileServer(http.Dir("/handler"))
	t, err := template.ParseFiles("./web/templates/index.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	err = t.Execute(w, nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func NewEventhandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
	}
	http.ServeFile(w, r, "./web/templates/NewEvent.html")
}

func AddEventHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
	}
	// fmt.Println("hi")
	var events []Event
	file, err := os.ReadFile("./database/events.json")
	if err != nil {
		http.Error(w, "hello", http.StatusInternalServerError)
		return
	}
	json.Unmarshal(file, &events)
	// if err != nil {
	// 	http.Error(w, "hi", http.StatusInternalServerError)
	// 	return
	// }
	newEvent := Event{
		r.FormValue("AddEvent"),
	}

	events = append(events, newEvent)

	jsonData, err := json.Marshal(events)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	err = os.WriteFile("./database/events.json", jsonData, 0o666)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	t, err := template.ParseFiles("./web/templates/NewEvent.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	data:=Data{Events: events}
	err = t.Execute(w, data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

}
