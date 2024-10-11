package service

import (
	"fmt"
	"os"
	"time"

	"github.com/faiface/beep/speaker"
	"github.com/faiface/beep/wav"
	// "github.com/gen2brain/beeep"
)

func playSound() {
	// Load sound file for the alert (change file path according to your system)
	f, err := os.Open("./service/Ring08.wav") // Ensure you have an alert.wav file in your project directory
	if err != nil {
		fmt.Println("Error opening sound file:", err)
		return
	}
	defer f.Close()

	streamer, format, err := wav.Decode(f)
	if err != nil {
		fmt.Println("Error decoding sound file:", err)
		return
	}
	defer streamer.Close()

	// Initialize the speaker
	if err := speaker.Init(format.SampleRate, format.SampleRate.N(time.Second/10)); err != nil {
		fmt.Println("Error initializing speaker:", err)
		return
	}

	speaker.Play(streamer)

	// Play sound for 5 seconds
	time.Sleep(4 * time.Second)

	// Stop the sound
	speaker.Close()
}

// func showPopup(message string) {
// 	// Display a simple notification using beeep
// 	err := beeep.Alert("Alert", message, "DuskyOwls_EN-US9845705930_UHD.png") // Replace with your image path if necessary
// 	if err != nil {
// 		fmt.Println("Error showing popup:", err)
// 	}
// }

// func setAlarm(duration time.Duration) {
// 	time.AfterFunc(duration, func() {
// 		// When the alarm goes off, play sound and show notification
// 		playSound()         // Play the alert sound
// 		// showPopup("Time's up!") // Show the popup notification
// 	})
// }

// func main() {
// 	var seconds int
// 	fmt.Print("Enter alarm time in seconds: ")
// 	_, err := fmt.Scan(&seconds)
// 	if err != nil || seconds <= 0 {
// 		fmt.Println("Invalid input. Please enter a positive number.")
// 		return
// 	}

// 	setAlarm(time.Duration(seconds) * time.Second)

// 	// Using a loop to keep the program running until it is manually stopped
// 	for {
// 		time.Sleep(1 * time.Second) // Sleep for a short time to prevent busy waiting
// 	}
// }
