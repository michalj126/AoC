package main

import (
	"bufio"
	"fmt"
	"os"
	"runtime"
	"slices"
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

	for scanner.Scan() {
		line := scanner.Text()

		sum += CalculatePoints(ParseLine(line))
	}

	fmt.Println("Total points: ", sum)

	duration := time.Since(start)

	fmt.Println("Execution time: ", duration)
	PrintMemUsage()
}

func CalculatePoints(winningNumbers []int, drawnNumbers []int) (points int) {
	points = 0

	for _, dn := range drawnNumbers {
		if slices.Contains(winningNumbers, dn) {
			if points == 0 {
				points = 1
			} else {
				points *= 2
			}
		}
	}

	return
}

func ParseLine(line string) (winningNumbers []int, drawnNumbers []int) {
	line = strings.ReplaceAll(line, "  ", " ")
	numbers := strings.Split(strings.Split(line, ": ")[1], " | ")
	winningNumbers = ConvertToInt(strings.Split(numbers[0], " "))
	drawnNumbers = ConvertToInt(strings.Split(numbers[1], " "))

	return
}

func ConvertToInt(input []string) []int {
	result := make([]int, len(input))

	for i, v := range input {
		vi, err := strconv.Atoi(v)
		Check(err)

		result[i] = vi
	}

	return result
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
