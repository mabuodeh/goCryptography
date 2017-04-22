package set1

import (
	"encoding/base64"
	"encoding/hex"
)

func HexToBase64(hexMsg string) string {
	bitMsg, err := hex.DecodeString(hexMsg)
	if err != nil {
		panic(err)
	}
	base64Msg := base64.StdEncoding.EncodeToString(bitMsg)

	return base64Msg
}
