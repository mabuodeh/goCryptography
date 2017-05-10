package blockAndStreamCiphers

import (
	"fmt"

	"github.com/goCryptography/utilities"
)

// ModifyCredentials takes an email, creates a stream token, and encrypts email using a key. The message is modified, then decrypted
func ModifyCredentials(email string) {
	blockSize := 16
	// encode email in token stream
	msg := utilities.ProfileFor(email)
	fmt.Println(msg)
	// get key
	key, _ := getKeyAndIv(blockSize)
	// encrypt data with key
	encryptedData := EncryptEcb([]byte(msg), key)

	// hacking here

	// decrypt data with key
	plaintext := DecryptEcb(encryptedData, key)
	fmt.Println(string(plaintext))
}
