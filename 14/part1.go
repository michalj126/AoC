package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

type Path struct {
	coords []Coords
}

type Coords struct {
	x int
	y int
}

var paths = []Path{}

type Cave struct {
	minX    int
	maxX    int
	minY    int
	maxY    int
	caveMap [][]string
}

var cave = Cave{
	minX: -1,
	maxX: 0,
	minY: 0,
	maxY: 0,
}

func main() {
	path := os.Args[1]

	f, err := os.Open(path)
	check(err)

	defer f.Close()

	parse(f)

	createMap()

	result := run()

	renderMap()

	fmt.Println(result)
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func run() (counter int) {
	source := Coords{
		x: 500 - cave.minX,
		y: 0,
	}

	xTemp := source.x
	yTemp := source.y

	counter = -1

	end := false

	for ok := true; ok; ok = end == false {
		xTemp = source.x
		yTemp = source.y

		counter++

		for {
			if xTemp-1 < 0 {
				end = true
				break
			}

			if xTemp+1 > len(cave.caveMap[0])-1 {
				end = true
				break
			}

			yTemp++

			if cave.caveMap[yTemp+1][xTemp] == "#" || cave.caveMap[yTemp+1][xTemp] == "o" {
				if cave.caveMap[yTemp+1][xTemp-1] != "#" && cave.caveMap[yTemp+1][xTemp-1] != "o" {
					xTemp--
				} else if cave.caveMap[yTemp+1][xTemp+1] != "#" && cave.caveMap[yTemp+1][xTemp+1] != "o" {
					xTemp++
				} else {
					cave.caveMap[yTemp][xTemp] = "o"
					break
				}
			}
		}
	}

	return
}

func renderMap() {
	for i := range cave.caveMap {
		for _, v := range cave.caveMap[i] {
			fmt.Print(v)
		}
		fmt.Println()
	}
}

func createMap() {
	width := cave.maxX - cave.minX
	height := cave.maxY - cave.minY

	for i := 0; i < height+1; i++ {
		cave.caveMap = append(cave.caveMap, []string{})

		for j := 0; j < width+1; j++ {
			cave.caveMap[i] = append(cave.caveMap[i], ".")
		}
	}

	for _, p := range paths {
		for i := 0; i < len(p.coords)-1; i++ {
			if p.coords[i].x == p.coords[i+1].x {
				x := p.coords[i].x - cave.minX
				y := p.coords[i].y
				yStart := p.coords[i].y
				yStop := p.coords[i+1].y

				for i := 0; i <= int(math.Abs(float64(yStart)-float64(yStop))); i++ {
					cave.caveMap[y][x] = "#"
					if yStart < yStop {
						y++
					} else {
						y--
					}
				}
			}

			if p.coords[i].y == p.coords[i+1].y {
				x := p.coords[i].x - cave.minX
				y := p.coords[i].y
				xStart := p.coords[i].x
				xStop := p.coords[i+1].x

				for i := 0; i <= int(math.Abs(float64(xStart)-float64(xStop))); i++ {
					cave.caveMap[y][x] = "#"
					if xStart < xStop {
						x++
					} else {
						x--
					}
				}
			}
		}
	}
}

func parse(f *os.File) {
	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		path := scanner.Text()

		coords := strings.Split(path, " -> ")

		paths = append(paths, Path{})

		for _, cv := range coords {
			coord := strings.Split(cv, ",")

			x, err := strconv.Atoi(coord[0])
			check(err)

			y, err := strconv.Atoi(coord[1])
			check(err)

			if x > cave.maxX {
				cave.maxX = x
			}

			if cave.minX == -1 || x < cave.minX {
				cave.minX = x
			}

			if y > cave.maxY {
				cave.maxY = y
			}

			paths[len(paths)-1].coords = append(paths[len(paths)-1].coords, Coords{
				x: x,
				y: y,
			})
		}
	}
}
