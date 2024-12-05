package main

import (
	"bufio"
	"fmt"
	"os"
	"runtime"
	"strconv"
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

	sum := 0

	left := []int{}
	multipliers := make(map[int]int)

	for scanner.Scan() {
		line := scanner.Text()

		pair := strings.Split(line, "   ")

		leftNumber, err := strconv.Atoi(string(pair[0]))
		Check(err)
		rightNumber, err := strconv.Atoi(string(pair[1]))
		Check(err)

		left = append(left, leftNumber)
		multipliers[rightNumber] = multipliers[rightNumber] + 1
	}

	for _, v := range left {
		sum += v * multipliers[v]
	}

	fmt.Println("Total distance: ", sum)

	duration := time.Since(start)

	fmt.Println("Execution time: ", duration)
	PrintMemUsage()
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
