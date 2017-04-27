package set1

import (
	"crypto/aes"
	"fmt"
)

// DecryptEcb takes a byte data, string key and decrypts the file data using aes 128 ecb
func DecryptEcb(inputBytes, byteKey []byte) []byte {
	// byteKey := []byte("YELLOW SUBMARINE")

	cipher, err := aes.NewCipher(byteKey)
	if err != nil {
		fmt.Println("aes 128 ecb cipher not found!")
		panic(err)
	}

	decryptedBytes := make([]byte, 0)
	for i := 0; i < len(inputBytes); i += len(byteKey) {
		temp := make([]byte, len(byteKey))
		cipher.Decrypt(temp, inputBytes[i:i+len(byteKey)])
		decryptedBytes = append(decryptedBytes, temp...)
	}
	return decryptedBytes

}

// EncryptEcb takes a text file and key string, and encryptes the file data using aes 128 ecb
func EncryptEcb(inputBytes, byteKey []byte) []byte {

	cipher, err := aes.NewCipher(byteKey)
	if err != nil {
		fmt.Println("aes 128 ecb cipher not found!")
		panic(err)
	}

	encryptedBytes := make([]byte, 0)
	for i := 0; i < len(inputBytes); i += len(byteKey) {
		temp := make([]byte, len(byteKey))
		cipher.Encrypt(temp, inputBytes[i:i+len(byteKey)])
		encryptedBytes = append(encryptedBytes, temp...)
	}
	return encryptedBytes

}
