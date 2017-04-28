package set2

import (
	"crypto/rand"
	"fmt"

	"github.com/goCryptography/set1/functions"
)

// EncryptionOracle encrypts data, and states which block type it encrypted with
func EncryptionOracle(byteData []byte) string {
	blockSize := 16
	// byteKey := []byte("YELLOW SUBMARINE")
	byteKey := make([]byte, blockSize)

	// halfSize := *big.NewInt(5)
	// temp, err := rand.Int(rand.Reader, &halfSize)
	// if err != nil {
	// 	panic(err)
	// }
	temp := make([]byte, 1)
	_, err := rand.Read(temp)
	if err != nil {
		panic(err)
	}
	randBegSize := int(temp[0])/256 + 10

	temp = make([]byte, 1)
	_, err = rand.Read(temp)
	if err != nil {
		panic(err)
	}
	randEndSize := int(temp[0])/256 + 10

	begByte := make([]byte, randBegSize)
	_, err = rand.Read(begByte)
	if err != nil {
		panic(err)
	}
	endByte := make([]byte, randEndSize)
	_, err = rand.Read(endByte)
	if err != nil {
		panic(err)
	}

	byteData = append(begByte, byteData...)
	byteData = append(byteData, endByte...)

	_, err = rand.Read(byteKey)
	if err != nil {
		panic(err)
	}
	fmt.Printf("rand key: %v\n", byteKey)
	// fmt.Println(string(byteData))

	// add 5-10 random bytes to beginning and end of data
	_, err = rand.Read(byteKey)
	if err != nil {
		panic(err)
	}
	fmt.Printf("beg rand key: %v\n", byteKey)
	byteData = append(byteData)

	// encrypt by cbc
	// encryptedData := EncryptCbc(byteData, byteKey, []byte{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0})
	// encrypt by ecb
	encryptedData := set1.EncryptEcb(byteData, byteKey)

	// fmt.Println(encryptedData)
	// encryptedData = set1.DecryptEcb(encryptedData, byteKey)
	// fmt.Println(string(encryptedData))
	// send encrypted data, check whether it's ecb or cbc
	duplicates := set1.DetectAesEcbLine(encryptedData)

	if duplicates == "" {
		return "CBC"
	}
	return "ECB"

}
