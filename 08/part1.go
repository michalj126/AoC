package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	path := os.Args[1]

	f, err := os.Open(path)
	check(err)

	defer f.Close()

	m := parse(f)

	fmt.Println(countVisibleTrees(&m))
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func parse(f *os.File) (m [][]int) {
	scanner := bufio.NewScanner(f)

	m = [][]int{{}}

	for scanner.Scan() {
		line := scanner.Text()

		for i, v := range line {
			vInt, err := strconv.Atoi(string(v))
			check(err)

			if len(m) == i {
				m = append(m, []int{})
			}

			m[i] = append(m[i], vInt)
		}
	}

	return
}

func countVisibleTrees(m *[][]int) (counter int) {
	counter = 0

	counter += len((*m))*2 + len((*m)[0])*2 - 4

	for y, row := range (*m)[1 : len((*m))-1] {
		for x := range (*m)[y+1][1 : len(row)-1] {
			result := isVisibleInside((*m)[x+1][y+1], m, x+1, y+1)

			if result {
				counter++
			}
		}
	}

	return
}

func isTopVisible(tree int, m *[][]int, x int, y int) bool {
	for i := 0; i < y; i++ {
		if tree <= (*m)[x][i] {
			return false
		}
	}

	return true
}

func isBottomVisible(tree int, m *[][]int, x int, y int) bool {
	for i := y + 1; i < len((*m)[0]); i++ {
		if tree <= (*m)[x][i] {
			return false
		}
	}

	return true
}

func isLeftVisible(tree int, m *[][]int, x int, y int) bool {
	for i := 0; i < x; i++ {
		if tree <= (*m)[i][y] {
			return false
		}
	}

	return true
}

func isRightVisible(tree int, m *[][]int, x int, y int) bool {
	for i := x + 1; i < len((*m)); i++ {
		if tree <= (*m)[i][y] {
			return false
		}
	}

	return true
}

func isVisibleInside(tree int, m *[][]int, x int, y int) bool {
	if isTopVisible(tree, m, x, y) {
		return true
	}

	if isBottomVisible(tree, m, x, y) {
		return true
	}

	if isLeftVisible(tree, m, x, y) {
		return true
	}

	if isRightVisible(tree, m, x, y) {
		return true
	}

	return false
}
