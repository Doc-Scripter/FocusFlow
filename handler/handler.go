package handler

import (
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"

	"Focus/structs"
)
type Users struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}
type TotalUsers struct {
	Total []Users
}
func AddEventHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
	}
	// fmt.Println("hi")
	var events []structs.Event
	var schedule []structs.Period

	file, err := os.ReadFile(structs.UserEvents)
	if err != nil {
		http.Error(w, "hello", http.StatusInternalServerError)
		return
	}
	json.Unmarshal(file, &events)
	timeFile, err := os.ReadFile("./database/time.json")
	if err != nil {
		http.Error(w, "hello", http.StatusInternalServerError)
		return
	}

	newEvent := structs.Event{
		AddEvent: r.FormValue("AddEvent"),
		Date:     r.FormValue("Date"),
		Time:     r.FormValue("Time"),
	}
	newSchedule :=structs.Period{
		Date: r.FormValue("Date"),
		Time: r.FormValue("Time"),
	}
	json.Unmarshal(timeFile, &schedule)

	// alert := service.Input()
	schedule = append(schedule, newSchedule)
	events = append(events, newEvent)

	jsonTime, err := json.MarshalIndent(schedule, "", "\t")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	err = os.WriteFile(structs.UserSchedule, jsonTime, 0o666)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

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

	t, err := template.ParseFiles("./web/templates/index2.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	data := structs.Data{
		Events: events,
		// Alert: alert,
	}

	err = t.Execute(w, data)
	if err != nil {
		log.Println(err)
		http.Error(w, "error", http.StatusInternalServerError)
		return

	}
}

func DeleteEventHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return

	}
	var events []structs.Event

	file, err := os.ReadFile("./database/events.json")
	if err != nil {
		http.Error(w, "hello", http.StatusInternalServerError)
		return

	}
	json.Unmarshal(file, &events)
	for k, v := range events {
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
	tmpl, err := template.ParseFiles("./web/templates/index2.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	data := structs.Data{Events: events}
	tmpl.Execute(w, data)
}

func RegisterHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("hey")
	if r.Method != "POST" {
		http.Error(w, "Method not Allowed", http.StatusMethodNotAllowed)
		return
	}
	var user []Users

	file, err := os.ReadFile(structs.UserEvents)
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
	json.Unmarshal([]byte(structs.UserData), &user)

	user = append(user, newUsers)

	JsonUser, err := json.MarshalIndent(user, "", "\t")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	os.WriteFile(structs.UserData, JsonUser, 0o777)
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
	temp, err := template.ParseFiles("./web/templates/login.html")
	if err != nil {
		http.Error(w, err.Error(), 404)
		return
	}
	temp.Execute(w, nil)
}

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Method not Allowed", http.StatusMethodNotAllowed)
		return
	}
	var AllUsers TotalUsers
	user := Users{
		Username: r.FormValue("username"),
		// Email:    r.FormValue("email"),
		Password: r.FormValue("password"),
	}

	ExistingUsers, err := os.ReadFile(structs.UserData)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	var authenticated bool
	json.Unmarshal(ExistingUsers, &AllUsers.Total)
	
	for _, v := range AllUsers.Total {
		if v.Username == user.Username || v.Password == user.Password {
			structs.IsOnline[v.Username] = true
			authenticated = true
		}
	}
	if !authenticated {
		http.Redirect(w, r, "/login", http.StatusSeeOther)
		return
	}
	tmpl, err := template.ParseFiles("./web/templates/index2.html")
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	var existingData []structs.Event
	file, err := os.ReadFile("./database/events.json")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.Unmarshal(file, &existingData)
	log.Println(existingData)
	data := structs.Data{Events: existingData}
	tmpl.Execute(w, data)
}

func LogoutHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Method not Allowed", http.StatusMethodNotAllowed)
		return
	}

	user := Users{
		Username: r.FormValue("username"),
	}

	structs.IsOnline[user.Username] = false

	http.Redirect(w, r, "/login", http.StatusSeeOther)
}

func ContactPageHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, "Method not Allowed", http.StatusMethodNotAllowed)
		return
	}
	
	var displayOnline structs.OnlineUsers
	for k := range structs.IsOnline {
		displayOnline.OnlineUser = append(displayOnline.OnlineUser, k)
	}
	tmpl, err := template.ParseFiles("./web/templates/contacts.html")
	if err != nil {
		http.Error(w, err.Error(), 404)
		return
	}
	tmpl.Execute(w, displayOnline)
}
