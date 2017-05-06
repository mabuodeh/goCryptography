package blockAndStreamCiphers

import (
	"crypto/aes"
	"fmt"
)

// EncryptEcb takes a text file and key, and encryptes the file data using aes 128 ecb
func EncryptEcb(inputBytes, byteKey []byte) []byte {

	cipher, err := aes.NewCipher(byteKey)
	if err != nil {
		fmt.Println("aes 128 ecb cipher not found!")
		panic(err)
	}

	// encrypts data with key, block by block
	encryptedBytes := make([]byte, 0)
	for i := 0; i < len(inputBytes); i += len(byteKey) {
		temp := make([]byte, len(byteKey))
		cipher.Encrypt(temp, inputBytes[i:i+len(byteKey)])
		encryptedBytes = append(encryptedBytes, temp...)
	}
	return encryptedBytes

}
