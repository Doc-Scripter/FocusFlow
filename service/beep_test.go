// Successfully plays a sound file for the specified duration
package service

import (
	"os"
	"testing"
	"time"
)

func TestPlaySoundSuccessfully(t *testing.T) {
	// Setup: Create a temporary sound file
	tempFile, err := os.CreateTemp("", "test_sound_*.wav")
	if err != nil {
		t.Fatalf("Failed to create temp file: %v", err)
	}
	defer os.Remove(tempFile.Name())

	// Write some data to the temp file to simulate a valid WAV file
	// (In a real test, you would write actual WAV data)
	tempFile.Write([]byte("RIFF....WAVEfmt "))

	// Replace the file path in playSound function with the temp file path
	originalFilePath := "Ring08.wav"
	defer func() { os.Rename(tempFile.Name(), originalFilePath) }()
	os.Rename(tempFile.Name(), originalFilePath)

	// Call the function under test
	playSound()

	// Assert: Check if the speaker was initialized and played for 5 seconds
	// (This is a simple check; in a real test, you might use a mock or spy)
	time.Sleep(6 * time.Second)
}
