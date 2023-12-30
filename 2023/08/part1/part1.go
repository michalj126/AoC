package main

import (
	"bufio"
	"container/list"
	"fmt"
	"os"
	"runtime"
	"strconv"
	"strings"
	"time"
	"unicode"
)

type Node struct {
	Left  string
	Right string
}

func main() {
	start := time.Now()

	path := os.Args[1]

	f, err := os.Open(path)
	Check(err)

	defer f.Close()

	scanner := bufio.NewScanner(f)

	nodes := make(map[string]Node)
	instructions := list.New()
	startNode := "AAA"
	endNode := "ZZZ"

	for i := 0; scanner.Scan(); i++ {
		line := scanner.Text()

		if i == 0 {
			for _, rune := range line {
				instructions.PushBack(rune)
			}
		}

		if i > 1 {
			fields := strings.FieldsFunc(line, func(r rune) bool {
				return !unicode.IsLetter(r)
			})

			nodes[fields[0]] = Node{
				Left:  fields[1],
				Right: fields[2],
			}
		}
	}

	node := startNode
	instruction := instructions.Front()
	steps := 0

	for node != endNode {
		if instruction == nil {
			instruction = instructions.Front()
		}

		if instruction.Value == int32(82) {
			node = nodes[node].Right
		} else {
			node = nodes[node].Left
		}

		instruction = instruction.Next()

		steps++
	}

	fmt.Println("Steps: ", steps)

	duration := time.Since(start)

	fmt.Println("Execution time: ", duration)
	PrintMemUsage()
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
