package audio

import (
	"fmt"
	"os"
	"strconv"
)

type metadata struct {
	meta_type uint8
	data_len  int
	data      []uint8
	isVorbis  bool
	isLast    bool
}

type frame struct {
	full_frame  []uint8
	header      []uint8
	data        []uint8
	checksum    []uint8
	valid_frame bool
}

type flac struct {
	audio           os.File
	full_file       []uint8
	magic_number    []uint8
	metadata_blocks []metadata
	frames          []frame
	isFlac          bool
}

func NewFrame(frame_segment []uint8) *frame {
	// Load in the frame
	frame := frame{full_frame: frame_segment}

	// Define the header segment
	frame.header = frame_segment[0:4]

	// Define the data segment by converting to string and then to int
	string_form := fmt.Sprintf("%d", frame.header[1:4])
	data_size, data_error := strconv.Atoi(string_form)
	if data_error != nil {
		data_size = 4
	}
	frame.data = frame_segment[4 : data_size-1]
	// Define the CRC-16 polynomial checksum, and check for validity
	// Use checksum validity to determine valid_frame variable

	// return the completed Frame
	return &frame
}

func NewMetadata(metadata_segment []uint8) *metadata {
	// Load in metadata
	metadata := metadata{meta_type: metadata_segment[0]}
	fmt.Println(metadata_segment[1:3])
	data_len_string := fmt.Sprintf("%d", metadata_segment[1:3])
	data_len_val, data_len_err := strconv.Atoi(data_len_string)
	if data_len_err != nil {
		panic(data_len_err)
	}
	metadata.data_len = data_len_val
	metadata.data = metadata_segment[4 : len(metadata_segment)-1]

	// Set up standard metadata blocks (1st byte: MSB == 1 means last metadata block (others are type of metadata), 2nd-4th byte: indicates size of metadata block)
	// Set up Vorbis Comment block/s (1st byte: MSB == 1 means last metadata block (if remaining bits are 4, Vorbis comment), 2nd-4th byte: indicates size of metadata block)
	metadata.isVorbis = (metadata.meta_type == uint8(4))
	metadata.isLast = (metadata.meta_type >= uint8(128))
	return &metadata
}

func NewFlac(audio os.File) *flac {
	// Load in the audio file
	flac := flac{audio: audio}

	// Load in the file as a byte array
	var err error
	flac.full_file, err = os.ReadFile(audio.Name())
	if err != nil {
		fmt.Println("Failed to read audio file")
	}

	// Create the magic number
	flac.magic_number = flac.full_file[0:4]

	// Determine validity as a .flac file
	if fmt.Sprintf("%s", flac.magic_number) == "fLaC" {
		flac.isFlac = true
	} else {
		flac.isFlac = false
	}

	// // Define metadata blocks
	// isHeader := 0
	// bytes_remaining := 0
	// var metadata_bytes []uint8
	// for byte_index, bytes := range flac.full_file {
	// 	if byte_index < 4 {
	// 		continue
	// 	}
	// 	// Read one byte at a time to construct metadata blocks
	// 	if isHeader < 4 {
	// 		metadata_bytes = append(metadata_bytes, bytes)
	// 		isHeader++
	// 	}

	// 	// If we hit 4 bytes, reset the metadata. If this is the last block, exit
	// 	if isHeader == 4 {
	// 		// Save the data, isVorbis, and isLast parts of metadata
	// 		fmt.Println(metadata_bytes[3])
	// 		data_len_string := fmt.Sprintf("%d%d%d", metadata_bytes[1], metadata_bytes[2], metadata_bytes[3])
	// 		fmt.Println(data_len_string)
	// 		data_len_val, data_len_err := strconv.Atoi(data_len_string)
	// 		if data_len_err != nil {
	// 			panic(data_len_err)
	// 		}
	// 		bytes_remaining = data_len_val
	// 	}

	// 	if isHeader == 4 && bytes_remaining != 0 {
	// 		metadata_bytes = append(metadata_bytes, bytes)
	// 		bytes_remaining--
	// 	}
	// 	// Create block
	// 	if bytes_remaining == 0 && isHeader == 4 {
	// 		fmt.Println(metadata_bytes)
	// 		metadata_block := NewMetadata(metadata_bytes)
	// 		flac.metadata_blocks = append(flac.metadata_blocks, *metadata_block)
	// 		if metadata_block.isLast {
	// 			break
	// 		}
	// 		metadata_bytes = nil
	// 		isHeader = 0
	// 	}

	// }
	// // Indicate the beginning of stream blocks
	// var non_stream_blocks int
	// non_stream_blocks = 4 + len(flac.metadata_blocks)

	// for byte_index, bytes := range flac.full_file {
	// 	// Use the remaining stream to fill the frame array
	// 	if byte_index < non_stream_blocks {
	// 		continue
	// 	}
	// 	// Create frames until EOF

	// 	// Append frame to the frames array

	// }

	// Return the fully defined .flac object
	return &flac
}

// Necessary getter functions for all flac structs
func (m metadata) VorbisComment() {
	// If this is a Vorbis Comment, print the string involved
	if m.isVorbis {
		fmt.Printf("%s\n", m.data)
	}
}

func (f flac) MagicNumber() string {
	return fmt.Sprintf("%s", f.magic_number)
}

func (f flac) ValidFlac() bool {
	return f.isFlac
}

func (f flac) PrintMetadata() {
	for _, block := range f.metadata_blocks {
		fmt.Printf("%s\n\n", block.data)
	}
}
