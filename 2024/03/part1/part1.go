package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
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

		re := regexp.MustCompile(`mul\((\d+,\d+)\)`)

		instructions := re.FindAllString(line, -1)

		for _, instruction := range instructions {
			values := strings.Split(instruction[4:len(instruction)-1], ",")

			a, err := strconv.Atoi(values[0])
			Check(err)

			b, err := strconv.Atoi(values[1])
			Check(err)

			sum += a * b
		}
	}

	duration := time.Since(start)

	fmt.Println("Sum of the multiplications: ", sum)

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
