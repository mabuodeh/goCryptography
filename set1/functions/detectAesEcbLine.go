package set1

import "fmt"

// DetectAesEcbLine reads a hex AES ECB file, and tries to detect which line is encrypted
func DetectAesEcbLine(byteData []byte) string {

	// file, err := os.Open(fileLoc)
	// if err != nil {
	// 	fmt.Println("8.txt not found!")
	// }
	// input := bufio.NewScanner(file)

	// byteLines := bytes.Split(byteData, []byte("\n"))
	// fmt.Println(byteLines)
	const blockSize = 16

	finalDu := 0
	finalLine := ""
	subStr := make(map[[blockSize]byte]int)
	for i := range byteData {
		var temp [16]byte
		// temp := make([]byte, 16)

		if i < (len(byteData) - blockSize) {
			copy(temp[:], byteData[i:i+blockSize])
			// fmt.Println(byteData[i : i+blockSize])
			subStr[temp]++
		} else {
			break
		}
	}
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
