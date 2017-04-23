package set1

import "encoding/hex"

// RepeatingKeyXor takes a sentence and and a key, encrypting the sentence with the key. It returns a hex string
func RepeatingKeyXor(s string, key string) string {
	// initialize key counter
	keyCount := 0
	// initialize encrypted string as a byte array
	sBytes := []byte(s)
	keyBytes := []byte(key)
	var xorMsg []byte
	// loop through all the bytes
	for _, bytes := range sBytes {
		// xor the byte with key[counter]
		xorMsg = append(xorMsg, bytes^keyBytes[keyCount])
		// counter = (counter + 1) % len(key)
		keyCount = (keyCount + 1) % len(keyBytes)
	}

	// return HEX encoding of the string
	return hex.EncodeToString(xorMsg)
}
