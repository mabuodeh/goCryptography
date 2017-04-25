package main

import (
	"encoding/hex"
	"fmt"

	"github.com/goCryptography/set1/functions"
)

func main() {
	fileDirectory := "../src/github.com/goCryptography/set1/files/"

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
	cha3Out, ch3Key, _ := set1.SingleByteBruteForceHexString(cha3hexMsg)

	fmt.Println("Challenge 3")
	fmt.Println(cha3Out)
	fmt.Println(hex.EncodeToString([]byte{byte(ch3Key)}))
	fmt.Println()

	// Challenge 4

	cha4Out, _ := set1.SingleByteBruteForceHexStringList(fileDirectory + "4.txt")
	fmt.Println("Challenge 4")
	fmt.Println(cha4Out)
	fmt.Println()

	// Challenge 5

	cha5Key := "ICE"
	cha5In := "Burning 'em, if you ain't quick and nimble\nI go crazy when I hear a cymbal"

	cha5Out := set1.RepeatingKeyXor(cha5In, cha5Key)
	cha5ExOut := "0b3637272a2b2e63622c2e69692a23693a2a3c6324202d623d63343c2a26226324272765272a282b2f20430a652e2c652a3124333a653e2b2027630c692b20283165286326302e27282f"

	fmt.Println("Challenge 5")
	fmt.Println(cha5Out)
	fmt.Println(cha5ExOut)
	fmt.Println(cha5ExOut == cha5Out)
	fmt.Println()

	fmt.Println("Challenge 6")
	fmt.Println(set1.BreakRepeatingKey(fileDirectory + "6.txt"))

	fmt.Println("Challenge 7")
	fmt.Println(set1.EncryptAesEcbWithKey(fileDirectory+"7b.txt", "YELLOW SUBMARINE"))
	fmt.Println(set1.DecryptAesEcbWithKey(fileDirectory+"test.txt", "YELLOW SUBMARINE"))

	fmt.Println("Challenge 8")
	set1.DetectAesEcbLine(fileDirectory + "8.txt")

}
