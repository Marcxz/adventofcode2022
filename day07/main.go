package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type fileTree struct {
	path string
	name string
	size int
}

type directoryTree struct {
	name       string
	path       string
	totalSize  int
	totalFiles int
	found      bool
	files      []fileTree
}

func main() {

	commands := readFile("input.txt")
	mDirectory, arrF := createFileTree(commands)
	mDirectory, arrF = calculateDirectoryMetadata(mDirectory, arrF)
	arrDFiltered, size := getAlmostDirectories(mDirectory, 100000)
	for _, d := range arrDFiltered {
		fmt.Printf("%s %d \n", d.path, d.totalSize)
	}
	fmt.Println(size)

}

func getAlmostDirectories(m map[string]directoryTree, size int) ([]directoryTree, int) {
	s := 0
	mF := make([]directoryTree, 0)
	for _, d := range m {
		if d.totalSize <= size {
			mF = append(mF, d)
			s += d.totalSize
		}

	}
	return mF, s
}
func createFileTree(commands []string) (map[string]directoryTree, []fileTree) {
	m := make(map[string]directoryTree, 0)
	arrF := make([]fileTree, 0)
	path := ""
	td := ""
	for _, s := range commands {
		// inspect to evaluate the commands
		arr := strings.Split(s, " ")
		if arr[0] == "$" {
			// cd command
			if arr[1] == "cd" {
				// root
				if arr[2] == "/" {
					path = "/"
					// directory
				} else if arr[2] != ".." {
					if path == "/" {
						path += arr[2]
					} else {
						path += "/" + arr[2]
					}
					// upper directory
				} else {
					arrPath := strings.Split(path, "/")
					arrPath = arrPath[:len(arrPath)-1]
					path = strings.Join(arrPath, "/")
				}
				td = ""
				// is a command directory
			} else if arr[0] == "ls" {
				td = ""
			}
		} else if arr[0] == "dir" {
			// td = arr[1]
			// is a file
		} else if size, err := strconv.Atoi(arr[0]); err == nil {

			file := fileTree{
				name: arr[1],
				size: size,
				path: "",
			}

			if td == "" {
				file.path = path
			} else {
				// is in root
				if path == "/" {
					file.path = path + td
				} else {
					file.path = path + "/" + td
				}
			}
			arrF = append(arrF, file)
			// Check the directories lets see
			if !m[file.path].found {
				m[file.path] = directoryTree{
					name:       file.path,
					path:       file.path,
					totalSize:  0,
					totalFiles: 0,
					files:      make([]fileTree, 0),
					found:      true,
				}
			}
		}
	}
	return m, arrF
}

func calculateDirectoryMetadata(m map[string]directoryTree, arrF []fileTree) (map[string]directoryTree, []fileTree) {
	for _, f := range arrF {
		// evaluate each directory
		for _, d := range m {

			if strings.Contains(f.path, d.path) {
				d.files = append(d.files, f)
				d.totalFiles++
				d.totalSize += f.size
				m[d.path] = d
			}
		}
	}
	return m, arrF
}

func readFile(input string) []string {
	arr := make([]string, 0)
	// read the file
	f, err := os.Open(input)
	if err != nil {
		log.Fatal(err)
		return nil
	}
	defer f.Close()

	// create an file scanner
	fs := bufio.NewScanner(f)
	for fs.Scan() {
		arr = append(arr, fs.Text())
	}

	return arr
}
