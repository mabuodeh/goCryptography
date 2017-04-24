package set1

import (
	"bufio"
	"fmt"
	"os"
)

// DetectAesEcbLine reads a hex AES ECB file, and tries to detect which line is encrypted
func DetectAesEcbLine(fileLoc string) {
	// input, err := ioutil.ReadFile(fileLoc)
	// if err != nil {
	// 	fmt.Println("8.txt not found!")
	// 	panic(err)
	// }
	//
	// byteForm, err := hex.DecodeString(string(input))
	// if err != nil {
	// 	fmt.Println("unable to decode hex string into bytes!")
	// 	panic(err)
	// }
	file, err := os.Open(fileLoc)
	if err != nil {
		fmt.Println("8.txt not found!")
	}
	input := bufio.NewScanner(file)

	finalDu := 0
	finalLine := ""
	for input.Scan() {
		// fmt.Println(input.Text())
		line := input.Text()
		subStr := make(map[string]int)
		for i := range line {
			if i <= len(line)-16 {
				subStr[line[i:i+16]]++
			} else {
				break
			}
		}
		for _, count := range subStr {
			if count > finalDu {
				finalDu = count
				finalLine = line
			}
		}
	}
	fmt.Println(finalDu)
	fmt.Println(finalLine)
}
