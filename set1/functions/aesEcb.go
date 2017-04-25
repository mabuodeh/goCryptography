package set1

import (
	"encoding/base64"
	"fmt"
	"io/ioutil"

	"github.com/spacemonkeygo/openssl"
)

// DecryptAesEcbWithKey takes a base64 AES ECB encrypted file and key, and decrypts it
func DecryptAesEcbWithKey(fileLoc string, key string) string {
	input, err := ioutil.ReadFile(fileLoc)
	if err != nil {
		fmt.Println("filenot found!")
		panic(err)
	}
	inputBytes, err := base64.StdEncoding.DecodeString(string(input))
	if err != nil {
		fmt.Println("Could not convert encrypted message to bytes!")
		panic(err)
	}
	byteKey := []byte("YELLOW SUBMARINE")

	cipher, err := openssl.GetCipherByName("aes-128-ecb")
	if err != nil {
		fmt.Println("aes 128 ecb cipher not found!")
		panic(err)
	}

	decryptionTool, err := openssl.NewDecryptionCipherCtx(cipher, nil, byteKey, nil)
	if err != nil {
		fmt.Println("unable to create decryption tool!")
		panic(err)
	}

	decryptedMsg, err := decryptionTool.DecryptUpdate(inputBytes)
	if err != nil {
		fmt.Println("message could not be decrypted!")
		panic(err)
	}

	return string(decryptedMsg)

}

// EncryptAesEcbWithKey takes a text file and encryptes it using AES ECB 128
func EncryptAesEcbWithKey(fileLoc string, key string) string {
	input, err := ioutil.ReadFile(fileLoc)
	if err != nil {
		fmt.Println("filenot found!")
		panic(err)
	}
	inputBytes := []byte(input)
	// if err != nil {
	// 	fmt.Println("Could not convert encrypted message to bytes!")
	// 	panic(err)
	// }
	byteKey := []byte("YELLOW SUBMARINE")

	cipher, err := openssl.GetCipherByName("aes-128-ecb")
	if err != nil {
		fmt.Println("aes 128 ecb cipher not found!")
		panic(err)
	}

	// decryptionTool, err := openssl.NewDecryptionCipherCtx(cipher, nil, byteKey, nil)
	encryptionTool, err := openssl.NewEncryptionCipherCtx(cipher, nil, byteKey, nil)
	if err != nil {
		fmt.Println("unable to create encryption tool!")
		panic(err)
	}

	encryptedMsg, err := encryptionTool.EncryptUpdate(inputBytes)
	if err != nil {
		fmt.Println("message could not be encrypted!")
		panic(err)
	}

	base64EncryptedMsg := base64.StdEncoding.EncodeToString(encryptedMsg)

	return string(base64EncryptedMsg)

}
