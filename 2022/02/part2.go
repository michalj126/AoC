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

	totalScore := 0

	for scanner.Scan() {
		line := scanner.Text()

		round := strings.Split(line, " ")

		if round[1] == "Y" {
			if round[0] == "A" {
				round[1] = "X"
			} else if round[0] == "B" {
				round[1] = "Y"
			} else {
				round[1] = "Z"
			}
		} else if round[1] == "X" {
			if round[0] == "A" {
				round[1] = "Z"
			} else if round[0] == "B" {
				round[1] = "X"
			} else {
				round[1] = "Y"
			}
		} else {
			if round[0] == "A" {
				round[1] = "Y"
			} else if round[0] == "B" {
				round[1] = "Z"
			} else {
				round[1] = "X"
			}
		}

		if round[0] == "A" {
			if round[1] == "Y" {
				totalScore += 6
			} else if round[1] == "X" {
				totalScore += 3
			} else {
				totalScore += 0
			}
		} else if round[0] == "B" {
			if round[1] == "Y" {
				totalScore += 3
			} else if round[1] == "X" {
				totalScore += 0
			} else {
				totalScore += 6
			}
		} else {
			if round[1] == "Y" {
				totalScore += 0
			} else if round[1] == "X" {
				totalScore += 6
			} else {
				totalScore += 3
			}
		}

		if round[1] == "X" {
			totalScore += 1
		} else if round[1] == "Y" {
			totalScore += 2
		} else {
			totalScore += 3
		}
	}

	fmt.Println(totalScore)
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
