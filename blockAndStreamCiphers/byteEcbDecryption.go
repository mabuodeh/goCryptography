package blockAndStreamCiphers

import "fmt"

// BreakEcb takes plaintext and a base64 string to be appended before the plaintext, encrypts, and then breaks encryption
func BreakEcb(plaintext, postAppend []byte) string {
	// part one
	// use oracle function, but key must not be random - done: use EncryptionOracle, send (plaintext, false)
	// take a base64 string,
	// decode it,

	// and append it to the plaintext
	// plaintext
	bytesToEncrypt := append(plaintext, postAppend...)
	byteKey, _ := getKeyAndIv(16)
	// encrypt using ECB
	ciphertext := EncryptEcb(bytesToEncrypt, byteKey)

	// check length of ciphertext
	//

	fmt.Print(ciphertext)

	// part two
	return ""

}
