package set2

import (
	"crypto/rand"
	"fmt"

	"github.com/goCryptography/set1/functions"
)

// EncryptionOracle encrypts data, and states which block type it encrypted with
func EncryptionOracle(byteData []byte) string {
	//
	// Create random bytes in the beginning and end, and create a random byte key
	//

	blockSize := 16

	fmt.Printf("data before additions: %v\n", byteData)

	// Create 1 random byte, convert it to an int between 5 and 10
	temp := make([]byte, 1)
	_, _ = rand.Read(temp)

	randBegSize := int((float64(temp[0])/256.0)*5) + 5
	fmt.Printf("beginning size: %d\n", randBegSize)

	temp = make([]byte, 1)
	_, _ = rand.Read(temp)

	randEndSize := int((float64(temp[0])/256.0)*5) + 5
	fmt.Printf("end size: %d\n", randEndSize)

	begByte := make([]byte, randBegSize)
	_, _ = rand.Read(begByte)
	fmt.Printf("beginning byte array: %v\n", begByte)

	endByte := make([]byte, randEndSize)
	_, _ = rand.Read(endByte)
	fmt.Printf("end byte array: %v\n", endByte)

	byteData = append(begByte, byteData...)
	byteData = append(byteData, endByte...)

	fmt.Printf("data after additions: %v\n", byteData)

	// Generate a random 16 byte byteKey
	byteKey := make([]byte, blockSize)
	_, _ = rand.Read(byteKey)

	fmt.Printf("rand key: %v\n", byteKey)

	//
	// Take byte key and byte data, and encrypt with either ecb or cbc
	//

	// encrypt by cbc
	// encryptedData := EncryptCbc(byteData, byteKey, []byte{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0})
	// encrypt by ecb
	encryptedData := set1.EncryptEcb(byteData, byteKey)

	// send encrypted data, check whether it's ecb or cbc
	duplicates := set1.DetectAesEcbLine(encryptedData)

	if duplicates == "" {
		return "CBC"
	}
	return "ECB"

}
