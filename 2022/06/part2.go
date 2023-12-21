package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	path := os.Args[1]

	f, err := os.Open(path)
	check(err)

	defer f.Close()

	scanner := bufio.NewScanner(f)

	datastream := ""

	for scanner.Scan() {
		datastream = scanner.Text()
	}

	_, index := findMarker(datastream)

	fmt.Println(index)
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func findMarker(datastream string) (bool, int) {
	datastreamLength := len(datastream)

	chunkSize := 14

	for i := range datastream {
		if i+chunkSize <= datastreamLength {
			chunk := datastream[i : i+chunkSize]

			if isMarker(chunk) {
				return true, strings.Index(datastream, chunk) + chunkSize
			}
		}
	}

	return false, -1
}

func isMarker(chunk string) bool {
	result := true

	for _, v := range chunk {
		if strings.Count(chunk, string(v)) > 1 {
			result = false
		}
	}

	return result
}
