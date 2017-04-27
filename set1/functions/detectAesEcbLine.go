package set1

import (
	"bytes"
	"fmt"
)

// DetectAesEcbLine reads a hex AES ECB file, and tries to detect which line is encrypted
func DetectAesEcbLine(byteData []byte) string {

	// file, err := os.Open(fileLoc)
	// if err != nil {
	// 	fmt.Println("8.txt not found!")
	// }
	// input := bufio.NewScanner(file)

	byteLines := bytes.Split(byteData, []byte("\n"))

	finalDu := 0
	finalLine := ""
	for _, bLine := range byteLines {
		// fmt.Println(input.Text())
		line := string(bLine)
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
	// fmt.Println(finalLine)
	return finalLine
}
