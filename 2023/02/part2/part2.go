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

type Game struct {
	id   int
	sets []Set
}

type Set struct {
	r int
	g int
	b int
}

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

		game := ParseLine(line)

		result := FindFewestNumberOfCubesOfEachColor(game.sets)

		sum += result.r * result.g * result.b
	}

	fmt.Println("Sum of IDs: ", sum)

	duration := time.Since(start)

	fmt.Println("Execution time: ", duration)
	PrintMemUsage()
}

func ParseLine(line string) (game Game) {
	lineSeparated := strings.Split(line, ":")
	id, err := strconv.Atoi(strings.Split(lineSeparated[0], " ")[1])
	Check(err)
	game.id = id
	sets := strings.Split(lineSeparated[1], ";")

	for si, set := range sets {
		game.sets = append(game.sets, Set{})

		for _, cubes := range strings.Split(set, ",") {
			numberOfCubes, err := strconv.Atoi(strings.Split(cubes, " ")[1])
			Check(err)
			if strings.Contains(cubes, "red") {
				game.sets[si].r = numberOfCubes
			} else if strings.Contains(cubes, "green") {
				game.sets[si].g = numberOfCubes
			} else if strings.Contains(cubes, "blue") {
				game.sets[si].b = numberOfCubes
			}
		}

	}

	return
}

func FindFewestNumberOfCubesOfEachColor(sets []Set) (setResult Set) {
	setResult.r, setResult.g, setResult.b = 0, 0, 0

	for _, set := range sets {
		if set.r > setResult.r {
			setResult.r = set.r
		}

		if set.g > setResult.g {
			setResult.g = set.g
		}

		if set.b > setResult.b {
			setResult.b = set.b
		}
	}

	return
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
