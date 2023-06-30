package main

import (
	audio "audio_conv"
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	// Create file object for the song
	current_path, path_err := os.Getwd()
	if path_err != nil {
		panic(path_err)
	}

	song_path := filepath.Join(current_path, "..", "src", "francis_forever.flac")
	francis, open_err := os.Open(song_path)
	if open_err != nil {
		panic(open_err)
	}

	// Import "Francis Forever" by Mitski as a .flac object
	francis_flac := audio.NewFlac(*francis)
	francis_data, data_err := os.Stat(song_path)
	if data_err != nil {
		panic(data_err)
	}

	// Continue with conversion from here
	// For now, just print the magic number....
	fmt.Println("Magic number: ", francis_flac.MagicNumber())
	fmt.Printf("File is a valid .flac? %t\n", francis_flac.ValidFlac())
	fmt.Println("CRC-8 test.... CRC-8(16): ", CRC_8_Check(16, CRC_8_Gen(642)))
	fmt.Println("CRC-16 test.... CRC-16(25): ", CRC_16_Check(25, CRC_16_Gen(25)))
	raw_analysis(*francis, francis_data.Size())
	// fmt.Println("Metadata incoming....")
	// francis_flac.PrintMetadata()
}
