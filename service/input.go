package service

import (
	"Focus/structs"
	"encoding/json"
	"fmt"
	"os"
	"sync"
	"time"
)

func Input() string {
	var wg sync.WaitGroup
	var existingData []structs.Event
	file, err := os.ReadFile(structs.UserData)
	if err != nil {
		return ""
	}
	json.Unmarshal(file, &existingData)
	var inputDate string
	var inputTime string
	var inputDateTime string
	for _, v := range existingData {
		inputDate = v.Date
		inputTime = v.Time
	}
	inputDateTime=inputDate+inputTime
	//Parse date
	targetTime, _ := time.Parse("2006-01-02", inputDateTime)
	//calculate duration until target time
	durationUntilTarget := targetTime.Sub(time.Now())

	//Create timer
	timer := time.NewTimer(durationUntilTarget)
	// Event.time = time.Now().Add(time.Duration(timed) * time.Second)
	// Words.alarm = *time.NewTimer(time.Until(Words.time))

	go func() {
		defer wg.Done()
		<-timer.C
		fmt.Print(" time\n")
	}()
	wg.Add(1)
	wg.Wait()
	return "alert"
}
