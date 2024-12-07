package main

import (
	"bufio"
	"fmt"
	"math"
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

	for scanner.Scan() {
		line := scanner.Text()

		levels := []int{}

		for _, rune := range strings.Split(line, " ") {
			level, err := strconv.Atoi(rune)
			Check(err)

			levels = append(levels, level)
		}

		result, index := IsSafe(levels)

		if !result {
			reducedLevels := make([]int, len(levels))
			copy(reducedLevels, levels)
			reducedLevels = append(reducedLevels[:index], reducedLevels[index+1:]...)
			result, _ = IsSafe(reducedLevels)
		}

		if !result {
			reducedLevels := make([]int, len(levels))
			copy(reducedLevels, levels)
			reducedLevels = append(reducedLevels[:index+1], reducedLevels[index+2:]...)
			result, _ = IsSafe(reducedLevels)
		}

		if !result && index-1 >= 0 {
			reducedLevels := make([]int, len(levels))
			copy(reducedLevels, levels)
			reducedLevels = append(reducedLevels[:index-1], reducedLevels[index:]...)
			result, _ = IsSafe(reducedLevels)
		}

		if result {
			sum++
		}
	}

	duration := time.Since(start)

	fmt.Println("Safe reports: ", sum)

	fmt.Println("Execution time: ", duration)
	PrintMemUsage()
}

func IsSafe(levels []int) (bool, int) {
	init := levels[0] - levels[1]
	if init == 0 {
		return false, 0
	}

	isIncrease := init < 0

	for i := 0; i < len(levels)-1; i++ {
		delta := levels[i] - levels[i+1]

		if delta == 0 {
			return false, i
		}

		if isIncrease && delta > 0 {
			return false, i
		}

		if !isIncrease && delta < 0 {
			return false, i
		}

		if delta < 0 {
			delta = int(math.Abs(float64(delta)))
		}

		if delta > 3 {
			return false, i
		}
	}

	return true, 0
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
