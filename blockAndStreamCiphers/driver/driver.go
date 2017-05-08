package main

import (
	"encoding/base64"
	"fmt"

	"github.com/goCryptography/blockAndStreamCiphers"
)

func main() {

	fmt.Println("Challenge 12")
	postAppend := "Um9sbGluJyBpbiBteSA1LjAKV2l0aCBteSByYWctdG9wIGRvd24gc28gbXkgaGFpciBjYW4gYmxvdwpUaGUgZ2lybGllcyBvbiBzdGFuZGJ5IHdhdmluZyBqdXN0IHRvIHNheSBoaQpEaWQgeW91IHN0b3A/IE5vLCBJIGp1c3QgZHJvdmUgYnkK"
	byteData, _ := base64.StdEncoding.DecodeString(postAppend)

	blockAndStreamCiphers.BreakEcb(byteData)

}
