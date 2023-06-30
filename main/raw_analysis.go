package main

import (
	"fmt"
	"os"
)

func raw_analysis(raw_file os.File, file_size int64) {
	fmt.Println("Initiating byte analysis....")

	// raw_file, file_error := os.Open("src/francis_forever.flac")
	// raw_info, file_info_error := os.Stat()
	// if file_error != nil {
	// 	fmt.Println("File read error....")
	// }
	// if file_info_error != nil {
	// 	fmt.Println("File info error....")
	// }
	byte_array := make([]byte, file_size)
	len, byte_error := raw_file.ReadAt(byte_array, 0)
	if byte_error != nil {
		fmt.Println("Byte conversion error")
	}
	fmt.Println("Size:", len, "bytes....")
	for i, byte_val := range byte_array {
		if i == 4 {
			break
		}
		fmt.Printf("%08b : %s\n", byte_val, string(byte_val))
	}
}
