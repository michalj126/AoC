package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Pair struct {
	begin int
	end   int
}

func main() {
	path := os.Args[1]

	f, err := os.Open(path)
	check(err)

	defer f.Close()

	scanner := bufio.NewScanner(f)

	sum := 0

	for scanner.Scan() {
		line := scanner.Text()

		firstPair, secondPair := splitPairs(line)

		if contains(firstPair, secondPair) || contains(secondPair, firstPair) {
			sum++
		}
	}

	fmt.Println(sum)
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func splitPairs(input string) (Pair, Pair) {
	splittedPairs := strings.Split(input, ",")

	splittedFirstPair := strings.Split(splittedPairs[0], "-")
	splittedSecondPair := strings.Split(splittedPairs[1], "-")

	firstPair := Pair{
		begin: str2Int(splittedFirstPair[0]),
		end:   str2Int(splittedFirstPair[1]),
	}

	secondPair := Pair{
		begin: str2Int(splittedSecondPair[0]),
		end:   str2Int(splittedSecondPair[1]),
	}

	return firstPair, secondPair
}

func str2Int(input string) int {
	i, err := strconv.Atoi(input)
	check(err)

	return i
}

func contains(x Pair, y Pair) bool {
	if y.begin >= x.begin && y.end <= x.end {
		return true
	}

	return false
}
