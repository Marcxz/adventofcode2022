package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	arr := readFile("input.txt")
	// fmt.Println(arr)
	_, n := getNumberVisibleFromOutside(arr)

	fmt.Println(n)

	_, major := getTopViewTrees(arr)
	fmt.Println(major)
}

func readFile(input string) [][]int {
	arr := make([][]int, 0)
	f, err := os.Open(input)
	if err != nil {
		log.Fatal(err)
		return nil
	}
	defer f.Close()

	fs := bufio.NewScanner(f)

	for fs.Scan() {
		l := fs.Text()
		arrl := []int{}
		for i := 0; i < len(l); i++ {
			arrl = append(arrl, int(l[i])-48)
		}
		arr = append(arr, arrl)
	}

	return arr
}

func getRow(arr [][]int, row int) []int {
	return arr[row]
}

func getColumn(arr [][]int, col int) []int {
	colArr := []int{}
	for i := 0; i < len(arr); i++ {
		colArr = append(colArr, arr[i][col])
	}
	return colArr
}

func getNumberVisibleFromOutside(arr [][]int) (map[string]bool, int) {
	rows := len(arr)
	columns := len(arr[0])
	m := make(map[string]bool)
	n := 0
	// Analize rows
	for i := 0; i < rows; i++ {
		row := getRow(arr, i)
		treeHeight := row[0]
		// left to right
		for idx := 0; idx < len(row); idx++ {
			s := strconv.Itoa(i) + "," + strconv.Itoa(idx)
			if idx == 0 {
				m[s] = true
			} else if row[idx] > treeHeight {
				m[s] = true
				treeHeight = row[idx]
			}

		}
		// right to left
		treeHeight = row[len(row)-1]
		for idx := len(row) - 1; idx >= 0; idx-- {
			s := strconv.Itoa(i) + "," + strconv.Itoa(idx)
			if idx == len(row)-1 {
				m[s] = true
			} else if row[idx] > treeHeight {
				m[s] = true
				treeHeight = row[idx]
			}
		}
	}
	// Analize columns
	for i := 0; i < columns; i++ {
		column := getColumn(arr, i)

		// Analize top to bottom
		treeHeight := column[0]
		for idx := 0; idx < len(column); idx++ {
			s := strconv.Itoa(idx) + "," + strconv.Itoa(i)
			if idx == 0 {
				m[s] = true
			} else if column[idx] > treeHeight {
				m[s] = true
				treeHeight = column[idx]
			}
		}
		// Analize bottom to top
		treeHeight = column[len(column)-1]
		for idx := len(column) - 1; idx >= 0; idx-- {
			s := strconv.Itoa(idx) + "," + strconv.Itoa(i)
			if idx == len(column)-1 {
				m[s] = true
			} else if column[idx] > treeHeight {
				m[s] = true
				treeHeight = column[idx]
			}
		}
	}

	for _, i := range m {
		if i {
			n++
		}
	}
	return m, n

}

func getTopViewTrees(arr [][]int) (map[string]int, int) {
	n := 1
	nRow := len(arr)
	nColumn := len(arr[0])
	m := make(map[string]int)
	i, j := 0, 0
	for i < nRow && j < nColumn {
		s := strconv.Itoa(i) + "," + strconv.Itoa(j)
		m[s] = 1
		// analize number to left 1,2
		cont := 0
		for pointer := j - 1; pointer >= 0; pointer-- {
			cont++
			if arr[i][j] <= arr[i][pointer] {
				break
			}
		}
		m[s] *= cont
		// analize number to right
		cont = 0
		for pointer := j + 1; pointer < nColumn; pointer++ {
			cont++
			if arr[i][j] <= arr[i][pointer] {
				break
			}
		}
		m[s] *= cont

		// analize number to top
		cont = 0
		for pointer := i - 1; pointer >= 0; pointer-- {
			cont++
			if arr[i][j] <= arr[pointer][j] {
				break
			}
		}
		m[s] *= cont

		// analize number to down
		cont = 0
		for pointer := i + 1; pointer < nRow; pointer++ {
			cont++
			if arr[i][j] <= arr[pointer][j] {
				break
			}
		}
		m[s] *= cont

		j++
		if j == nColumn && i < nRow {
			j = 0
			i++
		}
	}

	// get the top tree can see
	for _, v := range m {
		if n < v {
			n = v
		}
	}
	return m, n
}
