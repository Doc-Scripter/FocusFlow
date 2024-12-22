package service

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"sort"
	"time"

	"Focus/structs"
)

func GetEvents() ([]structs.Period, error) {
	file, err := os.ReadFile(structs.UserSchedule)
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

		// var soonestAlarm *structs.Period
		// var timeBefore time.Time
		now := time.Now()
		EventTimes := []time.Time{}
		for i := range events {
			eventTime, err := time.Parse("2006-01-02 15:04", events[i].Date+" "+events[i].Time)
			if err != nil {
				log.Println("Error parsing event time:", err)
				continue
			}
			localEventTime := eventTime.In(time.Local)
			localEventTime=localEventTime.Add(-3 * time.Hour)
			if localEventTime.After(now) {
				fmt.Println(localEventTime)

				EventTimes = append(EventTimes, localEventTime)
			}
		}
		sort.Slice(EventTimes, func(i, j int) bool {
			return EventTimes[i].Before(EventTimes[j])
		})
		
		if len(EventTimes) > 1 {
			time.Sleep(time.Until(EventTimes[0]))
			playSound()
		}
		time.Sleep(3 * time.Second)
	}

}

// func getEventTime(event *structs.Period) time.Time {
// 	eventTime, err := time.Parse("2006-01-02 15:04", event.Date+" "+event.Time)
// 	if err != nil {
// 		log.Println("Error parsing event time:", err)
// 		return time.Time{} // Return zero Time if there's an error
// 	}
// 	return eventTime
// }
