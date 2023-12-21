package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Stack struct {
	crates []string
}

func main() {
	path := os.Args[1]

	f, err := os.Open(path)
	check(err)

	defer f.Close()

	scanner := bufio.NewScanner(f)

	stacks := []Stack{}

	firstPartEnd := false

	for scanner.Scan() {
		line := scanner.Text()

		if len(line) == 0 {
			firstPartEnd = true
		}

		if !firstPartEnd && string(line[1]) != "1" {
			splittedStacks := splitStacks(line)

			for i, v := range splittedStacks {
				if len(stacks) == i {
					stacks = append(stacks, Stack{})
				}

				if v != "   " {
					stacks[i].crates = append(stacks[i].crates, v)
				}
			}
		}

		if len(line) != 0 && string(line[0]) == "m" {
			command := strings.Split(line, " ")

			amount, err := strconv.Atoi(command[1])
			check(err)
			from, err := strconv.Atoi(command[3])
			check(err)
			to, err := strconv.Atoi(command[5])
			check(err)

			moveCrate(stacks, amount, from-1, to-1)
		}
	}

	printResult(stacks)
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func printResult(stacks []Stack) {
	result := []string{}

	for _, v := range stacks {
		result = append(result, strings.Trim(v.crates[0], "[]"))
	}

	fmt.Println(strings.Join(result, ""))
}

func moveCrate(stacks []Stack, amount int, from int, to int) {
	move := make([]string, len(stacks[from].crates))
	copy(move, stacks[from].crates)
	move = move[:amount]

	stacks[from].crates = stacks[from].crates[amount:]

	stacks[to].crates = append(move, stacks[to].crates...)
}

func splitStacks(input string) []string {
	separated := []string{}

	for i, v := range input {
		if 0 != (i+1)%4 {
			separated = append(separated, string(v))
		} else {
			separated = append(separated, ",")
		}
	}

	splittedStacks := strings.Split(strings.Join(separated, ""), ",")

	return splittedStacks
}
