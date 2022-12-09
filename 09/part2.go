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

var tMoves = []Coordinates{{
	x: 0,
	y: 0,
}}

var rope = []Coordinates{}

func main() {
	path := os.Args[1]

	f, err := os.Open(path)
	check(err)

	defer f.Close()

	scanner := bufio.NewScanner(f)

	initRope(10)

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

func initRope(length int) {
	for i := 0; i < length; i++ {
		rope = append(rope, Coordinates{
			x: 0,
			y: 0,
		})
	}
}

func move(input string) {
	splitted := strings.Split(input, " ")

	if splitted[0] == "R" {
		steps, err := strconv.Atoi(splitted[1])
		check(err)

		for i := 0; i < steps; i++ {
			rope[0].x++

			for i := range rope {
				if len(rope)-1 > i {
					moveT(i + 1)
				}
			}
		}
	}

	if splitted[0] == "U" {
		steps, err := strconv.Atoi(splitted[1])
		check(err)

		for i := 0; i < steps; i++ {
			rope[0].y++

			for i := range rope {
				if len(rope)-1 > i {
					moveT(i + 1)
				}
			}
		}
	}

	if splitted[0] == "L" {
		steps, err := strconv.Atoi(splitted[1])
		check(err)

		for i := 0; i < steps; i++ {
			rope[0].x--

			for i := range rope {
				if len(rope)-1 > i {
					moveT(i + 1)
				}
			}
		}
	}

	if splitted[0] == "D" {
		steps, err := strconv.Atoi(splitted[1])
		check(err)

		for i := 0; i < steps; i++ {
			rope[0].y--

			for i := range rope {
				if len(rope)-1 > i {
					moveT(i + 1)
				}
			}
		}
	}

}

func moveT(index int) {
	if !isTouch(index) {
		if rope[index].x < rope[index-1].x && rope[index].y == rope[index-1].y {
			rope[index].x++
		}

		if rope[index].x > rope[index-1].x && rope[index].y == rope[index-1].y {
			rope[index].x--
		}

		if rope[index].y < rope[index-1].y && rope[index].x == rope[index-1].x {
			rope[index].y++
		}

		if rope[index].y > rope[index-1].y && rope[index].x == rope[index-1].x {
			rope[index].y--
		}

		if rope[index].x < rope[index-1].x && rope[index].y < rope[index-1].y {
			rope[index].x++
			rope[index].y++
		}

		if rope[index].x < rope[index-1].x && rope[index].y > rope[index-1].y {
			rope[index].x++
			rope[index].y--
		}

		if rope[index].x > rope[index-1].x && rope[index].y > rope[index-1].y {
			rope[index].x--
			rope[index].y--
		}

		if rope[index].x > rope[index-1].x && rope[index].y < rope[index-1].y {
			rope[index].x--
			rope[index].y++
		}

		if !isTMoveContain() {
			tMoves = append(tMoves, rope[index])
		}
	}
}

func isTMoveContain() bool {
	for _, v := range tMoves {
		if v.x == rope[len(rope)-1].x && v.y == rope[len(rope)-1].y {
			return true
		}
	}

	return false
}

func isTouch(index int) bool {
	if rope[index].x == rope[index-1].x && rope[index].y == rope[index-1].y {
		return true
	}

	if rope[index].x+1 == rope[index-1].x && rope[index].y == rope[index-1].y {
		return true
	}

	if rope[index].x-1 == rope[index-1].x && rope[index].y == rope[index-1].y {
		return true
	}

	if rope[index].y+1 == rope[index-1].y && rope[index].x == rope[index-1].x {
		return true
	}

	if rope[index].y-1 == rope[index-1].y && rope[index].x == rope[index-1].x {
		return true
	}

	if rope[index].y+1 == rope[index-1].y && rope[index].x+1 == rope[index-1].x {
		return true
	}

	if rope[index].y+1 == rope[index-1].y && rope[index].x-1 == rope[index-1].x {
		return true
	}

	if rope[index].y-1 == rope[index-1].y && rope[index].x-1 == rope[index-1].x {
		return true
	}

	if rope[index].y-1 == rope[index-1].y && rope[index].x+1 == rope[index-1].x {
		return true
	}

	return false
}
