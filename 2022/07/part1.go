package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Directory struct {
	id            int
	parent        int
	totalFileSize int
	totalSize     int
	children      []*Directory
}

var filesystem = Directory{}

var id = 0

var currentDirId = 0

func main() {
	path := os.Args[1]

	f, err := os.Open(path)
	check(err)

	defer f.Close()

	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		line := scanner.Text()

		if isCommand(line) {
			execCommand(line)
		} else if !isDir(line) {
			dir := findDir(currentDirId, &filesystem)
			dir.totalFileSize += getSize(line)
		}
	}

	sumDir(&filesystem)
	fmt.Println(findAnswer(&filesystem, 0))
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func isDir(input string) bool {
	if input[:3] == "dir" {
		return true
	}

	return false
}

func isCommand(input string) bool {
	if input[:1] == "$" {
		return true
	}

	return false
}

func execCommand(input string) {
	if strings.Contains(input, "cd") {
		if input[5:] == "/" {
			filesystem = Directory{
				id:            0,
				parent:        0,
				totalFileSize: 0,
				children:      []*Directory{},
			}

			currentDirId = 0
		} else if input[5:] == ".." {
			dir := findDir(currentDirId, &filesystem)
			currentDirId = dir.parent
		} else {
			dir := findDir(currentDirId, &filesystem)

			id++

			dir.children = append(dir.children, &Directory{
				id:            id,
				parent:        currentDirId,
				totalFileSize: 0,
				children:      []*Directory{},
			})

			currentDirId = id
		}
	}
}

func getSize(input string) int {
	splitted := strings.Split(input, " ")
	size, err := strconv.Atoi(splitted[0])
	check(err)

	return size
}

func findDir(id int, dir *Directory) *Directory {
	if dir.id == id {
		return dir
	}

	for _, v := range dir.children {
		result := findDir(id, v)

		if result != nil {
			return result
		}
	}

	return nil
}

func sumDir(dir *Directory) {
	totalSize := 0
	totalSize += dir.totalFileSize
	for _, d := range dir.children {
		sumDir(d)
		totalSize += d.totalSize
	}

	dir.totalSize = totalSize
}

func findAnswer(dir *Directory, totalSize int) int {
	if dir.totalSize <= 100000 {
		totalSize += dir.totalSize
	}

	for _, d := range dir.children {
		totalSize = findAnswer(d, totalSize)
	}

	return totalSize
}
