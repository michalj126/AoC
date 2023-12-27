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

	raceTime := 0
	raceDistance := 0

	for scanner.Scan() {
		line := scanner.Text()

		splited := strings.ReplaceAll(strings.Split(line, ":")[1], " ", "")
		iV, err := strconv.Atoi(splited)
		Check(err)

		if strings.Contains(line, "Time") {
			raceTime = iV
		} else {
			raceDistance = iV
		}
	}

	waysToWin := FindWaysToWin(raceTime, raceDistance)

	fmt.Println("Ways to win: ", waysToWin)

	duration := time.Since(start)

	fmt.Println("Execution time: ", duration)
	PrintMemUsage()
}

func FindWaysToWin(rT int, rD int) (waysToWin int) {
	distance := 0
	pushTime := 0

	for distance <= rD {
		pushTime++
		distance = (rT - pushTime) * pushTime
	}

	waysToWin = rT - (pushTime*2 - 1)

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
