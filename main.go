package main

import (
	"Focus/handler"
	"net/http"
)

func main() {
	http.HandleFunc("/", handler.Homehandler)
	http.ListenAndServe(":8080", nil)
}
