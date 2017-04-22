package main

import (
	"fmt"
	"goCryptography/set1/functions"
)

func main() {

	// Challenge 1
	cha1Msg := "49276d206b696c6c696e6720796f757220627261696e206c696b65206120706f69736f6e6f7573206d757368726f6f6d"
	cha1ExO := "SSdtIGtpbGxpbmcgeW91ciBicmFpbiBsaWtlIGEgcG9pc29ub3VzIG11c2hyb29t"

	cha1O := set1.HexToBase64(cha1Msg)

	fmt.Println(cha1ExO)
	fmt.Println(cha1O)
	fmt.Println(cha1ExO == cha1O)

	// hexInput1 := "1c0111001f010100061a024b53535009181c"
	// hexInput2 := "686974207468652062756c6c277320657965"
	//
	// expectedOutput := "746865206b696420646f6e277420706c6179"
	//
	// fmt.Println(hex.EncodeToString(output))
	// fmt.Println(expectedOutput)
	//
	// fmt.Println(hex.EncodeToString(output) == expectedOutput)

	// XorHexStrings()

}
