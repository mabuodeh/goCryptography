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
		fmt.Println("one nil")
		return false
	}

	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			// fmt.Printf("%v %v", a[i], b[i])

			// fmt.Println("not equal")
			return false
		}
	}
	return true
}

// BreakEcb (ONLY BREAKS FIRST BLOCK) encrypts data using ECB, detects blocksize, block type, and breaks it byte by byte
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

	// blocks := [16]byte{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}

	// initialize finalText
	finalText := ""
	finalEcbSen := ""

	for blk := 1; blk <= noOfBlocks; blk++ {
		// blk := 2
		// {
		genLen := blk * blockSize
		fmt.Println(finalText)
		finalEcbSen += finalText
		finalText = ""

		// loop for the blockSize to create A's i = [15:0]
		// CHANGE 15 to len(msg) - 1
		for i := genLen - 1; i >= 0; i-- {
			// i := len(byteData) - 1
			// {
			// initialize map to store 16-byte values
			// CHANGE to store i sized values
			combinations := make(map[int]tuple)

			myStr = ""
			// myStr += finalText
			// create a string with i A's
			myStr += strings.Repeat("A", i)
			// append finalText to string
			myStr += finalText
			// loop over j = [0:256] for the 8th byte.
			// for j := 32; j < 127; j++ {
			for j := 0; j < 256; j++ {
				// j := 1
				// {
				// append j to string
				// myStr += string(j)
				dictStr := myStr
				dictStr += string(j)

				// append string to byteData before encrypting
				encryptedData = EncryptEcb(append([]byte(dictStr), byteData...), byteKey)
				// obtain and store the first 16 bytes of the encrypted data
				// CHANGE to store len(msg) values instead of 16
				tempKey := make([]byte, genLen)
				copy(tempKey[:], encryptedData[:genLen])
				tempVal := make([]byte, genLen)
				copy(tempVal[:], []byte(dictStr))

				// fmt.Println(string(tempVal))
				// combinations[string(tempKey)] = tempVal
				combinations[j] = tuple{key: tempKey, value: tempVal}
				// fmt.Println(string(tempKey))

				// endloop
			}
			// append only i A's then encrypt
			dictStr := strings.Repeat("A", i)
			encryptedData = EncryptEcb(append([]byte(dictStr), byteData...), byteKey)
			// compare the 16 bytes with the stored data
			// CHANGE 16 to len(msg)
			tempKey := make([]byte, genLen)
			copy(tempKey[:], encryptedData[:genLen])
			// fmt.Println(tempKey)
			// fmt.Println(len(combinations))
			// newVal := combinations[string(tempKey)]
			newVal := make([]byte, genLen)
			for _, kv := range combinations {
				// fmt.Println(kv.key)
				// fmt.Println(kv.value)
				if testEq(kv.key, tempKey) {
					// fmt.Println("true")
					// fmt.Printf("equal! %v\n", kv.value)
					newVal = kv.value
					break
				}
			}
			// fmt.Println(string(newVal))
			// fmt.Println(len(tempKey))
			// fmt.Println(newVal)
			// preappend last byte to finalText
			finalText = finalText + string(newVal[len(newVal)-1])
			// fmt.Println(finalText)
		}
	}
	fmt.Printf("final text: %s\n\n", finalEcbSen)

}
