package blockAndStreamCiphers

import (
	"crypto/aes"
	"fmt"
)

// EncryptEcb takes a text file and key, and encryptes the file data using aes 128 ecb
func EncryptEcb(byteData, byteKey []byte) []byte {

	cipher, err := aes.NewCipher(byteKey)
	if err != nil {
		fmt.Println("aes 128 ecb cipher not found!")
		panic(err)
	}
	// fmt.Printf("total size before padding: %d\n", len(byteData))
	inputBytes, _ := Pkcs7Padding(byteData, len(byteKey))
	// fmt.Printf("total size before encryption: %d\n", len(inputBytes))
	// encrypts data with key, block by block
	encryptedBytes := make([]byte, 0)
	for i := 0; i < len(inputBytes); i += len(byteKey) {
		temp := make([]byte, len(byteKey))
		if i+len(byteKey) < len(inputBytes) {
			cipher.Encrypt(temp, inputBytes[i:i+len(byteKey)])
		} else {
			// fmt.Println(len(inputBytes[i:len(inputBytes)]))
			cipher.Encrypt(temp, inputBytes[i:len(inputBytes)])
		}
		encryptedBytes = append(encryptedBytes, temp...)
	}
	return encryptedBytes

}
