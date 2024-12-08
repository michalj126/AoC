package main

import (
	"bufio"
	"fmt"
	"os"
	"runtime"
	"strings"
	"time"
)

type Direction int

const (
	t Direction = iota
	b
	l
	r
	tl
	tr
	bl
	br
)

func main() {
	start := time.Now()

	path := os.Args[1]

	f, err := os.Open(path)
	Check(err)

	defer f.Close()

	scanner := bufio.NewScanner(f)

	grid := [][]string{}
	word := "XMAS"

	for scanner.Scan() {
		line := scanner.Text()

		grid = append(grid, strings.Split(line, ""))

	}

	counter := FindWords(grid, word)

	duration := time.Since(start)

	fmt.Println("XMAS words: ", counter)

	fmt.Println("Execution time: ", duration)
	PrintMemUsage()
}

func FindWords(grid [][]string, word string) (counter int) {
	gh, gw := len(grid), len(grid[0])

	for py := 0; py < gh; py++ {
		for px := 0; px < gw; px++ {
			if grid[py][px] != string(word[0]) {
				continue
			}

			for d := t; d <= br; d++ {
				if Search(grid, word, px, py, d) {
					counter++
				}
			}
		}
	}

	return
}

func Search(grid [][]string, word string, x, y int, direction Direction) bool {
	gh, gw, lw := len(grid), len(grid[0]), len(word)

	switch direction {
	case t:
		if y+1-lw < 0 {
			return false
		}

		for i, w := range word {
			if grid[y-i][x] != string(w) {
				return false
			}
		}
	case b:
		if y+lw > gh {
			return false
		}

		for i, w := range word {
			if grid[y+i][x] != string(w) {
				return false
			}
		}
	case l:
		if x+1-lw < 0 {
			return false
		}

		for i, w := range word {
			if grid[y][x-i] != string(w) {
				return false
			}
		}
	case r:
		if x+lw > gw {
			return false
		}

		for i, w := range word {
			if grid[y][x+i] != string(w) {
				return false
			}
		}
	case tl:
		if y+1-lw < 0 || x+1-lw < 0 {
			return false
		}

		for i, w := range word {
			if grid[y-i][x-i] != string(w) {
				return false
			}
		}
	case tr:
		if y+1-lw < 0 || x+lw > gw {
			return false
		}

		for i, w := range word {
			if grid[y-i][x+i] != string(w) {
				return false
			}
		}
	case bl:
		if y+lw > gh || x+1-lw < 0 {
			return false
		}

		for i, w := range word {
			if grid[y+i][x-i] != string(w) {
				return false
			}
		}
	case br:
		if y+lw > gh || x+lw > gw {
			return false
		}

		for i, w := range word {
			if grid[y+i][x+i] != string(w) {
				return false
			}
		}
	}

	return true
}

func Check(e error) {
	if e != nil {
		panic(e)
	}
}

func PrintMemUsage() {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	fmt.Printf("Alloc = %v MiB", m.Alloc/1024/1024)
	fmt.Printf(" | %v KiB", m.Alloc/1024)
	fmt.Printf("\tTotalAlloc = %v MiB", m.TotalAlloc/1024/1024)
	fmt.Printf(" | %v KiB", m.TotalAlloc/1024)
	fmt.Printf("\tSys = %v MiB", m.Sys/1024/1024)
	fmt.Printf(" | Sys = %v KiB", m.Sys/1024)
	fmt.Printf("\tNumGC = %v\n", m.NumGC)
}
