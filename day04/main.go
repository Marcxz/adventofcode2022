package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	arr := readSectionInputs("input.txt")
	nPair := findNumberPairCompleteFilled(arr)
	fmt.Printf("The number of times that a elve is completed filled is %d \n", nPair)

	nOverlaps := findNumberPairOverlap(arr)
	fmt.Printf("The number of times that a section overlaps is %d \n", nOverlaps)
}

type pair struct {
	slotsA   int
	slotsB   int
	sections map[int]int
}

func readSectionInputs(input string) []pair {
	arr := make([]pair, 0)

	// read the file
	f, err := os.Open(input)
	if err != nil {
		log.Fatal(err)
		return nil
	}

	// Scan each line
	fs := bufio.NewScanner(f)
	for fs.Scan() {

		l := fs.Text()
		fp := strings.Split(l, ",")[0]
		sp := strings.Split(l, ",")[1]

		ifp, err := strconv.Atoi(strings.Split(fp, "-")[0])
		if err != nil {
			log.Fatal(err)
			return nil
		}

		ffp, err := strconv.Atoi(strings.Split(fp, "-")[1])
		if err != nil {
			log.Fatal(err)
			return nil
		}

		isp, err := strconv.Atoi(strings.Split(sp, "-")[0])
		if err != nil {
			log.Fatal(err)
			return nil
		}

		fsp, err := strconv.Atoi(strings.Split(sp, "-")[1])
		if err != nil {
			log.Fatal(err)
			return nil
		}

		camp := pair{
			slotsA:   0,
			slotsB:   0,
			sections: make(map[int]int, 0),
		}
		// fill map first pair
		for i := ifp; i <= ffp; i++ {
			camp.sections[i]++
			camp.slotsA++
		}

		// fill map first pair
		for i := isp; i <= fsp; i++ {
			camp.sections[i]++
			camp.slotsB++
		}
		arr = append(arr, camp)
	}
	return arr
}

func findNumberPairCompleteFilled(arr []pair) int {
	numberPair := 0
	for _, camp := range arr {
		nSection := 0
		for _, nTimes := range camp.sections {
			if nTimes == 2 {
				nSection++
			}
		}
		if nSection >= camp.slotsA || nSection >= camp.slotsB {
			numberPair++
		}
	}
	return numberPair
}

func findNumberPairOverlap(arr []pair) int {
	nOverlaps := 0
	for _, camp := range arr {
		for _, nTimes := range camp.sections {
			if nTimes == 2 {
				nOverlaps++
				break
			}
		}
	}
	return nOverlaps
}
