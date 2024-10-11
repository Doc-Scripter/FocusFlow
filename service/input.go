package service

import (
	"encoding/json"
	"log"
	"os"
	"time"

	"Focus/handler"
)

func GetEvents() ([]handler.Period, error) {
	var existingData []handler.Period
	file, err := os.ReadFile(handler.UserSchedule) // Adjust the path as necessary
	if err != nil {
		return nil, err
	}
	json.Unmarshal(file, &existingData)
	return existingData, nil
}

func StartAlarmLoop() {
	for {
		events, err := GetEvents()
		if err != nil {
			log.Println("Error reading events:", err)
			return
		}

		now := time.Now()
		var soonestAlarm *handler.Period
		var timeBefore time.Time
		for i := range events {
			eventTime, err := time.Parse("2006-01-02 15:04", events[i].Date+" "+events[i].Time)
			if err != nil {
				log.Println("Error parsing event time:", err)
				continue
			}
			// timeBefore=getEventTime(soonestAlarm)
			
			if eventTime.After(now) {//&&eventTime.Before(timeBefore)
				// timeBefore, err := time.Parse("2006-01-02 15:04", soonestAlarm.Date+" "+soonestAlarm.Time)
				// if err != nil {
				// 	log.Println("Error parsing event time:", err)
				// 	continue
				// }

				if soonestAlarm == nil || eventTime.Before(getEventTime(soonestAlarm)) {
				soonestAlarm = &events[i]
					}
			}
		}

		if soonestAlarm != nil {
			timeBefore, err := time.Parse("2006-01-02 15:04", soonestAlarm.Date+" "+soonestAlarm.Time)
			
			timeBefore= timeBefore.Add(-3 * time.Hour)
			log.Println(timeBefore)
			if err != nil {
				log.Println("Error parsing event time:", err)
				continue
			}
			sleepDuration := time.Until(timeBefore)
			log.Println(sleepDuration)
			time.Sleep(sleepDuration)

			// Trigger the alarm
			playSound()
			log.Println("Alert triggered for:", soonestAlarm.Date)

			// Optionally, remove the event after triggering
			// RemoveEvent(soonestAlarm)
		} else {
			// If there are no upcoming alarms, check every minute
			time.Sleep(1 * time.Second)
		}
	}
}
func getEventTime(event *handler.Period) time.Time {
	eventTime, err := time.Parse("2006-01-02 15:04", event.Date+" "+event.Time)
	if err != nil {
		log.Println("Error parsing event time:", err)
		return time.Time{} // Return zero Time if there's an error
	}
	return eventTime // Adjust for time zone if needed
}
// func Input() string {
// 	// var wg sync.WaitGroup
// 	var existingData []structs.Event
// 	file, err := os.ReadFile(structs.UserData)
// 	if err != nil {
// 		return ""
// 	}
// 	json.Unmarshal(file, &existingData)
// 	var inputDate string
// 	var inputTime string
// 	var inputDateTime string
// 	for _, v := range existingData {
// 		inputDate = v.Date
// 		inputTime = v.Time
// 	}
// 	inputDateTime = inputDate + inputTime
// 	log.Println(inputDateTime)
// 	//Parse date
// 	targetTime, _ := time.Parse("2006-01-02", inputDateTime)
// 	//calculate duration until target time
// 	// durationUntilTarget := targetTime.Sub(time.Now())
// 	durationUntilTarget := time.Until(targetTime)

// 	//Create timer
// 	timer := time.NewTimer(durationUntilTarget)
// 	// Event.time = time.Now().Add(time.Duration(timed) * time.Second)
// 	// Words.alarm = *time.NewTimer(time.Until(Words.time))

// 	// go func() {
// 	// 	defer wg.Done()
// 	// 	// fmt.Print(" time\n")
// 	// 	}()
// 	<-timer.C
// 	playSound()
// 	log.Println("Alert triggered")
// 	return "alert"
// }
