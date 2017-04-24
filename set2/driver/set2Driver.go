package main

import (
	"fmt"

	"github.com/goCryptography/set2/functions"
)

func main() {

	fmt.Println("Challenge 1")
	paddedData := set2.Pkcs7Padding([]byte("YELLOW SUBMARINE"), 16)
	fmt.Println(string(paddedData))

}
