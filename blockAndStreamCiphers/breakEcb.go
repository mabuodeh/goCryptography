package blockAndStreamCiphers

import (
	"fmt"
	"strings"
)

// BreakEcb encrypts data using ECB, detects blocksize, block type, and breaks it byte by byte
func BreakEcb(byteData []byte) {
	blockSize := 16
	byteKey, _ := getKeyAndIv(blockSize)

	// initially encrypt byteData and store the ciphertext length
	fmt.Println("encrypting..")
	encryptedData := EncryptEcb(byteData, byteKey)
	findSize := len(encryptedData)
	finalSize := 0
	// loop until blockSize found
	for i := 1; ; i++ {
		// Create a custom string (A's). Number of A's is based on loop i
		myStr := strings.Repeat("A", i)
		// append byteData to custom string
		plaintext := append([]byte(myStr), byteData...)
		// encrypt resultant plaintext
		encryptedData = EncryptEcb(plaintext, byteKey)
		// check ciphertext size
		// if ciphertext size changed
		if findSize != len(encryptedData) {
			// take difference. that's the blockSize
			finalSize = (len(encryptedData) - findSize)
			fmt.Printf("main size: %d, encrypted size: %d, block size: %d\n", findSize, len(encryptedData), finalSize)
			// break
			break
		}
	}

	// send encrypted data, check whether it's ecb or cbc
	myStr := strings.Repeat("A", 2*finalSize)
	encryptedData = EncryptEcb(append([]byte(myStr), byteData...), byteKey)
	duplicates := DetectAesEcbLine(encryptedData)

	if duplicates == "" {
		fmt.Println("CBC")
	} else {
		fmt.Println("ECB")
	}
}
