package service

import (
	"fmt"
	"strconv"
	"time"
)

var Words struct {
	name string
	task string
	time time.Time
}

func Input(str []string) {
for i:=range str{
	switch i {
	case 0:
		Words.name = str[i]
	case 1:
		Words.task = str[i]
	case 2:
		timed, _ := strconv.Atoi(str[i])
		Words.time = time.Now().Add(time.Duration(timed) * time.Hour)
	}


	if  i > 2 {
		// return Words
		fmt.Printf("Hi %s,your task %s is at %s\n",Words.name,Words.task, Words.time.Format("2006-01 15:04:05"))
	}
}
}