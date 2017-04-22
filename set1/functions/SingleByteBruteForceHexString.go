package set1

import (
	"encoding/hex"
	"fmt"
)

// SingleByteBruteForceHexString takes an encrypted hex string, xors it with different byte combinations
func SingleByteBruteForceHexString(hexStringMsg string) string {
	var ret []byte

	// convert hex string to bytes
	byteMsg, _ := hex.DecodeString(hexStringMsg)
	//oneByte := byte(58)
	// loop through an unsigned byte
	for oneByte := 0; oneByte < 128; oneByte++ {
		var xoredBytes []byte
		for _, bytes := range byteMsg {
			xoredBytes = append(xoredBytes, bytes^byte(oneByte))
		}
		fmt.Println(string(xoredBytes[:]))
		fmt.Println()
	}

	// xor byte with byteMsg
	// endloop
	return string(ret[:])
}
