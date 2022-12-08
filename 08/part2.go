package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type ScenicScore struct {
	top    int
	bottom int
	left   int
	right  int
}

func (o ScenicScore) getScore() int {
	return o.top * o.bottom * o.left * o.right
}

func main() {
	path := os.Args[1]

	f, err := os.Open(path)
	check(err)

	defer f.Close()

	m := parse(f)

	fmt.Println(findHighestScenicScore(&m))
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

func calcScenicScore(tree int, m *[][]int, x int, y int) int {
	scenicScore := ScenicScore{}

	scenicScore.top = calcTopViewDistance(tree, m, x, y)
	scenicScore.bottom = calcBottomViewDistance(tree, m, x, y)
	scenicScore.left = calcLeftViewDistance(tree, m, x, y)
	scenicScore.right = calcRightViewDistance(tree, m, x, y)

	return scenicScore.getScore()
}

func calcTopViewDistance(tree int, m *[][]int, x int, y int) (top int) {
	top = 0

	for i := y - 1; i >= 0; i-- {
		if tree > (*m)[x][i] {
			top++
		} else {
			top++
			return
		}
	}

	return
}

func calcBottomViewDistance(tree int, m *[][]int, x int, y int) (bottom int) {
	bottom = 0

	for i := y + 1; i < len((*m)); i++ {
		if tree > (*m)[x][i] {
			bottom++
		} else {
			bottom++
			return
		}
	}

	return
}

func calcLeftViewDistance(tree int, m *[][]int, x int, y int) (left int) {
	left = 0

	for i := x - 1; i >= 0; i-- {
		if tree > (*m)[i][y] {
			left++
		} else {
			left++
			return
		}
	}

	return
}

func calcRightViewDistance(tree int, m *[][]int, x int, y int) (right int) {
	right = 0

	for i := x + 1; i < len((*m)[0]); i++ {
		if tree > (*m)[i][y] {
			right++
		} else {
			right++
			return
		}
	}

	return
}

func findHighestScenicScore(m *[][]int) (highestScore int) {
	highestScore = 0

	for y, row := range (*m)[1 : len((*m))-1] {
		for x := range (*m)[y+1][1 : len(row)-1] {
			score := calcScenicScore((*m)[x+1][y+1], m, x+1, y+1)

			if score > highestScore {
				highestScore = score
			}
		}
	}

	return
}
