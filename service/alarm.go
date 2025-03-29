package service

import (
	"encoding/json"
	"log"
	"os"
	"time"

	"Focus/structs"
)

func GetEvents() ([]structs.Period, error) {
	file, err := os.ReadFile(structs.UserSchedule) // Adjust the path as necessary
	if err != nil {
		return nil, err
	}
	var existingData []structs.Period
	if err := json.Unmarshal(file, &existingData); err != nil {
		return nil, err
	}
	return existingData, nil
}

func StartAlarm() {
	for {
		events, err := GetEvents()
		if err != nil {
			log.Println("Error reading events:", err)
			return
		}

		now := time.Now()
		var soonestAlarm *structs.Period
		// var timeBefore time.Time
		for i := range events {
			eventTime, err := time.Parse("2006-01-02 15:04", events[i].Date+" "+events[i].Time)
			if err != nil {
				log.Println("Error parsing event time:", err)
				continue
			}
			// timeBefore=getEventTime(soonestAlarm)

			if eventTime.After(now) { //&&eventTime.Before(timeBefore)

				soonestAlarm = &events[i]

				if eventTime.Before(getEventTime(soonestAlarm)) {
					soonestAlarm = &events[i]
				}
			}

			if soonestAlarm != nil {
				timeBefore, err := time.Parse("2006-01-02 15:04", soonestAlarm.Date+" "+soonestAlarm.Time)

				timeBefore = timeBefore.Add(-3 * time.Hour)
				if err != nil {
					log.Println("Error parsing event time:", err)
					continue
				}
				sleepDuration := time.Until(timeBefore)
				time.Sleep(sleepDuration)

				// Trigger the alarm
				playSound()
				log.Println("Alert triggered at:", soonestAlarm.Time)

			} else {
				// If there are no upcoming alarms, check every minute
				time.Sleep(3 * time.Second)
			}
		}
	}
}
func getEventTime(event *structs.Period) time.Time {
	eventTime, err := time.Parse("2006-01-02 15:04", event.Date+" "+event.Time)
	if err != nil {
		log.Println("Error parsing event time:", err)
		return time.Time{} // Return zero Time if there's an error
	}
	return eventTime
}
