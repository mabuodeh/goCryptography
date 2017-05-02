package set2

import (
	"fmt"

	"github.com/goCryptography/set1/functions"
)

// BreakEcb takes plaintext and a base64 string to be appended before the plaintext, encrypts, and then breaks encryption
func BreakEcb(plaintext, preAppend []byte) string {
	// part one
	// use oracle function, but key must not be random - done: use EncryptionOracle, send (plaintext, false)
	// take a base64 string,
	// decode it,

	// and append it to the plaintext
	// plaintext
	bytesToEncrypt := append(preAppend, plaintext...)
	byteKey, _ := getKeyAndIv(16)
	// encrypt using ECB
	ciphertext := set1.EncryptEcb(bytesToEncrypt, byteKey)

	fmt.Print(ciphertext)

	// part two
	return ""

}
