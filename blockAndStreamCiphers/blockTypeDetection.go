package blockAndStreamCiphers

import (
	"crypto/rand"
	"fmt"
)

var consistentKey = make([]byte, 16)

// EncryptionOracle encrypts data, and states which block type it encrypted with
func EncryptionOracle(byteData []byte) string {
	//
	// Create random bytes in the beginning and end, and create a random byte key
	//

	blockSize := 16

	min := 5
	max := 10
	begBytes, endBytes := randPadding(min, max)

	fmt.Printf("beginning byte array: %v\n", begBytes)
	fmt.Printf("end byte array: %v\n", endBytes)

	// append random bytes to the beginning and end of data
	byteData = append(begBytes, byteData...)
	byteData = append(byteData, endBytes...)

	// fmt.Printf("data after additions: %v\n", byteData)
	byteKey, iv := getKeyAndIv(blockSize)

	encryptedData := randEncrypt(byteData, byteKey, iv)

	// send encrypted data, check whether it's ecb or cbc
	duplicates := DetectAesEcbLine(encryptedData)

	if duplicates == "" {
		return "CBC"
	}
	return "ECB"

}

func getKeyAndIv(blockSize int) ([]byte, []byte) {
	randKeyAndIv := make([]byte, 2*blockSize)
	_, _ = rand.Read(randKeyAndIv)

	// Generate a random 16 byte byteKey
	byteKey := randKeyAndIv[:blockSize]
	fmt.Printf("rand key: %v\n", byteKey)

	// Generate a random 16 byte IV
	iv := randKeyAndIv[blockSize:]
	fmt.Printf("rand iv: %v\n", iv)

	return byteKey, iv
}

// random padding, returns bytes to pad before and after data. Size veries depending on start and end ints given
func randPadding(start, end int) (begByte, endByte []byte) {
	// Create 1 random byte, convert it to an int between 5 and 10
	temp := make([]byte, 1)
	_, _ = rand.Read(temp)

	// randBegSize := int((float64(temp[0])/256.0)*5) + 5
	randBegSize := int((float64(temp[0])/256.0)*float64(start)) + (end - start)

	temp = make([]byte, 1)
	_, _ = rand.Read(temp)

	randEndSize := int((float64(temp[0])/256.0)*float64(start)) + (end - start)

	// fmt.Printf("beginning size: %d\n", randBegSize)
	// fmt.Printf("end size: %d\n", randEndSize)

	begByte = make([]byte, randBegSize)
	_, _ = rand.Read(begByte)

	endByte = make([]byte, randEndSize)
	_, _ = rand.Read(endByte)

	return begByte, endByte
}

func randEncrypt(byteData, byteKey, iv []byte) []byte {

	//
	// choose ecb or cbc randomly
	//
	// gets a random byte, converts it to an int which will be between 5 and 10.
	temp := make([]byte, 1)
	_, _ = rand.Read(temp)

	blockType := int((float64(temp[0])/256.0)*5) + 5
	// fmt.Println(blockType)

	// if random int is between 5 and 7, encrypt with ECB, if between 8 and 10 then CBC
	encryptedData := make([]byte, 0)
	if blockType <= 7 {
		fmt.Println("ecb")
		encryptedData = EncryptEcb(byteData, byteKey)
	} else {
		fmt.Println("cbc")
		encryptedData = EncryptCbc(byteData, byteKey, iv)
	}

	return encryptedData

}
