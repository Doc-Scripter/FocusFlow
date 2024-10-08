package structs

const UserData = "./database/users.json"

type Event struct {
	AddEvent string `json:"addEvent"`
	Date     string `json:"date"`
	Time     string `json:"time"`
}