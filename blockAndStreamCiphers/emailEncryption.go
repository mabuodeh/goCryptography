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
	// form:
	// email=adminfooba r@bar.com&uid=10 &role=user666666, length = 39
	// ignore first 7 bytes, cut next 5 bytes (ignore email=, cut admin)
	admin := encryptedData[6:11]
	// pre := encryptedData[:6]
	// post := encryptedData[11:]
	fmt.Println(len(encryptedData))
	fmt.Println(admin)
	// encryptedData = append(encryptedData[0:6], encryptedData[11:]...)

	fmt.Println(encryptedData)
	// email=foo@bar.com&uid=10&role=user, length = 34
	// subtract length of admin & user for paste beginning location (30)
	final := make([]byte, 0)
	final = append(final, encryptedData[:38]...)
	// final = append(final, encryptedData[:43]...)
	final = append(final, admin...)
	final = append(final, encryptedData[42:]...)
	fmt.Println(final)
	// paste admin at 30
	// email=foo@bar.com&uid=10&role=admin

	// decrypt data with key
	plaintext := DecryptEcb(final, key)
	fmt.Println(string(plaintext))
}
