package main

import (
	audio "audio_conv"
	"fmt"
	"os"
)

func main() {
	// Create file object for the song
	francis, open_err := os.Open("src/francis_forever.flac")
	if open_err != nil {
		fmt.Println("Error: File opening failed....")
	}

	// Import "Francis Forever" by Mitski as a .flac object
	francis_flac := audio.NewFlac(*francis)

	// Continue with conversion from here
	// For now, just print the magic number....
	fmt.Printf("Magic number: %s\n", francis_flac.MagicNumber())
	fmt.Println("Raw byte data: ", francis_flac.MagicNumber())
	fmt.Printf("File is a valid .flac? %t\n", francis_flac.ValidFlac())
	// fmt.Println("Metadata incoming....")
	// francis_flac.PrintMetadata()
}
