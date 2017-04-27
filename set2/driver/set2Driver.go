package main

import (
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
	// should give cbc
	cha11ByteData := set1.OpenBase64File(fileDirectory + "10.txt")
	fmt.Println(set2.EncryptionOracle(cha11ByteData))

}
