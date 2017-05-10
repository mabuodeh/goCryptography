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

	// blocks := [16]byte{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}

	// initialize finalText
	finalText := ""
	// loop for the blockSize to create A's i = [15:0]
	for i := 15; i >= 0; i-- {
		// i := 15
		// {
		// initialize map to store 16-byte values
		combinations := make(map[[16]byte][16]byte)
		// create a string with i A's
		myStr = strings.Repeat("A", i)
		// append finalText to string
		myStr += string(finalText)
		// loop over j = [0:256] for the 8th byte.
		for j := 32; j < 127; j++ {
			// j := 1
			// {
			// append j to string
			// myStr += string(j)
			dictStr := myStr
			dictStr += string(j)
			// fmt.Println(string(dictStr))
			// append string to byteData before encrypting
			encryptedData = EncryptEcb(append([]byte(dictStr), byteData...), byteKey)
			// obtain and store the first 16 bytes of the encrypted data
			tempKey := [16]byte{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
			copy(tempKey[:], encryptedData[0:16])
			tempVal := [16]byte{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
			copy(tempVal[:], []byte(dictStr))

			combinations[tempKey] = tempVal
			// fmt.Println(tempVal)

			// endloop
		}
		// append only i A's then encrypt
		dictStr := strings.Repeat("A", i)
		encryptedData = EncryptEcb(append([]byte(dictStr), byteData...), byteKey)
		// compare the 16 bytes with the stored data
		tempKey := [16]byte{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
		copy(tempKey[:], encryptedData[0:16])
		// fmt.Println(string(tempKey[:]))
		newVal := combinations[tempKey]
		// fmt.Println(newVal)
		// preappend last byte to finalText
		finalText = finalText + string(newVal[len(newVal)-1])
		fmt.Println(finalText)
	}
	fmt.Printf("final text: %s\n", finalText)

}
