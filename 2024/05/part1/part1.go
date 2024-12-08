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

type Rule struct {
	x string
	y string
}

func main() {
	start := time.Now()

	path := os.Args[1]

	f, err := os.Open(path)
	Check(err)

	defer f.Close()

	scanner := bufio.NewScanner(f)

	sum := 0

	rules := map[string][]Rule{}

	for scanner.Scan() {
		line := scanner.Text()

		if strings.Contains(line, "|") {
			pair := strings.Split(line, "|")

			rules[pair[0]] = append(rules[pair[0]], Rule{x: pair[0], y: pair[1]})

			continue
		}

		if line == "" {
			continue
		}

		update := strings.Split(line, ",")

		correct := true

		for i, p := range update {
			t := update[:i]

			for _, p := range rules[p] {
				if slices.Contains(t, p.y) {
					correct = false
					break
				}
			}

			if !correct {
				break
			}
		}

		if correct {
			mps := update[len(update)/2]

			mpi, err := strconv.Atoi(mps)
			Check(err)

			sum += mpi
		}
	}

	duration := time.Since(start)

	fmt.Println("Sum of middle pages: ", sum)

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
