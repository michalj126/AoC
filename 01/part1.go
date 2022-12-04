package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	path := os.Args[1]

	f, err := os.Open(path)
	check(err)

	defer f.Close()

	scanner := bufio.NewScanner(f)

	totalCalories := 0

	mostCalories := 0

	for scanner.Scan() {
		line := scanner.Text()

		if line == "" {
			totalCalories = 0
		} else {
			itemCalories, err := strconv.Atoi(line)
			check(err)

			totalCalories += itemCalories

			if mostCalories <= totalCalories {
				mostCalories = totalCalories
			}
		}
	}

	fmt.Println(mostCalories)
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
