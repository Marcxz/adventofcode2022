package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

var mConversion = map[string]int{
	"a": 1,
	"b": 2,
	"c": 3,
	"d": 4,
	"e": 5,
	"f": 6,
	"g": 7,
	"h": 8,
	"i": 9,
	"j": 10,
	"k": 11,
	"l": 12,
	"m": 13,
	"n": 14,
	"o": 15,
	"p": 16,
	"q": 17,
	"r": 18,
	"s": 19,
	"t": 20,
	"u": 21,
	"v": 22,
	"w": 23,
	"x": 24,
	"y": 25,
	"z": 26,
	"A": 27,
	"B": 28,
	"C": 29,
	"D": 30,
	"E": 31,
	"F": 32,
	"G": 33,
	"H": 34,
	"I": 35,
	"J": 36,
	"K": 37,
	"L": 38,
	"M": 39,
	"N": 40,
	"O": 41,
	"P": 42,
	"Q": 43,
	"R": 44,
	"S": 45,
	"T": 46,
	"U": 47,
	"V": 48,
	"W": 49,
	"X": 50,
	"Y": 51,
	"Z": 52,
}

type items struct {
	FirstCompartment  bool
	SecondCompartment bool
}

func main() {

	m := readRuckstacks("input.txt")
	sum := getSumRepeated(m)
	fmt.Printf("The sum in part one is %d \n", sum)

	sumPart2 := getSum3ElfGroup(m)
	fmt.Printf("The sum in part two is %d \n", sumPart2)
}

func readRuckstacks(input string) []map[string]items {
	arr := make([]map[string]items, 0)
	f, err := os.Open(input)
	if err != nil {
		log.Fatal(err)
		return nil
	}

	fs := bufio.NewScanner(f)
	for fs.Scan() {
		l := fs.Text()
		fc := l[:len(l)/2]
		sc := l[len(l)/2:]
		mi := make(map[string]items)

		for _, r := range fc {
			s := string(r)
			item := mi[s]
			item.FirstCompartment = true
			mi[s] = item
		}
		for _, r := range sc {
			s := string(r)
			item := mi[s]
			item.SecondCompartment = true
			mi[s] = item
		}
		arr = append(arr, mi)
	}
	return arr
}

func getSumRepeated(arr []map[string]items) int {
	sum := 0
	for _, m := range arr {
		for s, i := range m {
			if i.FirstCompartment && i.SecondCompartment {
				sum += mConversion[s]
			}
		}
	}
	return sum
}

func getSum3ElfGroup(arr []map[string]items) int {
	sum := 0
	for i := 0; i < len(arr); i = i + 3 {
		m1 := arr[i]
		m2 := arr[i+1]
		m3 := arr[i+2]

		for s, _ := range m1 {
			if (m2[s].FirstCompartment || m2[s].SecondCompartment) && (m3[s].FirstCompartment || m3[s].SecondCompartment) {
				sum += mConversion[s]
				break
			}
		}
	}
	return sum
}
