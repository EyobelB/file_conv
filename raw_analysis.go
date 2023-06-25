package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("Testing... 1, 2, 3....")
	raw_file, file_error := os.Open("src/francis_forever.flac")
	raw_info, file_info_error := os.Stat("src/francis_forever.flac")
	if file_error != nil {
		fmt.Println("File read error....")
	}
	if file_info_error != nil {
		fmt.Println("File info error....")
	}
	byte_array := make([]byte, raw_info.Size())
	len, byte_error := raw_file.ReadAt(byte_array, 0)
	if byte_error != nil {
		fmt.Println("Byte conversion error")
	}
	fmt.Println("Size: ", len, " bytes....")
	for i, byte_val := range byte_array {
		fmt.Printf("%08b\n", byte_val)
		if i == 3 {
			break
		}
	}
}
