package main

import (
	"encoding/base64"
	"fmt"
	"io/ioutil"

	"github.com/spacemonkeygo/openssl"
)

func main() {
	input, err := ioutil.ReadFile("../files/7.txt")
	if err != nil {
		fmt.Println("filenot found!")
		panic(err)
	}
	inputBytes, err := base64.StdEncoding.DecodeString(string(input))
	if err != nil {
		fmt.Println("Could not convert encrypted message to bytes!")
		panic(err)
	}
	key := []byte("YELLOW SUBMARINE")

	cipher, err := openssl.GetCipherByName("aes-128-ecb")
	if err != nil {
		fmt.Println("aes 128 ecb cipher not found!")
		panic(err)
	}

	decryptionTool, err := openssl.NewDecryptionCipherCtx(cipher, nil, key, nil)
	if err != nil {
		fmt.Println("message could not be decrypted!")
		panic(err)
	}

	decryptedMsg, err := decryptionTool.DecryptUpdate(inputBytes)
	if err != nil {
		fmt.Println("message could not be decrypted!")
		panic(err)
	}

	fmt.Println(string(decryptedMsg))

}
