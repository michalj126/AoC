package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Monkey struct {
	items               []int
	operation           Operation
	test                Test
	totalInspectedItems int
}

type Operation struct {
	sign string
	x    int
	y    int
}

type Test struct {
	divisibleBy int
	true        int
	false       int
}

type Round struct {
	counter int
	monkeys []Monkey
}

var round = Round{
	counter: 0,
	monkeys: []Monkey{},
}

var relief = 0

func main() {
	path := os.Args[1]

	f, err := os.Open(path)
	check(err)

	defer f.Close()

	parse(f)

	relief = findReliefValue()

	for i := 0; i < 10000; i++ {
		for iM, monkey := range round.monkeys {
			for _, item := range round.monkeys[iM].items {
				round.monkeys[iM].totalInspectedItems++

				if monkey.operation.sign == "+" {
					if monkey.operation.x == 0 && monkey.operation.y != 0 {
						item = item + monkey.operation.y
					} else if monkey.operation.x == 0 && monkey.operation.y == 0 {
						item = item + item
					}
				} else {
					if monkey.operation.x == 0 && monkey.operation.y != 0 {
						item = item * monkey.operation.y
					} else if monkey.operation.x == 0 && monkey.operation.y == 0 {
						item = item * item
					}
				}

				item %= relief

				throwTo := 0

				if item%monkey.test.divisibleBy == 0 {
					throwTo = monkey.test.true
				} else {
					throwTo = monkey.test.false
				}

				round.monkeys[throwTo].items = append(round.monkeys[throwTo].items, item)
				if len(round.monkeys[iM].items) == 1 {
					round.monkeys[iM].items = []int{}
				} else {
					round.monkeys[iM].items = round.monkeys[iM].items[1:]
				}
			}
		}

	}

	fmt.Println(calcLevelOfMonkeyBusiness())
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func calcLevelOfMonkeyBusiness() (levelOfMonkeyBusiness int) {
	twoMostActive := [2]int{0, 0}

	for _, monkey := range round.monkeys {
		if monkey.totalInspectedItems > twoMostActive[0] {
			twoMostActive[1] = twoMostActive[0]
			twoMostActive[0] = monkey.totalInspectedItems
		} else if monkey.totalInspectedItems > twoMostActive[1] {
			twoMostActive[1] = monkey.totalInspectedItems
		}
	}

	levelOfMonkeyBusiness = twoMostActive[0] * twoMostActive[1]

	return
}

func findReliefValue() (relief int) {
	relief = round.monkeys[0].test.divisibleBy

	for _, monkey := range round.monkeys[1:] {
		relief *= monkey.test.divisibleBy
	}

	return
}

func parse(f *os.File) {
	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		line := scanner.Text()

		if len(line) > 0 {
			line = strings.TrimSpace(line)

			if line[0:1] == "M" {
				round.monkeys = append(round.monkeys, Monkey{})
			}

			if line[0:1] == "S" {
				items := strings.Split(strings.Split(line, ": ")[1], ", ")

				for _, v := range items {
					vInt, err := strconv.Atoi(v)
					check(err)

					round.monkeys[len(round.monkeys)-1].items = append(round.monkeys[len(round.monkeys)-1].items, int(vInt))
				}
			}

			if line[0:1] == "O" {
				operation := strings.Split(strings.Split(line, " = ")[1], " ")

				if operation[0] != "old" {
					xInt, err := strconv.Atoi(operation[0])
					check(err)

					round.monkeys[len(round.monkeys)-1].operation.x = int(xInt)
				}

				if operation[2] != "old" {
					yInt, err := strconv.Atoi(operation[2])
					check(err)

					round.monkeys[len(round.monkeys)-1].operation.y = int(yInt)
				}

				round.monkeys[len(round.monkeys)-1].operation.sign = operation[1]
			}

			if line[0:1] == "T" {
				splitted := strings.Split(line, " ")
				divisibleBy, err := strconv.Atoi(splitted[len(splitted)-1])
				check(err)

				round.monkeys[len(round.monkeys)-1].test.divisibleBy = int(divisibleBy)
			}

			if line[0:1] == "I" {
				splittedLine := strings.Split(line, ": ")

				splitted := strings.Split(splittedLine[1], " ")
				throwTo, err := strconv.Atoi(splitted[len(splitted)-1])
				check(err)

				if strings.Split(splittedLine[0], " ")[1] == "true" {
					round.monkeys[len(round.monkeys)-1].test.true = throwTo
				} else {
					round.monkeys[len(round.monkeys)-1].test.false = throwTo
				}
			}
		}
	}
}
