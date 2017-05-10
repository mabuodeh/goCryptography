package main

import (
	"fmt"

	"github.com/goCryptography/utilities"
)

func main() {

	// fmt.Println("Challenge 12")
	// postAppend := "Um9sbGluJyBpbiBteSA1LjAKV2l0aCBteSByYWctdG9wIGRvd24gc28gbXkgaGFpciBjYW4gYmxvdwpUaGUgZ2lybGllcyBvbiBzdGFuZGJ5IHdhdmluZyBqdXN0IHRvIHNheSBoaQpEaWQgeW91IHN0b3A/IE5vLCBJIGp1c3QgZHJvdmUgYnkK"
	// byteData, _ := base64.StdEncoding.DecodeString(postAppend)
	//
	// blockAndStreamCiphers.BreakEcb(byteData)

	fmt.Println("Challenge 13")
	// str := "foo=bar&baz=qux&zap=zazzle"
	// fmt.Printf("before parsing: %v\n", str)
	// parsed := utilities.Parse(str)
	// fmt.Printf("after parsing: %v\n", parsed)
	// fmt.Printf("after unparsing: %v\n", utilities.Unparse(parsed))
	fmt.Println(utilities.ProfileFor("foo@email.com"))
	// utilities.ProfileFor("foo@email.com")

}
