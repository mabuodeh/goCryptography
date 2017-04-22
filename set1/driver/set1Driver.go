package main

import (
	"fmt"

	"github.com/goCryptography/set1/functions"
)

func main() {

	// Challenge 1
	cha1Msg := "49276d206b696c6c696e6720796f757220627261696e206c696b65206120706f69736f6e6f7573206d757368726f6f6d"
	cha1ExO := "SSdtIGtpbGxpbmcgeW91ciBicmFpbiBsaWtlIGEgcG9pc29ub3VzIG11c2hyb29t"

	cha1O := set1.HexToBase64(cha1Msg)

	fmt.Println("Challenge 1")
	fmt.Println(cha1ExO)
	fmt.Println(cha1O)
	fmt.Println(cha1ExO == cha1O)
	fmt.Println()

	// Challenge 2

	cha2HexInput1 := "1c0111001f010100061a024b53535009181c"
	cha2HexInput2 := "686974207468652062756c6c277320657965"

	cha2ExpectedOutput := "746865206b696420646f6e277420706c6179"

	cha2Out := set1.XorHexStrings(cha2HexInput1, cha2HexInput2)

	fmt.Println("Challenge 2")
	fmt.Println(cha2ExpectedOutput)
	fmt.Println(cha2Out)
	fmt.Println(cha2Out == cha2ExpectedOutput)
	fmt.Println()

	// Challenge 3

	cha3hexMsg := "1b37373331363f78151b7f2b783431333d78397828372d363c78373e783a393b3736"
	cha3Out := set1.SingleByteBruteForceHexString(cha3hexMsg)

	fmt.Println(cha3Out)

}
