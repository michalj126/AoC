package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type CPU struct {
	register int
	cycle    int
}

func NewCPU() *CPU {
	return &CPU{
		register: 1,
		cycle:    0,
	}
}

func (o *CPU) execInstruction(input string) {
	instruction := strings.Split(input, " ")

	if instruction[0] == "noop" {
		o.renderPixel()
		o.cycle++
	}

	if instruction[0] == "addx" {
		for i := 0; i < 2; i++ {
			o.renderPixel()
			o.cycle++
		}

		valueInt, err := strconv.Atoi(instruction[1])
		check(err)

		o.register += valueInt
	}
}

func (o *CPU) renderPixel() {
	if o.cycle%40 == 0 {
		fmt.Println()
	}

	if o.cycle%40 >= o.register-1 && o.cycle%40 <= o.register+1 {
		fmt.Print("#")
	} else {
		fmt.Print(".")
	}
}

var sumSignalStrength = 0

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
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
