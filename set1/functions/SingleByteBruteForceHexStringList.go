package set1

import (
	"bufio"
	"os"
)

// SingleByteBruteForceHexStringList takes a filename, opens a file of hex strings, returns encrypted string with index of coincidence
func SingleByteBruteForceHexStringList(fileLoc string) (string, float64) {
	finalStr, finalInd := "", 0.0

	file, err := os.Open(fileLoc)
	if err != nil {
		panic(err)
	}

	input := bufio.NewScanner(file)
	for input.Scan() {
		tempStr, _, tempInd := SingleByteBruteForceHexString(input.Text())
		if tempInd > finalInd {
			finalStr = tempStr
			finalInd = tempInd
		}
	}

	return finalStr, finalInd
}
