package blockAndStreamCiphers

import "fmt"

// DetectAesEcbLine checks which line has duplicates. Duplicates = ECB
func DetectAesEcbLine(byteData []byte) string {

	const blockSize = 16

	finalDu := 0
	finalLine := ""
	subStr := make(map[[blockSize]byte]int)

	// iterate over byteData, checking if there are any duplicates based on block size
	for i := range byteData {
		var temp [16]byte
		// temp := make([]byte, 16)

		// i cannot be within the last block of 16, since we take blocks of [i:i+16), so we'll get out of bounds
		if i < (len(byteData) - blockSize) {
			copy(temp[:], byteData[i:i+blockSize])
			// fmt.Println(byteData[i : i+blockSize])
			subStr[temp]++
		} else {
			break
		}
	}
	// if any of the blocks was duplicated, it's ECB
	for sub, count := range subStr {
		// if there are duplicates, then count > 1
		if count > 1 {
			finalDu = count
			finalLine = string(sub[:])
			fmt.Printf("%v repeated %d times\n", sub, count)
		}
	}

	fmt.Println(finalDu)
	fmt.Printf("duplicated block: %v\n", []byte(finalLine))
	// fmt.Println(finalLine)
	return finalLine
}
