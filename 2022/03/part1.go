package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Rucksack struct {
	firstCompartment  string
	secondCompartment string
}

func main() {
	path := os.Args[1]

	f, err := os.Open(path)
	check(err)

	defer f.Close()

	scanner := bufio.NewScanner(f)

	sum := 0

	for scanner.Scan() {
		rucksackContent := scanner.Text()

		rucksack := splitRucksackContent(rucksackContent)

		for _, v := range reduceToUnique(rucksack.firstCompartment) {
			result := strings.ContainsRune(rucksack.secondCompartment, v)

			if result {
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

func splitRucksackContent(rucksackContent string) Rucksack {
	half := len(rucksackContent) / 2
	rucksack := Rucksack{
		firstCompartment:  rucksackContent[:half],
		secondCompartment: rucksackContent[half:],
	}

	return rucksack
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
