package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Coordinates struct {
	x int
	y int
}

var hPos = Coordinates{
	x: 0,
	y: 0,
}

var tPos = Coordinates{
	x: 0,
	y: 0,
}

var tMoves = []Coordinates{{
	x: 0,
	y: 0,
}}

func main() {
	path := os.Args[1]

	f, err := os.Open(path)
	check(err)

	defer f.Close()

	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		line := scanner.Text()

		move(line)
	}

	fmt.Println(len(tMoves))
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func move(input string) {
	splitted := strings.Split(input, " ")

	if splitted[0] == "R" {
		steps, err := strconv.Atoi(splitted[1])
		check(err)

		for i := 0; i < steps; i++ {
			hPos.x++
			moveT()
		}
	}

	if splitted[0] == "U" {
		steps, err := strconv.Atoi(splitted[1])
		check(err)

		for i := 0; i < steps; i++ {
			hPos.y++
			moveT()
		}
	}

	if splitted[0] == "L" {
		steps, err := strconv.Atoi(splitted[1])
		check(err)

		for i := 0; i < steps; i++ {
			hPos.x--
			moveT()
		}
	}

	if splitted[0] == "D" {
		steps, err := strconv.Atoi(splitted[1])
		check(err)

		for i := 0; i < steps; i++ {
			hPos.y--
			moveT()
		}
	}

}

func moveT() {
	if !isTouch() {
		if tPos.x < hPos.x && tPos.y == hPos.y {
			tPos.x++
		}

		if tPos.x > hPos.x && tPos.y == hPos.y {
			tPos.x--
		}

		if tPos.y < hPos.y && tPos.x == hPos.x {
			tPos.y++
		}

		if tPos.y > hPos.y && tPos.x == hPos.x {
			tPos.y--
		}

		if tPos.x < hPos.x && tPos.y < hPos.y {
			tPos.x++
			tPos.y++
		}

		if tPos.x < hPos.x && tPos.y > hPos.y {
			tPos.x++
			tPos.y--
		}

		if tPos.x > hPos.x && tPos.y > hPos.y {
			tPos.x--
			tPos.y--
		}

		if tPos.x > hPos.x && tPos.y < hPos.y {
			tPos.x--
			tPos.y++
		}

		if !isTMoveContain() {
			tMoves = append(tMoves, tPos)
		}
	}
}

func isTMoveContain() bool {
	for _, v := range tMoves {
		if v.x == tPos.x && v.y == tPos.y {
			return true
		}
	}

	return false
}

func isTouch() bool {
	if tPos.x == hPos.x && tPos.y == hPos.y {
		return true
	}

	if tPos.x+1 == hPos.x && tPos.y == hPos.y {
		return true
	}

	if tPos.x-1 == hPos.x && tPos.y == hPos.y {
		return true
	}

	if tPos.y+1 == hPos.y && tPos.x == hPos.x {
		return true
	}

	if tPos.y-1 == hPos.y && tPos.x == hPos.x {
		return true
	}

	if tPos.y+1 == hPos.y && tPos.x+1 == hPos.x {
		return true
	}

	if tPos.y+1 == hPos.y && tPos.x-1 == hPos.x {
		return true
	}

	if tPos.y-1 == hPos.y && tPos.x-1 == hPos.x {
		return true
	}

	if tPos.y-1 == hPos.y && tPos.x+1 == hPos.x {
		return true
	}

	return false
}
