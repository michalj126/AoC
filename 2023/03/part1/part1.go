package main

import (
	"bufio"
	"fmt"
	"os"
	"runtime"
	"strconv"
	"time"
)

type EngineSchematic struct {
	parts   []Part
	symbols []Symbol
}

type Part struct {
	x          int
	y          int
	partNumber string
}

type Symbol struct {
	x      int
	y      int
	symbol string
}

func main() {
	start := time.Now()

	path := os.Args[1]

	f, err := os.Open(path)
	Check(err)

	defer f.Close()

	scanner := bufio.NewScanner(f)

	engineSchematic := EngineSchematic{}

	for i := 0; scanner.Scan(); i++ {
		line := scanner.Text()

		FindEnginePart(&engineSchematic, line, i)
	}

	sum := SumUpNumbersOfAllEngineParts(engineSchematic)

	fmt.Println("Sum of all of the part numbers: ", sum)

	duration := time.Since(start)

	fmt.Println("Execution time: ", duration)
	PrintMemUsage()
}

func SumUpNumbersOfAllEngineParts(engineSchematic EngineSchematic) (sum int) {
	sum = 0

	for _, p := range engineSchematic.parts {
		if IsEnginePart(p, engineSchematic.symbols) {
			partNumber, err := strconv.Atoi(p.partNumber)
			Check(err)

			sum += partNumber
		}
	}

	return
}

func IsEnginePart(part Part, symbols []Symbol) bool {
	l, t, r, b := part.x-1, part.y-1, part.x+len(part.partNumber), part.y+1

	for _, s := range symbols {
		if (s.x >= l && s.x <= r) && (s.y >= t && s.y <= b) {
			return true
		}
	}

	return false
}

func FindEnginePart(engineSchematic *EngineSchematic, line string, lineNumber int) {
	isPart := false
	x := 0
	y := 0

	for ri, rune := range line {
		if rune >= 48 && rune <= 57 {
			if !isPart {
				isPart = true
				x = ri
				y = lineNumber

				engineSchematic.parts = append(engineSchematic.parts, Part{
					x:          x,
					y:          y,
					partNumber: "",
				})
			}

			engineSchematic.parts[len(engineSchematic.parts)-1].partNumber += string(rune)
		} else {
			isPart = false
		}

		if (rune < 48 || rune > 57) && rune != 46 {
			engineSchematic.symbols = append(engineSchematic.symbols, Symbol{
				x:      ri,
				y:      lineNumber,
				symbol: string(rune),
			})
		}
	}
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
