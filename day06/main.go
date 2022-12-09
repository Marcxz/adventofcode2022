package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	strCode := readStringSignal("input.txt")
    p4 := findFirstBlockAllCharactersDifferent(strCode, 4)
	p14 := findFirstBlockAllCharactersDifferent(strCode, 14)
	fmt.Println(p4)
	fmt.Println(p14)

}

func findFirstBlockAllCharactersDifferent(str string, block int) int {
	p := -1
	for i := 0; i < len(str); i++ {
		s := str[i : i+block]
		if findIfAllCharactersAreDifferent(s) {
			p = i + block
			break
		}
	}
	return p
}
func findIfAllCharactersAreDifferent(str string) bool {
	isDifferent := true
	m := make(map[string]bool)
	for _, c := range str {
		if m[string(c)] {
			isDifferent = false
			break
		}
		m[string(c)] = true
	}
	return isDifferent
}

func readStringSignal(input string) string {

	// read the file
	f, err := os.Open(input)
	if err != nil {
		log.Fatal(err)
		return ""
	}
	defer f.Close()

	// create the file scanner
	fs := bufio.NewScanner(f)
	fs.Scan()
	return fs.Text()
}
