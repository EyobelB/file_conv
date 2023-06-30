package main

// This file contains calculators for various checksums used in various audio formats

func CRC_16_Gen(send int) int {
	// Function here is x^16 + x^15 + x^2 + x^0 => 11000000 00000010 1
	generator := 98309

	// Add 16 trailing zeros to the send value
	send_check := send * 65536

	// Calculate the remainder of division, return checksum
	return send_check % generator
}

func CRC_8_Gen(send int) int {
	// Function here is x^8 + x^2 + x^1 + x^0 => 10000011 1
	generator := 263

	// Add 8 trailing zeros to the send value
	send_check := send * 256

	// Calculate the remainder of division
	return send_check % generator
}

func CRC_16_Check(recv int, recv_checksum int) bool {
	// Function here is x^16 + x^15 + x^2 + x^0 => 11000000 00000010 1
	generator := 98309

	// Add 16 trailing zeros to the received value
	recv_test := recv * 65536

	// Calculate the remainder of division, return checksum
	return recv_test%generator == recv_checksum
}

func CRC_8_Check(recv int, recv_checksum int) bool {
	// Function here is x^8 + x^2 + x^1 + x^0 => 10000011 1
	generator := 263

	// Add 8 trailing zeros to the received value
	recv_test := recv * 256

	// Calculate the remainder of division
	return recv_test%generator == recv_checksum
}
