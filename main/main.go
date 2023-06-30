package main

import (
	audio "audio_conv"
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	// If user inputs no arguments, quit the program and require that arguments are submitted
	if len(os.Args) < 2 {
		panic("Invalid arguments, run with -h or --help option for help text")
	}
	// Use arguments to determine input and output path
	input_file := ""
	output_file := ""

	// Check if help text is requested
	if os.Args[1] == "-h" || os.Args[1] == "--help" {
		help_text := "This program is a file conversion utility that is meant to one day\n"
		help_text += "    allow for the conversion any two files occupying the same medium.\n"
		help_text += "    Current options include the following:\n\n"
		help_text += "-i: Input file to be converted (make sure to enter absolute or relative path to execution location)\n"
		help_text += "-o: Output file to be converted (make sure to enter absolute or relative path to execution location)\n"
		help_text += "-h/--help: Prints out the help text displayed here"
		panic(help_text)
	}

	// Create logic to determine input and output files
	if os.Args[1] == "-i" {
		input_file = os.Args[2]
	}
	if os.Args[3] == "-o" {
		output_file = os.Args[4]
	}

	// Create file object for the song
	current_path, path_err := os.Getwd()
	if path_err != nil {
		panic(path_err)
	}

	song_path := filepath.Join(current_path, input_file)
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
	fmt.Println("File is a valid .flac?", francis_flac.ValidFlac())
	fmt.Println("CRC-8 test.... CRC-8(16): ", CRC_8_Check(16, CRC_8_Gen(642)))
	fmt.Println("CRC-16 test.... CRC-16(25): ", CRC_16_Check(25, CRC_16_Gen(25)))
	fmt.Println("Output file path:", filepath.Join(current_path, output_file))
	raw_analysis(*francis, francis_data.Size())
	// fmt.Println("Metadata incoming....")
	// francis_flac.PrintMetadata()
}
