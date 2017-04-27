package set2

import "github.com/goCryptography/set1/functions"

// EncryptionOracle encrypts data, and states which block type it encrypted with
func EncryptionOracle(byteData []byte) string {

	byteKey := []byte("YELLOW SUBMARINE")

	// encrypt by cbc
	encryptedData := EncryptCbc(byteData, byteKey, []byte{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0})
	// encrypt by ecb
	// encryptedData := set1.EncryptEcb(byteData, byteKey)

	// fmt.Println(encryptedData)
	// encryptedData = set1.DecryptEcb(encryptedData, byteKey)
	// fmt.Println(string(encryptedData))
	// send encrypted data, check whether it's ecb or cbc
	duplicates := set1.DetectAesEcbLine(encryptedData)

	if duplicates == "" {
		return "CBC"
	}
	return "ECB"

}
