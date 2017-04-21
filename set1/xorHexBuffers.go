package main

import (
	"encoding/hex"
	"fmt"
)

func XorHexStrings() {
	hexInput1 := "1c0111001f010100061a024b53535009181c"
	hexInput2 := "686974207468652062756c6c277320657965"

	byteIn1, _ := hex.DecodeString(hexInput1)
	byteIn2, _ := hex.DecodeString(hexInput2)

	var output []byte

	for i, val := range byteIn1 {
		output = append(output, val^byteIn2[i])
	}

	expectedOutput := "746865206b696420646f6e277420706c6179"

	fmt.Println(hex.EncodeToString(output))
	fmt.Println(expectedOutput)

	fmt.Println(hex.EncodeToString(output) == expectedOutput)
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
