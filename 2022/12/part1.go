package main

import (
	"bufio"
	"fmt"
	"os"
)

type Square struct {
	x         int
	y         int
	height    int
	mark      string
	direction string
	visited   bool
}

type CurrentPos struct {
	x int
	y int
}

var heightmap = [][]Square{}

func main() {
	path := os.Args[1]

	f, err := os.Open(path)
	check(err)

	defer f.Close()

	parse(f)

	counter := findPath(CurrentPos{0, 0})

	fmt.Println(counter)
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

var history = []Square{}

func findPath(cp CurrentPos) (counter int) {
	d := []Square{start}

	counter = 0

	found := false

	for ok := true; ok; ok = found != true {
		counter++

		dcopy := make([]Square, len(d))
		copy(dcopy, d)

		d = nil

		for _, v := range dcopy {
			if v.x+1 < len(heightmap) {
				right := heightmap[v.x+1][v.y]
				right.direction = ">"

				if v.direction != "<" && right.height-v.height <= 1 && !heightmap[v.x+1][v.y].visited {
					d = append(d, right)
					heightmap[v.x+1][v.y].visited = true
				}
			}

			if v.x-1 >= 0 {
				left := heightmap[v.x-1][v.y]
				left.direction = "<"

				if v.direction != ">" && left.height-v.height <= 1 && !heightmap[v.x-1][v.y].visited {
					d = append(d, left)
					heightmap[v.x-1][v.y].visited = true
				}
			}

			if v.y+1 < len(heightmap[0]) {
				down := heightmap[v.x][v.y+1]
				down.direction = "V"

				if v.direction != "^" && down.height-v.height <= 1 && !heightmap[v.x][v.y+1].visited {
					d = append(d, down)
					heightmap[v.x][v.y+1].visited = true
				}
			}

			if v.y-1 >= 0 {
				up := heightmap[v.x][v.y-1]
				up.direction = "^"

				if v.direction != "V" && up.height-v.height <= 1 && !heightmap[v.x][v.y-1].visited {
					d = append(d, up)
					heightmap[v.x][v.y-1].visited = true
				}
			}

			if v.mark == "E" {
				counter--
				found = true
				break
			}
		}

	}

	return
}

func render() {
	for y := range heightmap[0] {
		for x := range heightmap {
			fmt.Print(heightmap[x][y].direction)
		}
		fmt.Println()
	}
}

var start = Square{}

func parse(f *os.File) {
	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		line := scanner.Text()

		for i, square := range line {
			height := 0

			if string(square) == "S" {
				height = 0
			} else if string(square) == "E" {
				height = 25
			} else {
				height = int(square - 97)
			}

			if len(heightmap) == i {
				heightmap = append(heightmap, []Square{})
			}

			heightmap[i] = append(heightmap[i], Square{
				x:         i,
				y:         len(heightmap[i]),
				height:    height,
				mark:      string(square),
				direction: ".",
				visited:   false,
			})

			if string(square) == "S" {
				start = heightmap[i][len(heightmap[i])-1]
			}
		}
	}
}
