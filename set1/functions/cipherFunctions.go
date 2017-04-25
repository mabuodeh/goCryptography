package set1

import (
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"io/ioutil"

	"github.com/spacemonkeygo/openssl"
)

// DecryptWithCipherKeyIv takes a base64 encrypted file, cipher name, string key, and hex string iv, and decrypts the file data
func DecryptWithCipherKeyIv(fileLoc, cipherName, key, ivStr string) string {
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

	var iv []byte
	if ivStr == "" {
		iv = nil
	} else {
		iv, err = hex.DecodeString(ivStr)
		if err != nil {
			fmt.Println("initialization vector issue at decrypt aes!")
		}
	}

	cipher, err := openssl.GetCipherByName(cipherName)
	if err != nil {
		fmt.Println("aes 128 ecb cipher not found!")
		panic(err)
	}

	decryptionTool, err := openssl.NewDecryptionCipherCtx(cipher, nil, byteKey, iv)
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

// EncryptWithCipherKeyIv takes a text file, cipher name, key string, and hex string iv, and encryptes the file data
func EncryptWithCipherKeyIv(fileLoc, cipherName, key, ivStr string) string {
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

	var iv []byte
	if ivStr == "" {
		iv = nil
	} else {
		iv, err = hex.DecodeString(ivStr)
		if err != nil {
			fmt.Println("initialization vector issue at encrypt aes!")
		}
	}

	cipher, err := openssl.GetCipherByName(cipherName)
	if err != nil {
		fmt.Println("aes 128 ecb cipher not found!")
		panic(err)
	}

	// decryptionTool, err := openssl.NewDecryptionCipherCtx(cipher, nil, byteKey, nil)
	encryptionTool, err := openssl.NewEncryptionCipherCtx(cipher, nil, byteKey, iv)
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
