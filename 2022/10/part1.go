package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type CPU struct {
	register          int
	cycle             int
	sumSignalStrength int
}

func NewCPU() *CPU {
	return &CPU{
		register:          1,
		cycle:             0,
		sumSignalStrength: 0,
	}
}

func (o *CPU) execInstruction(input string) {
	instruction := strings.Split(input, " ")

	if instruction[0] == "noop" {
		o.cycle++
		o.calcSumSignalStrength()
	}

	if instruction[0] == "addx" {
		for i := 0; i < 2; i++ {
			o.cycle++
			o.calcSumSignalStrength()
		}

		valueInt, err := strconv.Atoi(instruction[1])
		check(err)

		o.register += valueInt
	}
}

func (o *CPU) calcSumSignalStrength() {
	if (o.cycle == 20 || (o.cycle-20)%40 == 0) && o.cycle <= 220 {
		o.sumSignalStrength += (o.cycle * o.register)
	} else if (o.cycle-20)%40 == 0 {
	}
}

func main() {
	path := os.Args[1]

	f, err := os.Open(path)
	check(err)

	defer f.Close()

	scanner := bufio.NewScanner(f)

	cpu := NewCPU()

	for scanner.Scan() {
		line := scanner.Text()

		cpu.execInstruction(line)
	}

	fmt.Println(cpu.sumSignalStrength)
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
