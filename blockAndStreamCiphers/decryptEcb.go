package blockAndStreamCiphers

import (
	"crypto/aes"
	"fmt"
)

// DecryptEcb takes a byte data, string key and decrypts the file data using aes 128 ecb
func DecryptEcb(inputBytes, byteKey []byte) []byte {
	// creating a new aes ecb 128 block cipher
	cipher, err := aes.NewCipher(byteKey)
	if err != nil {
		fmt.Println("aes 128 ecb cipher not found!")
		panic(err)
	}

	// manually decrypting, block by block
	decryptedBytes := make([]byte, 0)
	for i := 0; i < len(inputBytes); i += len(byteKey) {
		temp := make([]byte, len(byteKey))
		cipher.Decrypt(temp, inputBytes[i:i+len(byteKey)])
		decryptedBytes = append(decryptedBytes, temp...)
	}
	return decryptedBytes

}
