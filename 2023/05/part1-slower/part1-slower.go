package main

import (
	"bufio"
	"fmt"
	"math"
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

	// sum := 0
	seeds := [8][]int{}
	// maps := [][]int{}
	currentMap := -1
	seedTemp := []int{}

	for scanner.Scan() {
		line := scanner.Text()
		
		if strings.Contains(line, "seeds") {
			seedsi := ConvertToInt(strings.Split(strings.Split(line, ": ")[1], " "))
			// seeds = append(seeds, seedsi)
			// seeds = [7][len(seedsi)]int{}
			seeds[0] = seedsi
			// for _, _ = range seedsi {
			// 	seeds = append(seeds, []int{})
			// }
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
				for i := range seeds[currentMap-1] {
					if len(seeds[currentMap]) == 0 {
						seeds[currentMap] = append(seeds[currentMap], seeds[currentMap-1]...)
						seedTemp = seeds[currentMap-1]
						// seeds[currentMap] = make([]int, len(seeds[currentMap-1]))
					}

					// if 

					dest, state := FindDest(seedTemp[i], line)
					if !state {
						seeds[currentMap][i] = dest
					}
					// seeds[currentMap][i] = FindDest(seedTemp[i], line)
					// seeds[currentMap][i] = FindDest(seeds[currentMap-1][i], line)
					// seeds[currentMap][i] = FindDest(v, line)
				}
			}
		}
	}

	min := slices.Min(seeds[7])

	// for _, v := range seeds {
	// 	fmt.Println(v)
	// }

	// fmt.Println("Current map: ", currentMap)
	// fmt.Println("Seeds: ", seeds)
	fmt.Println("The lowest location number: ", min)

	duration := time.Since(start)

	fmt.Println("Execution time: ", duration)
	PrintMemUsage()
}

func FindDest(seed int, line string) (dest int, notFound bool) {
	dest = -1
	notFound = false
	splited := ConvertToInt(strings.Split(line, " "))
	d, s, r := splited[0], splited[1], splited[2]

	if seed >= s && seed < s+r {
		// dest = int(math.Abs(float64(d-s)))+seed
		dest = int(math.Abs(float64(seed-s)))+d
	} else {
		// dest = seed
		notFound = true
	}

	return
}
// func ParseLine(line string) (winningNumbers []int, drawnNumbers []int) {
// 	line = strings.ReplaceAll(line, "  ", " ")
// 	numbers := strings.Split(strings.Split(line, ": ")[1], " | ")
// 	winningNumbers = ConvertToInt(strings.Split(numbers[0], " "))
// 	drawnNumbers = ConvertToInt(strings.Split(numbers[1], " "))
//
// 	return
// }

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
