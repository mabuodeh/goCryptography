package blockAndStreamCiphers

import (
	"fmt"
	"strings"
)

type tuple struct {
	key   []byte
	value []byte
}

func testEq(a, b []byte) bool {

	if a == nil && b == nil {
		return true
	}

	if a == nil || b == nil {
		return false
	}

	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

// BreakEcb encrypts data using ECB, detects blocksize, block type, and breaks it byte by byte
func BreakEcb(byteData []byte) {
	blockSize := 16
	// byteData := Pkcs7Padding(byteDataT, blockSize)
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
	_, noOfBlocks := Pkcs7Padding(byteData, blockSize)

	// send encrypted data, check whether it's ecb or cbc
	myStr := strings.Repeat("A", 2*finalSize)
	encryptedData = EncryptEcb(append([]byte(myStr), byteData...), byteKey)
	duplicates := DetectAesEcbLine(encryptedData)

	if duplicates == "" {
		fmt.Println("CBC")
	} else {
		fmt.Println("ECB")
	}

	// initialize finalText
	finalText := ""
	genLen := noOfBlocks * blockSize

	// loop for the blockSize to create A's i = [getLen:0]
	for i := genLen - 1; i >= 0; i-- {
		// initialize map to store getLen-byte key/value pairs
		combinations := make(map[int]tuple)
		// create a string with i A's
		myStr = strings.Repeat("A", i)
		// append finalText to string
		myStr += finalText
		// loop over j = [0:256] for the last byte.
		for j := 0; j < 256; j++ {
			dictStr := myStr
			dictStr += string(j)

			// append string to byteData before encrypting
			encryptedData = EncryptEcb(append([]byte(dictStr), byteData...), byteKey)
			// obtain and store the first getLen bytes of the encrypted data
			tempKey := make([]byte, genLen)
			copy(tempKey[:], encryptedData[:genLen])
			tempVal := make([]byte, genLen)
			copy(tempVal[:], []byte(dictStr))
			combinations[j] = tuple{key: tempKey, value: tempVal}
		} // endloop
		// append only i A's then encrypt
		dictStr := strings.Repeat("A", i)
		encryptedData = EncryptEcb(append([]byte(dictStr), byteData...), byteKey)
		// compare the encrypted bytes bytes with the stored data
		tempKey := make([]byte, genLen)
		copy(tempKey[:], encryptedData[:genLen])
		newVal := make([]byte, genLen)
		for _, kv := range combinations {
			if testEq(kv.key, tempKey) {
				newVal = kv.value
				break
			}
		}
		finalText = finalText + string(newVal[len(newVal)-1])
	}

	// Getting rid of additional bytes as a result of unequal blockSize
	finalLength := len(finalText) - len(byteData)
	finalText = finalText[:len(finalText)-finalLength]

	fmt.Printf("final text: %s\n\n", finalText)

}

// BreakPrependEcb encrypts data using ECB, detects blocksize, padding, block type, and breaks it byte by byte
func BreakPrependEcb(byteData []byte) {
	blockSize := 16
	// byteData := Pkcs7Padding(byteDataT, blockSize)
	byteKey, _ := getKeyAndIv(blockSize)

	//
	// Determine block size
	//
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

	//
	// send encrypted data, check whether it's ecb or cbc
	//
	myStr := strings.Repeat("A", 2*finalSize)
	encryptedData = EncryptEcb(append([]byte(myStr), byteData...), byteKey)
	duplicates := DetectAesEcbLine(encryptedData)

	if duplicates == "" {
		fmt.Println("CBC")
		panic("CBC!")
	} else {
		fmt.Println("ECB")
	}

	//
	// block size determined, and it's ecb. Time to break it
	//

	// initialize finalText
	finalText := ""

	_, noOfBlocks := Pkcs7Padding(byteData, blockSize)
	genLen := noOfBlocks * blockSize

	// loop for the blockSize to create A's i = [getLen:0]
	for i := genLen - 1; i >= 0; i-- {
		// initialize map to store getLen-byte key/value pairs
		combinations := make(map[int]tuple)
		// create a string with i A's
		myStr = strings.Repeat("A", i)
		// append finalText to string
		myStr += finalText
		// loop over j = [0:256] for the last byte.
		for j := 0; j < 256; j++ {
			dictStr := myStr
			dictStr += string(j)

			// append string to byteData before encrypting
			encryptedData = EncryptEcb(append([]byte(dictStr), byteData...), byteKey)
			// obtain and store the first getLen bytes of the encrypted data
			tempKey := make([]byte, genLen)
			copy(tempKey[:], encryptedData[:genLen])
			tempVal := make([]byte, genLen)
			copy(tempVal[:], []byte(dictStr))
			combinations[j] = tuple{key: tempKey, value: tempVal}
		} // endloop
		// append only i A's then encrypt
		dictStr := strings.Repeat("A", i)
		encryptedData = EncryptEcb(append([]byte(dictStr), byteData...), byteKey)
		// compare the encrypted bytes bytes with the stored data
		tempKey := make([]byte, genLen)
		copy(tempKey[:], encryptedData[:genLen])
		newVal := make([]byte, genLen)
		for _, kv := range combinations {
			if testEq(kv.key, tempKey) {
				newVal = kv.value
				break
			}
		}
		finalText = finalText + string(newVal[len(newVal)-1])
	}

	// Getting rid of additional bytes as a result of unequal blockSize
	finalLength := len(finalText) - len(byteData)
	finalText = finalText[:len(finalText)-finalLength]

	fmt.Printf("final text: %s\n\n", finalText)

}
