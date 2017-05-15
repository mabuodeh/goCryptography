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
	// email=foofoofoof
	// adminfoobar@bar.
	// com&uid=10&role=
	// user121212121212121212121212, length = 52 + 12
	// 4 blocks, 64 bytes
	first := encryptedData[:16]
	second := encryptedData[16:32]
	third := encryptedData[32:48]
	// fourth := encryptedData[48:]

	final := make([]byte, 0)
	final = append(final, first...)
	final = append(final, second...)
	final = append(final, third...)
	final = append(final, second...)

	// admin := first[6:11]
	//
	// mod := make([]byte, 0)
	// mod = append(mod, third[:6]...)
	// mod = append(mod, admin...)
	// mod = append(mod, third[11:]...)

	// decrypt data with key
	plaintext := DecryptEcb(final, key)
	fmt.Println(string(plaintext))
}
