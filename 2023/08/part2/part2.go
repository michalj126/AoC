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

type Ghost struct {
	Node  string
	Steps int
}

func main() {
	start := time.Now()

	path := os.Args[1]

	f, err := os.Open(path)
	Check(err)

	defer f.Close()

	scanner := bufio.NewScanner(f)

	nodes := make(map[string]Node)
	ghosts := []Ghost{}
	instructions := list.New()

	for i := 0; scanner.Scan(); i++ {
		line := scanner.Text()

		if i == 0 {
			for _, rune := range line {
				instructions.PushBack(rune)
			}
		}

		if i > 1 {
			fields := strings.FieldsFunc(line, func(r rune) bool {
				return !unicode.IsLetter(r) && !unicode.IsDigit(r)
			})

			nodes[fields[0]] = Node{
				Left:  fields[1],
				Right: fields[2],
			}

			if strings.HasSuffix(fields[0], "A") {
				ghosts = append(ghosts, Ghost{Node: fields[0], Steps: 0})
			}
		}
	}

	for i, g := range ghosts {
		node := g.Node
		instruction := instructions.Front()
		steps := 0

		for !strings.HasSuffix(node, "Z") {
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

		ghosts[i].Steps = steps
	}

	lcm := ghosts[0].Steps

	for i := 1; i < len(ghosts); i++ {
		lcm = LCM(lcm, ghosts[i].Steps)
	}

	fmt.Println("Steps: ", lcm)

	duration := time.Since(start)

	fmt.Println("Execution time: ", duration)
	PrintMemUsage()
}

func GCD(a, b int) (gcd int) {
	for b != 0 {
		temp := a
		a = b
		b = temp % b
	}

	gcd = a

	return
}

func LCM(a, b int) (lcm int) {
	lcm = a * b / GCD(a, b)

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
