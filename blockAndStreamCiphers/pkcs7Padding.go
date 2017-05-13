package blockAndStreamCiphers

import "bytes"

// Pkcs7Padding takes a byte slice of data and the total size of a block and appends empty bytes to the end
func Pkcs7Padding(data []byte, blockSize int) ([]byte, int) {
	// blockSize := 16
	paddedSize := blockSize - len(data)%blockSize
	// fmt.Printf("padding: %d\n", paddedSize)
	paddedPart := bytes.Repeat([]byte{byte(paddedSize)}, paddedSize)
	// fmt.Println(paddedPart)

	finalData := append(data, paddedPart...)
	// return append(data, hex.EncodeToString(paddedPart)...)
	return finalData, len(finalData) / 16
}

// Pkcs7Unpadding removes padding
func Pkcs7Unpadding(data []byte) []byte {
	// location of last val
	loc := 0
	// start with last value
	lastVal := data[len(data)-1]
	// iterate in reverse until a different value is reached
	for i := len(data) - 2; i >= 0; i-- {
		if data[i] != lastVal {
			loc = i + 1
			break
		}
	}
	return data[:loc]
}
