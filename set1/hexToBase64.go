package main

import (
	"encoding/base64"
	"encoding/hex"
	"fmt"
)

func main() {
	msg := "49276d206b696c6c696e6720796f757220627261696e206c696b65206120706f69736f6e6f7573206d757368726f6f6d"
	bitMsg, err := hex.DecodeString(msg)
	if err != nil {
		panic(err)
	}
	output1 := base64.StdEncoding.EncodeToString(bitMsg)
	output2 := "SSdtIGtpbGxpbmcgeW91ciBicmFpbiBsaWtlIGEgcG9pc29ub3VzIG11c2hyb29t"

	fmt.Println(output1)
	fmt.Println(output2)
	fmt.Println(output1 == output2)

}
