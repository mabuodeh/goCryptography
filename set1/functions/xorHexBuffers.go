package set1

import (
	"encoding/hex"
	"fmt"
)

// XorHexStrings takes two hex strings and returns a resultant xor hex string
func XorHexStrings(hexInput1, hexInput2 string) string {
	if len(hexInput1) != len(hexInput2) {
		panic("Size not equal!")
	}

	byteIn1, _ := hex.DecodeString(hexInput1)
	byteIn2, _ := hex.DecodeString(hexInput2)

	var output []byte

	for i, val := range byteIn1 {
		output = append(output, val^byteIn2[i])
	}

	return hex.EncodeToString(output)
}

// XorByteArrays takes two byte arrays and returns a resultant xor byte array
func XorByteArrays(input1, input2 []byte) []byte {
	if len(input1) != len(input2) {
		fmt.Printf("Size not equal. %d != %d\n", len(input1), len(input2))
		panic("At XorByteArrays")
	}

	var output []byte

	for i, val := range input1 {
		output = append(output, val^input2[i])
	}

	return output
}
