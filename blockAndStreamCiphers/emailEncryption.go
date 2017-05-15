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
	// 3 blocks, 48 bytes
	first := encryptedData[:16]
	second := encryptedData[16:32]
	third := encryptedData[32:]

	final := make([]byte, 0)
	final = append(final, first...)
	final = append(final, second...)

	admin := first[6:11]

	mod := make([]byte, 0)
	mod = append(mod, third[:6]...)
	mod = append(mod, admin...)
	mod = append(mod, third[11:]...)

	final = append(final, mod...)
	fmt.Println(first)
	fmt.Println(mod)

	// form:
	// email=adminfooba
	// r@bar.com&uid=10
	// &role=user666666, length = 42 + 6
	// ignore first 7 bytes, cut next 5 bytes (ignore email=, cut admin)
	// admin := encryptedData[6:11]
	// // pre := encryptedData[:6]
	// // post := encryptedData[11:]
	// fmt.Println(len(encryptedData))
	// fmt.Println(admin)
	// // encryptedData = append(encryptedData[0:6], encryptedData[11:]...)
	//
	// fmt.Println(encryptedData)
	// // subtract length of admin & user for paste beginning location (30)
	// final := make([]byte, 0)
	// final = append(final, encryptedData[:40]...)
	// // final = append(final, encryptedData[:43]...)
	// final = append(final, admin...)
	// final = append(final, encryptedData[43:]...)
	// fmt.Println(final)

	// decrypt data with key
	plaintext := DecryptEcb(final, key)
	fmt.Println(string(plaintext))
}
