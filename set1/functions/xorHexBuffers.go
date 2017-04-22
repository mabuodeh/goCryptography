package set1

import "encoding/hex"

// XorHexStrings takes two hex strings and returns a resultant xor hex string
func XorHexStrings(hexInput1 string, hexInput2 string) string {
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
