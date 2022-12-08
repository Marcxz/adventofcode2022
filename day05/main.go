package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

var stack = map[int][]string{
	1: {"D", "L", "J", "R", "V", "G", "F"},
	2: {"T", "P", "M", "B", "V", "H", "J", "S"},
	3: {"V", "H", "M", "F", "D", "G", "P", "C"},
	4: {"M", "D", "P", "N", "G", "Q"},
	5: {"J", "L", "H", "N", "F"},
	6: {"N", "F", "V", "Q", "D", "G", "T", "Z"},
	7: {"F", "D", "B", "L"},
	8: {"M", "J", "B", "S", "V", "D", "N"},
	9: {"G", "L", "D"},
}

type moves struct {
	times int
	from  int
	to    int
}

func main() {
	moveSet := createMoves("input.txt")

	// Part 1
	/*
		executeMoves(stack, moveSet)
		topCreates := getLastTopCreates(stack)
		fmt.Println("The top creates of each stack is ", topCreates)

	*/

	// Part 2
	executeMovesCreateMover9001(stack, moveSet)
	topCreates := getLastTopCreates(stack)
	fmt.Println("The top creates of each stack using Create Mover 9001 is ", topCreates)
}

func getLastTopCreates(stack map[int][]string) string {
	creates := ""
	for i := 1; i <= 9; i++ {
		creates += stack[i][len(stack[i])-1]
	}
	return creates
}

func executeMoves(stack map[int][]string, moveSet []moves) {
	for _, m := range moveSet {
		move(stack, m.times, m.from, m.to)
	}
}

func move(stack map[int][]string, times int, from int, to int) {
	for i := 0; i < times; i++ {
		letterToMove := stack[from][len(stack[from])-1]
		// remove item from stack
		stack[from] = stack[from][:len(stack[from])-1]
		// insert item to stack
		stack[to] = append(stack[to], letterToMove)
	}
}

func executeMovesCreateMover9001(stack map[int][]string, moveSet []moves) {
	for _, m := range moveSet {
		moveCreateMover9001(stack, m.times, m.from, m.to)
	}
}

func moveCreateMover9001(stack map[int][]string, times int, from int, to int) {
	arrToMove := stack[from][len(stack[from])-times:]
	// remove item from stack
	stack[from] = stack[from][:len(stack[from])-times]
	// insert item to stack
	stack[to] = append(stack[to], arrToMove...)
}

func createMoves(input string) []moves {
	arr := make([]moves, 0)

	// read the file
	f, err := os.Open(input)
	if err != nil {
		log.Fatal(err)
		return nil
	}

	// read over lines
	fs := bufio.NewScanner(f)

	for fs.Scan() {

		t, err := strconv.Atoi(strings.Split(fs.Text(), " ")[1])
		if err != nil {
			log.Fatal(err)
			return nil
		}

		from, err := strconv.Atoi(strings.Split(fs.Text(), " ")[3])
		if err != nil {
			log.Fatal(err)
			return nil
		}

		to, err := strconv.Atoi(strings.Split(fs.Text(), " ")[5])
		if err != nil {
			log.Fatal(err)
			return nil
		}

		m := moves{
			times: t,
			from:  from,
			to:    to,
		}
		arr = append(arr, m)
	}
	return arr
}
