package blockAndStreamCiphers

import "bytes"

// Pkcs7Padding takes a byte slice of data and the total size of a block and appends empty bytes to the end
func Pkcs7Padding(data []byte, blockSize int) []byte {
	// blockSize := 16
	paddedSize := blockSize - len(data)%blockSize
	// fmt.Printf("padding: %d\n", paddedSize)
	paddedPart := bytes.Repeat([]byte{byte(paddedSize)}, paddedSize)
	// fmt.Println(paddedPart)

	// return append(data, hex.EncodeToString(paddedPart)...)
	return append(data, paddedPart...)
}
