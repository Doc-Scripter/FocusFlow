package structs

const UserData = "./database/users.json"

type Event struct {
	AddEvent string `json:"addEvent"`
	Date     string `json:"date"`
	Time     string `json:"time"`
}

type Data struct {
	Events []Event
	Alert  string
}



const UserEvents = "./database/events.json"

const UserSchedule ="./database/time.json"

var IsOnline = make(map[string]bool)

type OnlineUsers struct {
	OnlineUser []string
}

type Period struct{
	Date string`json:"date"`
	Time string `json:"time"`
}

