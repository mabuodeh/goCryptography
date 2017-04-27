package set1

import (
	"encoding/base64"
	"fmt"
	"io/ioutil"
)

// OpenTextFile opens a text file and returns the data in bytes
func OpenTextFile(fileLoc string) []byte {
	input, err := ioutil.ReadFile(fileLoc)
	if err != nil {
		fmt.Println("filenot found!")
		panic(err)
	}

	return input

}

// WriteToFile writes base64 to files
func WriteToFile(byteData []byte, fileLoc string) {
	err := ioutil.WriteFile(fileLoc, byteData, 0644)
	fmt.Println(err)
}

// OpenBase64File opens a base64 file and returns data bytes
func OpenBase64File(fileLoc string) []byte {
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
	return inputBytes
}

// OpenHexFile opens a hex file and returns data bytes
func OpenHexFile(fileLoc string) []byte {
	input, err := ioutil.ReadFile(fileLoc)
	if err != nil {
		fmt.Println("filenot found!")
		panic(err)
	}
	// fmt.Println(input)
	// inputBytes, err := hex.DecodeString(string(input))
	// if err != nil {
	// 	fmt.Println("Could not convert encrypted hex message to bytes!")
	// 	panic(err)
	// }
	return input
}
