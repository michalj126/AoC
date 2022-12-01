package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

func main() {
	path := os.Args[1]

	f, err := os.Open(path)
	check(err)

	scanner := bufio.NewScanner(f)

	totalCaloriesCollection := []int{}

	totalCalories := 0

	for scanner.Scan() {
		line := scanner.Text()

		if line == "" {
			totalCaloriesCollection = append(totalCaloriesCollection, totalCalories)
			totalCalories = 0
		} else {
			itemCalories, err := strconv.Atoi(line)
			check(err)

			totalCalories += itemCalories
		}
	}

	totalCaloriesCollection = append(totalCaloriesCollection, totalCalories)

	sort.Ints(totalCaloriesCollection)

	totalTopThree := 0

	for i := 1; i < 4; i++ {
		totalTopThree += totalCaloriesCollection[len(totalCaloriesCollection)-i]
	}

	fmt.Println(totalTopThree)
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
