package service

import (
	"fmt"
	"strconv"
	"sync"
	"time"
)

var Words struct {
	name  string
	task  string
	time  time.Time
	alarm time.Timer
}

func Input(str []string) {
	var wg sync.WaitGroup
	for i := range str {
		switch i {
		case 0:
			Words.name = str[i]
		case 1:
			Words.task = str[i]
		case 2:
			timed, _ := strconv.Atoi(str[i])
			Words.time = time.Now().Add(time.Duration(timed) * time.Second)
			Words.alarm = *time.NewTimer(time.Until(Words.time))
		}

		if i == 2{
			wg.Add(1)
			go func() {
				defer wg.Done()
				<-Words.alarm.C
				fmt.Printf("%s time\n",Words.task)
			}()

			fmt.Printf("Hi %s,your task \"%s\" is at %s\n", Words.name, Words.task, Words.time.Format("2006-01 15:04:05"))
		}
	}
	wg.Wait()
}
