package set1

import (
	"encoding/hex"
	"unicode"
)

// SingleByteBruteForceHexString takes an encrypted hex string, xors it with different byte combinations
func SingleByteBruteForceHexString(hexStringMsg string) (string, byte, float64) {
	var ret []byte
	finalInd := 0.0
	var finalKey byte

	// convert hex string to bytes
	byteMsg, _ := hex.DecodeString(hexStringMsg)
	//oneByte := byte(58)
	// loop through an unsigned byte
	for oneByte := 0; oneByte < 128; oneByte++ {
		var xoredBytes []byte
		for _, bytes := range byteMsg {
			xoredBytes = append(xoredBytes, bytes^byte(oneByte))
		}
		tempInd := CalcCharFrequency(string(xoredBytes[:]))
		if tempInd > finalInd {
			ret = xoredBytes
			finalInd = tempInd
			finalKey = byte(oneByte)
		}
	}

	return string(ret[:]), finalKey, finalInd
}

// CalcCharFrequency takes a string, and returns the Index of Coincidence of that string
func CalcCharFrequency(s string) float64 {
	indOfCo := 0.0
	letterCount := 0

	charCount := make(map[rune]int)

	for _, char := range s {
		if ('A' <= (char) && (char) <= 'Z') || ('a' <= (char) && (char) <= 'z') || ((char) == ' ') {
			charCount[unicode.ToUpper(char)]++
		}
		letterCount++
	}

	for _, count := range charCount {
		indOfCo += float64(count * (count - 1))
	}
	indOfCo /= float64(letterCount * (letterCount - 1))

	return indOfCo
}
