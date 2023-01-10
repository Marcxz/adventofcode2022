package main

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Command struct {
	direction string
	steps     int
}

type Position struct {
	x int
	y int
}

func main() {
	commands := readFile("input.txt")
	m := evaluateCommands(commands)
	fmt.Println(len(m))
}

func evaluateCommands(arrC []Command) map[string]int {
	m := make(map[string]int)
	h := &Position{
		x: 0,
		y: 0,
	}
	t := &Position{
		x: 0,
		y: 0,
	}
	
	// position 0
	st := strconv.Itoa(t.x) + "," + strconv.Itoa(t.y)
	m[st]++

	for _, c := range arrC {
		// do the steps by command
		for i := 0; i < c.steps; i++ {
			// mark the position
			st := strconv.Itoa(t.x) + "," + strconv.Itoa(t.y)
			m[st]++
			evaluatePosition(c.direction, h, t)
		}
	}
	return m
}

func evaluatePosition(command string, h *Position, t *Position) {
	switch command {
	case "U":
		// h and t in the same position
		// tail in the 3 upper position
		if t.y > h.y {
			// do nothing
			// the 3 center positions
		} else if h.y == t.y {
			// do nothing
		} else if h.y > t.y {
			if h.x > t.x {
				t.x++
				t.y++
			} else if h.x == t.x {
				t.y++
			} else {
				t.x--
				t.y++
			}
		}
		h.y++
		break
	case "R":
		if h.x > t.x {
			// upper left
			if t.y > h.y {
				t.x++
				t.y--
				// center left
			} else if t.y == h.y {
				t.x++
				// down left
			} else if t.y < h.y {
				t.y++
				t.x++
			}
		}
		h.x++
		break
	case "L":
		if h.x < t.x {
			// upper right t
			if t.y > h.x {
				t.x--
				t.y--
				// center right t
			} else if t.y == h.y {
				t.x--
			} else if t.y < h.y {
				t.x--
				t.y++
			}
		}
		h.x--
		break
	case "D":
		if t.y > h.y {
			// upper left t
			if t.x < h.x {
				t.y--
				t.x++
			} else if t.x == h.x {
				t.y--
			} else if t.x > h.x {
				t.y--
				t.x--
			}
			h.y--
		}
		break
	}
}

func readFile(input string) []Command {
	arr := make([]Command, 0)

	// read the file
	f, err := os.Open(input)
	if err != nil {
		log.Fatal(err)
		return nil
	}
	defer f.Close()

	// read the file
	fs := bufio.NewScanner(f)
	for fs.Scan() {
		l := fs.Text()
		arrL := strings.Split(l, " ")

		d := arrL[0]
		st, err := strconv.Atoi(arrL[1])
		if err != nil {
			log.Fatal(err)
			return nil
		}
		c := Command{
			direction: d,
			steps:     st,
		}
		arr = append(arr, c)
	}
	return arr
}
