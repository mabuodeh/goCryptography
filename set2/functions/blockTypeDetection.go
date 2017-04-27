package set2

import (
	"github.com/goCryptography/set1/functions"
)

// EncryptionOracle checks which block type was used
func EncryptionOracle(byteData []byte) string {

	// send encrypted data, check whether it's ecb or cbc
	duplicates := set1.DetectAesEcbLine(byteData)

	if duplicates == "" {
		return "CBC"
	} else {
		return "ECB"
	}

}
