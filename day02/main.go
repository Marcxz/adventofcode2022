package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

type gameStrategy struct {
	enemyGame string
	youGame   string
}

func main() {
	ags := readStrategy("input.txt")
	es := calculateScoreStrategy(ags)
	fmt.Printf("The total strategy score is: %d \n", es)
	es2 := calculateScoreStrategyPart2(ags)
	fmt.Printf("The total strategy part 2 score is: %d \n", es2)
}

func readStrategy(input string) []gameStrategy {
	ags := make([]gameStrategy, 0)
	// read input file
	f, err := os.Open(input)
	if err != nil {
		log.Fatal(err)
		return nil
	}

	defer f.Close()

	fs := bufio.NewScanner(f)
	for fs.Scan() {
		g := gameStrategy{
			enemyGame: strings.Split(fs.Text(), " ")[0],
			youGame:   strings.Split(fs.Text(), " ")[1],
		}
		ags = append(ags, g)
	}
	return ags
}

func calculateScoreStrategy(ags []gameStrategy) int {
	s := 0
	for _, g := range ags {
		if g.enemyGame == "A" {
			// draw
			if g.youGame == "X" {
				s += 1 + 3
				// win
			} else if g.youGame == "Y" {
				s += 2 + 6
				// lose
			} else if g.youGame == "Z" {
				s += 3 + 0
			}
		} else if g.enemyGame == "B" {
			// lose
			if g.youGame == "X" {
				s += 1 + 0
				// draw
			} else if g.youGame == "Y" {
				s += 2 + 3
				// win
			} else if g.youGame == "Z" {
				s += 3 + 6
			}
		} else if g.enemyGame == "C" {
			// win
			if g.youGame == "X" {
				s += 1 + 6
				// lose
			} else if g.youGame == "Y" {
				s += 2 + 0
				// draw
			} else if g.youGame == "Z" {
				s += 3 + 3
			}
		}
	}
	return s
}

func calculateScoreStrategyPart2(ags []gameStrategy) int {
	s := 0
	for _, g := range ags {
		if g.enemyGame == "A" {
			// lose
			if g.youGame == "X" {
				s += 3 + 0
				// draw
			} else if g.youGame == "Y" {
				s += 1 + 3
				// win
			} else if g.youGame == "Z" {
				s += 2 + 6
			}
		} else if g.enemyGame == "B" {
			// lose
			if g.youGame == "X" {
				s += 1 + 0
				// draw
			} else if g.youGame == "Y" {
				s += 2 + 3
				// win
			} else if g.youGame == "Z" {
				s += 3 + 6
			}
		} else if g.enemyGame == "C" {
			// lose
			if g.youGame == "X" {
				s += 2 + 0
				// draw
			} else if g.youGame == "Y" {
				s += 3 + 3
				// win
			} else if g.youGame == "Z" {
				s += 1 + 6
			}
		}
	}
	return s
}
