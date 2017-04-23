package set1

import (
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"io/ioutil"
)

// BreakRepeatingKey takes the location of a base64 encrypted file. It will determine key length, then decrypt it
func BreakRepeatingKey(fileLoc string) string {
	input, err := ioutil.ReadFile(fileLoc)
	if err != nil {
		panic(err)
	}
	//byteForm := make([]byte, base64.StdEncoding.DecodedLen(len(input)))
	//_, err = base64.StdEncoding.Decode(byteForm, input)
	byteForm, err := base64.StdEncoding.DecodeString(string(input))
	if err != nil {
		panic(err)
	}

	keysize := GetKeysize(byteForm)
	fmt.Printf("Keysize: %d\n", keysize)
	//keysize := 29

	blocks := TransposeBlocks(byteForm, keysize)
	finalKeyByte := make([]byte, keysize)
	for i := 0; i < keysize; i++ {
		_, tempKey, _ := SingleByteBruteForceHexString(hex.EncodeToString(blocks[i]))
		finalKeyByte = append(finalKeyByte, byte(tempKey))
	}
	finalKey := string(finalKeyByte)
	// finalKey = strings.Join(finalKey[:], "\n")
	fmt.Println(finalKey)

	decryptedMsg := RepeatingKeyXor(string(byteForm[:]), finalKey)
	decryptedMsgBytes, _ := hex.DecodeString(decryptedMsg)
	return string(decryptedMsgBytes)
}

// GetKeysize takes a byte array, and calculates its possible keysize by transposition and index of coincidence techniques
func GetKeysize(b []byte) int {
	// initialize final keysize and final index of coincidence values
	finalKeysize := 0
	finalIndOfCo := 0.0

	// loop from 2 to 40 (these are the keysizes to test)
	for keysize := 2; keysize <= 40; keysize++ {
		// initialize a block index of coincidence variable
		blockIndOfCo := 0.0
		// transpose the byte array for the given keysize
		transposed := TransposeBlocks(b, keysize)
		// loop through the blocks
		for _, block := range transposed {
			// obtain the index of coincidence of each block
			// add the index of coincidence to the total for this keysize
			_, _, temp := SingleByteBruteForceHexString(hex.EncodeToString(block))
			blockIndOfCo += temp
		} // endloop
		// divide the block index of coincidence with the number of blocks of this keysize
		blockIndOfCo /= float64(len(transposed))
		//fmt.Printf("keysize and ind of co: %d,  %d, %f\n", len(transposed), keysize, blockIndOfCo)
		// if the block index of coincidence is higher than the final index of coincidence
		if blockIndOfCo > finalIndOfCo {
			// replace the final with the block index of coincidence
			finalIndOfCo = blockIndOfCo
			// replace the final keysize with the current loop keysize
			finalKeysize = keysize
		} // endif

	} //end loop

	return finalKeysize
}

// TransposeBlocks takes a byte array, and transposes it into blocks of keysize bytes.
func TransposeBlocks(b []byte, keysize int) [][]byte {
	ret := make([][]byte, keysize)

	// initialize block counter
	blockCounter := 0
	// initialize bytes per block counter
	// bytesPerBlockCounter := 0
	// loop through the bytes
	for _, bytes := range b {
		// append byte to ret[blockCounter]
		ret[blockCounter] = append(ret[blockCounter], bytes)
		// increment bytesPerBlockCounter
		// bytesPerBlockCounter++

		// if bytes per block counter >= keysize
		// if bytesPerBlockCounter >= 1 {
		// reinitialize bytes per block counter
		// bytesPerBlockCounter = 0
		// increment block counter
		blockCounter = (blockCounter + 1) % keysize
		// } // endif
	} // endloop

	return ret
}
