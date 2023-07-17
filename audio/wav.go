package audio

import (
	"fmt"
	"os"
	"strconv"
)

type fmt_chunk struct {
	marker           string
	chunk_len        int
	waveform_format  int
	channel_count    int
	sample_rate      int
	byte_rate        int
	bytes_per_sample int
	bits_per_sample  int
}

type data_chunk struct {
	marker   string
	data_len int
	data     []uint8
}

type wav struct {
	audio     os.File
	wav_len   int
	full_file []uint8
	format    fmt_chunk
	data      data_chunk
	isRIFF    bool
	isWAV     bool
}

func NewFmt(subframe []uint8) *fmt_chunk {
	// Create an internal file seeker
	file_seek := 0

	// Grab the initial "fmt " letters (including the space)
	fmt_string := fmt.Sprintf("%s", subframe[file_seek:file_seek+4])
	fmt_obj := fmt_chunk{marker: fmt_string}

	// Error check for initial letters
	if fmt_string != "fmt " {
		panic("Incorrect format marker")
	} else {
		file_seek += 4
	}

	// Grab the subchunk size, which is the size after this value
	var conv_err error
	fmt_obj.chunk_len, conv_err = strconv.Atoi(fmt.Sprintf("%d", subframe[file_seek:file_seek+4]))
	if conv_err != nil {
		panic("Error converting chunk to integer")
	}
	file_seek += 4

	// Determine audio format
	fmt_obj.waveform_format, conv_err = strconv.Atoi(fmt.Sprintf("%d", subframe[file_seek:file_seek+4]))
	if conv_err != nil {
		panic("Error converting chunk to integer")
	}
	file_seek += 4

	// Error check for non PCM values
	if fmt_obj.chunk_len != 16 || fmt_obj.waveform_format != 1 {
		panic("Non PCM values detected....")
	}

	// Determine number of audio channels
	fmt_obj.channel_count, conv_err = strconv.Atoi(fmt.Sprintf("%d", subframe[file_seek:file_seek+2]))
	if conv_err != nil {
		panic("Error converting chunk to integer")
	}
	file_seek += 2

	// Initialize sample rate
	fmt_obj.sample_rate, conv_err = strconv.Atoi(fmt.Sprintf("%d", subframe[file_seek:file_seek+4]))
	if conv_err != nil {
		panic("Error converting value to integer")
	}
	file_seek += 4

	// Grab the byte rate, AKA sample rate * channel count * bytesPerSample
	fmt_obj.byte_rate, conv_err = strconv.Atoi(fmt.Sprintf("%d", subframe[file_seek:file_seek+4]))
	if conv_err != nil {
		panic("Error converting chunk to integer")
	}
	file_seek += 4

	// Get the number of bytes in a sample across all channels
	fmt_obj.bytes_per_sample, conv_err = strconv.Atoi(fmt.Sprintf("%d", subframe[file_seek:file_seek+2]))
	if conv_err != nil {
		panic("Error converting chunk to integer")
	}
	file_seek += 2

	// Indicate bits per sample
	fmt_obj.bits_per_sample, conv_err = strconv.Atoi(fmt.Sprintf("%d", subframe[file_seek:file_seek+2]))
	if conv_err != nil {
		panic("Error converting chunk to integer")
	}
	file_seek += 2

	// Return address of format chunk
	return &fmt_obj
}

func NewData(subframe []uint8) *data_chunk {
	// Create an internal file seeker
	file_seek := 0

	// Grab the initial "fmt " letters (including the space)
	data_string := fmt.Sprintf("%s", subframe[file_seek:file_seek+4])
	data_obj := data_chunk{marker: data_string}
	file_seek += 4

	// Verify the validity of the data subframe
	if data_obj.marker != "data" {
		panic("Incorrect data marker....")
	}

	// Determine the size of the data subframe
	var conv_err error
	data_obj.data_len, conv_err = strconv.Atoi(fmt.Sprintf("%d", subframe[file_seek:file_seek+4]))
	if conv_err != nil {
		panic("Error converting chunk to integer")
	}
	file_seek += 4

	// Loop across the range of bytes to store each of the PCM data values as uint8
	var data_array []uint8
	for i := file_seek; i < data_obj.data_len; i += 4 {
		var byte_val int
		byte_val, conv_err = strconv.Atoi(fmt.Sprintf("%b", subframe[i:i+4]))
		if conv_err != nil {
			panic("Bad data byte type conversion....")
		}
		data_array = append(data_array, uint8(byte_val))
	}
	data_obj.data = data_array

	// Return data chunk address
	return &data_obj
}

func NewWav(audio os.File) *wav {
	// Load in the audio file
	wav_obj := wav{audio: audio}

	// Load file as a byte array
	var read_error error
	wav_obj.full_file, read_error = os.ReadFile(audio.Name())
	if read_error != nil {
		panic("Failed to read audio file....")
	}

	// Create file seek int to keep track of current byte
	file_seek := 0

	// Create initialization values for the file
	riff_string := fmt.Sprintf("%s", wav_obj.full_file[file_seek:file_seek+4])
	if riff_string != "RIFF" {
		wav_obj.isRIFF = false
		panic("File not a subset of the .RIFF format....")
	} else {
		wav_obj.isRIFF = true
		file_seek += 4
	}

	// Determine length of the wav file
	var conv_err error
	wav_obj.wav_len, conv_err = strconv.Atoi(fmt.Sprintf("%d", wav_obj.full_file[file_seek:file_seek+4]))
	if conv_err != nil {
		panic("Error converting value to integer")
	}
	file_seek += 4

	// Grab the WAVE subcategory string
	wave_string := fmt.Sprintf("%s", wav_obj.full_file[file_seek:file_seek+4])
	if wave_string != "WAVE" {
		wav_obj.isWAV = false
		panic("File not a .WAV file....")
	} else {
		wav_obj.isWAV = true
		file_seek += 4
	}

	// Create the format chunk, starting with the "fmt" text
	format := NewFmt(wav_obj.full_file[file_seek:len(wav_obj.full_file)])
	wav_obj.format = *format
	file_seek += 24

	// Create the data chunk
	data := NewData(wav_obj.full_file[file_seek:len(wav_obj.full_file)])
	wav_obj.data = *data

	// Return pointer to wav object
	return &wav_obj
}
