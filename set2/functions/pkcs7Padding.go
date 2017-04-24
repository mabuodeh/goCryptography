package set2

import (
	"bytes"
	"encoding/hex"
	"fmt"
)

// Pkcs7Padding takes a byte slice of data and the total size of a block and appends empty bytes to the end
func Pkcs7Padding(data []byte, blockSize int) []byte {
	// blockSize := 16
	paddedSize := blockSize - len(data)%blockSize
	fmt.Println(paddedSize)
	paddedPart := bytes.Repeat([]byte{byte(paddedSize)}, paddedSize)

	return append(data, hex.EncodeToString(paddedPart)...)
}
