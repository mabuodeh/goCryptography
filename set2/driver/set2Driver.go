package main

import (
	"encoding/base64"
	"fmt"

	"github.com/goCryptography/set1/functions"
	"github.com/goCryptography/set2/functions"
)

func main() {
	fileDirectory := "../src/github.com/goCryptography/set2/files/"

	fmt.Println("Challenge 9")
	paddedData := set2.Pkcs7Padding([]byte("YELLOW SUBMARINE"), 16)
	fmt.Println(string(paddedData))
	fmt.Println()

	fmt.Println("Challenge 10")
	byteData := set1.OpenBase64File(fileDirectory + "10.txt")
	// set2.DecryptCbc(byteData, []byte("YELLOW SUBMARINE"), []byte{48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48})
	fmt.Println(string(set2.DecryptCbc(byteData, []byte("YELLOW SUBMARINE"), []byte{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0})))

	fmt.Println("Challenge 11")
	cha11ByteData := set1.OpenTextFile(fileDirectory + "text.txt")
	fmt.Println(set2.EncryptionOracle(cha11ByteData, true))
	// cha11ByteData = set1.OpenBase64File("../src/github.com/goCryptography/set1/files/8.txt")
	// fmt.Println(set2.EncryptionOracle(cha11ByteData))

	fmt.Println("Challenge 12")
	// plaintext12 := set1.OpenTextFile(fileDirectory + "text.txt")
	plaintext12 := "A"
	base64Str12 := "Um9sbGluJyBpbiBteSA1LjAKV2l0aCBteSByYWctdG9wIGRvd24gc28gbXkgaGFpciBjYW4gYmxvdwpUaGUgZ2lybGllcyBvbiBzdGFuZGJ5IHdhdmluZyBqdXN0IHRvIHNheSBoaQpEaWQgeW91IHN0b3A/IE5vLCBJIGp1c3QgZHJvdmUgYnkK"
	preAppend, _ := base64.StdEncoding.DecodeString(base64Str12)
	fmt.Println(set2.BreakEcb(plaintext12, preAppend))
}
