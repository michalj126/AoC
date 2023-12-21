package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Group struct {
	rucksacks []string
}

func main() {
	path := os.Args[1]

	f, err := os.Open(path)
	check(err)

	defer f.Close()

	groups := loadChunk(f, 3)

	sum := 0

	for _, group := range groups {
		reducedRucksackContent := reduceToUnique(group.rucksacks[0])

		for _, v := range reducedRucksackContent {
			if isBadge(group, v) {
				if v >= 97 {
					sum += int(v - 96)
				} else {
					sum += int(v - 38)
				}
			}
		}
	}

	fmt.Println(sum)
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func loadChunk(f *os.File, size int) []Group {
	scanner := bufio.NewScanner(f)

	groups := []Group{}
	rucksacks := []string{}

	line := 1

	for scanner.Scan() {
		rucksack := scanner.Text()

		if line > size {
			rucksacks = nil
			line = 1
		}

		rucksacks = append(rucksacks, rucksack)

		if line == 3 {
			group := Group{
				rucksacks: rucksacks,
			}
			groups = append(groups, group)
		}

		line++
	}

	return groups
}

func isBadge(group Group, char rune) bool {
	result := false

	for _, rucksack := range group.rucksacks[1:] {
		if strings.ContainsRune(rucksack, char) {
			result = true
		} else {
			return false
		}
	}

	return result
}

func reduceToUnique(input string) string {
	unique := ""

	for _, v := range input {
		if strings.ContainsRune(input, v) && !strings.ContainsRune(unique, v) {
			unique += string(v)
		}
	}

	return unique
}
