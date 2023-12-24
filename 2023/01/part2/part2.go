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

	stringNumbers := [9]string{"one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}

	sum := 0

	for scanner.Scan() {
		line := scanner.Text()
		firstNumber := ""
		secondNumber := ""

		for ri, rune := range line {
			if rune >= 48 && rune <= 57 {
				if firstNumber == "" {
					firstNumber = string(rune)
				} else {
					secondNumber = string(rune)
				}
			} else {
				for sni, stringNumber := range stringNumbers {
					index := strings.Index(line[ri:], stringNumber)

					if index == 0 {
						if firstNumber == "" {
							firstNumber = strconv.Itoa(sni + 1)
						} else {
							secondNumber = strconv.Itoa(sni + 1)
						}
						break
					}
				}
			}
		}

		if secondNumber == "" {
			secondNumber = firstNumber
		}

		lineValue := firstNumber + secondNumber
		i, err := strconv.Atoi(lineValue)
		Check(err)

		sum += i
	}

	fmt.Println("Sum of all calibration values: ", sum)

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
