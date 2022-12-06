package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	inputFile := "input.txt"
	elfMap := createElfMap(inputFile)
	me, mc := findMostCaloriesElf(elfMap)
	fmt.Printf("The top elf calories is: %d and his calories is: %d \n", me, mc)
	mc1, mc2, mc3 := find3TopElfCalories(elfMap)
	fmt.Printf("The top 3 Elf Calories are: %d, %d, %d \n", mc1, mc2, mc3)
	fmt.Printf("The total Top 3 Elf calories are: %d \n", mc1+mc2+mc3)
}

func createElfMap(input string) map[int]int {
	ce := 0
	elfMap := make(map[int]int)

	// read input file
	f, err := os.Open(input)
	if err != nil {
		log.Fatal(err)
		return nil
	}

	defer f.Close()

	fs := bufio.NewScanner(f)
	for fs.Scan() {
		if fs.Text() == "" {
			ce++
		} else {
			c, err := strconv.Atoi(fs.Text())
			if err != nil {
				log.Fatal(err)
				return nil
			}
			elfMap[ce] += c
		}
	}
	return elfMap
}

func findMostCaloriesElf(elfMap map[int]int) (int, int) {
	var mc int
	var me int
	for e, c := range elfMap {
		if c > mc {
			mc = c
			me = e
		}
	}
	return me, mc
}

func find3TopElfCalories(elfMap map[int]int) (int, int, int) {
	mc1, mc2, mc3 := 0, 0, 0
	for _, c := range elfMap {
		if mc1 <= mc2 && mc1 <= mc3 {
			if c > mc1 {
				mc1 = c
			}
		} else if mc2 <= mc1 && mc2 <= mc3 {
			if c > mc2 {
				mc2 = c
			}
		} else {
			if c > mc3 {
				mc3 = c
			}
		}
	}
	return mc1, mc2, mc3
}
