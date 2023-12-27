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

	seeds := []int{}
	maps := make(map[int][][]int)
	currentMap := -1

	for scanner.Scan() {
		line := scanner.Text()

		if strings.Contains(line, "seeds") {
			seeds = ConvertToInt(strings.Split(strings.Split(line, ": ")[1], " "))
		} else {
			switch line {
			case "seed-to-soil map:":
				currentMap = 1
			case "soil-to-fertilizer map:":
				currentMap = 2
			case "fertilizer-to-water map:":
				currentMap = 3
			case "water-to-light map:":
				currentMap = 4
			case "light-to-temperature map:":
				currentMap = 5
			case "temperature-to-humidity map:":
				currentMap = 6
			case "humidity-to-location map:":
				currentMap = 7
			case "":
			default:
				iMap := ConvertToInt(strings.Split(line, " "))
				maps[currentMap] = append(maps[currentMap], iMap)
			}
		}
	}

	min := -1

	for i := 0; i < len(seeds); i += 2 {
		for j := seeds[i]; j < seeds[i]+seeds[i+1]; j++ {
			dest := j

			for k := 0; k < len(maps); k++ {
				destTemp := dest

				for _, v := range maps[k+1] {
					d, state := FindDest(destTemp, v)

					if !state {
						dest = d
					}
				}
			}

			if min == -1 || dest < min {
				min = dest
			}
		}
	}

	fmt.Println("The lowest location number: ", min)

	duration := time.Since(start)

	fmt.Println("Execution time: ", duration)
	PrintMemUsage()
}

func FindDest(seed int, m []int) (dest int, notFound bool) {
	dest = -1
	notFound = false
	d, s, r := m[0], m[1], m[2]

	if seed >= s && seed < s+r {
		dest = int(math.Abs(float64(seed-s))) + d
	} else {
		notFound = true
	}

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
