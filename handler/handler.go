package handler

import (
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"
	"os"
)

type Event struct {
	AddEvent string `json:"addEvent"`
	Date     string `json:"date"`
	Time     string `json:"time"`
}

type Data struct {
	Events []Event
}

const users = "./database/users.json"

func Homehandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, "Method not Allowed", http.StatusMethodNotAllowed)
		return
	}
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
	tmpl, err := template.ParseFiles("./web/templates/NewEvent.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	var existingData []Event
	file, err := os.ReadFile("./database/events.json")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.Unmarshal(file, &existingData)
	data := Data{Events: existingData}
	tmpl.Execute(w, data)
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

	newEvent := Event{
		AddEvent: r.FormValue("AddEvent"),
		Date:     r.FormValue("Date"),
		Time:     r.FormValue("Time"),
	}
	events = append(events, newEvent)

	jsonData, err := json.MarshalIndent(events, "", "\t")
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
	data := Data{Events: events}
	err = t.Execute(w, data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return

	}
}

func DeleteEventHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return

	}
	var events []Event

	file, err := os.ReadFile("./database/events.json")
	if err != nil {
		http.Error(w, "hello", http.StatusInternalServerError)
		return

	}
	json.Unmarshal(file, &events)
	for k, v := range events {
		fmt.Println(r.FormValue("DeleteEvent"))
		if v.AddEvent == r.FormValue("DeleteEvent") {
			events = append(events[:k], events[k+1:]...)
		}
	}
	jsonFile, err := json.Marshal(events)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return

	}
	err = os.WriteFile("./database/events.json", jsonFile, 0o666)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return

	}
	tmpl, err := template.ParseFiles("./web/templates/NewEvent.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	data := Data{Events: events}
	tmpl.Execute(w, data)
}

type Users struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func RegisterHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("hey")
	if r.Method != "POST" {
		http.Error(w, "Method not Allowed", http.StatusMethodNotAllowed)
	}
	var user []Users

	file, err := os.ReadFile(users)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.Unmarshal([]byte(file), &user)

	newUsers := Users{
		r.FormValue("username"),
		r.FormValue("email"),
		r.FormValue("password"),
	}
	json.Unmarshal([]byte(users), &user)

	user = append(user, newUsers)

	JsonUser, err := json.MarshalIndent(user, "", "\t")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	os.WriteFile(users, JsonUser, 0o777)
	http.Redirect(w, r, "/login", http.StatusSeeOther)
}

func RegisterPageHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("hi")
	if r.Method != "GET" {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}
	http.ServeFile(w, r, "./web/templates/Register.html")
}

func LoginPageHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, "Method not Allowed", http.StatusMethodNotAllowed)
		return
	}
	http.ServeFile(w, r, "./web/templates/login.html")
}
