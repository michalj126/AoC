package main

import (
	"bufio"
	"fmt"
	"os"
	"runtime"
	"strings"
	"time"
)

func main() {
	start := time.Now()

	path := os.Args[1]

	f, err := os.Open(path)
	Check(err)

	defer f.Close()

	scanner := bufio.NewScanner(f)

	grid := [][]string{}

	for scanner.Scan() {
		line := scanner.Text()

		grid = append(grid, strings.Split(line, ""))

	}

	counter := FindXMAS(grid)

	duration := time.Since(start)

	fmt.Println("XMAS words: ", counter)

	fmt.Println("Execution time: ", duration)
	PrintMemUsage()
}

func FindXMAS(grid [][]string) (counter int) {
	gh, gw := len(grid), len(grid[0])

	for py := 1; py < gh-1; py++ {
		for px := 1; px < gw-1; px++ {
			if grid[py][px] != "A" {
				continue
			}

			if CheckX(grid, px, py) {
				counter++
			}
		}
	}

	return
}

func CheckX(grid [][]string, x, y int) bool {
	a, b := false, false

	if grid[y-1][x-1] == "M" && grid[y+1][x+1] == "S" {
		a = true
	} else if grid[y-1][x-1] == "S" && grid[y+1][x+1] == "M" {
		a = true
	}

	if grid[y-1][x+1] == "M" && grid[y+1][x-1] == "S" {
		b = true
	} else if grid[y-1][x+1] == "S" && grid[y+1][x-1] == "M" {
		b = true
	}

	return a && b
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
