package main

import (
	"bufio"
	"fmt"
	"os"
	"runtime"
	"sort"
	"strconv"
	"strings"
	"time"
)

type Hand struct {
	Cards string
	Bid   int
}

func main() {
	start := time.Now()

	path := os.Args[1]

	f, err := os.Open(path)
	Check(err)

	defer f.Close()

	scanner := bufio.NewScanner(f)

	hands := []Hand{}

	for scanner.Scan() {
		line := scanner.Text()

		splited := strings.Split(line, " ")
		bid, err := strconv.Atoi(splited[1])
		Check(err)

		hands = append(hands, Hand{splited[0], bid})
	}

	sort.Slice(hands, func(i, j int) bool {
		return CompareHand(hands[i].Cards, hands[j].Cards)
	})

	totalWinnings := 0

	for i, h := range hands {
		totalWinnings += h.Bid * (i + 1)
	}

	fmt.Println("Total winnings: ", totalWinnings)

	duration := time.Since(start)

	fmt.Println("Execution time: ", duration)
	PrintMemUsage()
}

func CompareHand(a string, b string) (result bool) {
	aP, bP := GetPower(a), GetPower(b)

	if aP == bP {
		result = CompareSamePower(a, b)
		return
	}

	result = aP < bP

	return
}

func CompareSamePower(a, b string) (result bool) {
	cardPower := make(map[string]int)
	cardPower["A"] = 14
	cardPower["K"] = 13
	cardPower["Q"] = 12
	cardPower["J"] = 11
	cardPower["T"] = 10
	cardPower["9"] = 9
	cardPower["8"] = 8
	cardPower["7"] = 7
	cardPower["6"] = 6
	cardPower["5"] = 5
	cardPower["4"] = 4
	cardPower["3"] = 3
	cardPower["2"] = 2

	for i := range a {
		if a[i] != b[i] {
			return cardPower[string(a[i])] < cardPower[string(b[i])]
		}
	}

	return
}

func GetPower(cards string) (power int) {
	switch {
	case IsFive(cards):
		power = 6
	case IsFour(cards):
		power = 5
	case IsFull(cards):
		power = 4
	case IsThree(cards):
		power = 3
	case IsTwoPair(cards):
		power = 2
	case IsOnePair(cards):
		power = 1
	}

	return
}

func IsFive(cards string) bool {
	return strings.Count(cards, string(cards[0])) == 5
}

func IsFour(cards string) bool {
	return strings.Count(cards, string(cards[0])) == 4 || strings.Count(cards, string(cards[1])) == 4
}

func IsFull(cards string) (isFull bool) {
	for i := 0; i < 2; i++ {
		cards = strings.ReplaceAll(cards, string(cards[0]), "")
	}

	if cards == "" {
		isFull = true
	}

	return
}

func IsThree(cards string) (isThree bool) {
	ch1 := strings.Count(cards, string(cards[0]))
	ch2 := strings.Count(cards, string(cards[1]))
	ch3 := strings.Count(cards, string(cards[2]))

	if ch1 == 3 || ch2 == 3 || ch3 == 3 {
		isThree = true
	}

	return
}

func IsTwoPair(cards string) (isTwoPair bool) {
	for i := 0; i < 2; i++ {
		cards = strings.ReplaceAll(cards, string(cards[0]), "")
	}

	if len(cards) == 1 {
		isTwoPair = true
		return
	}

	cards = strings.ReplaceAll(cards, string(cards[0]), "")

	if len(cards) == 0 {
		isTwoPair = true
	}

	return
}

func IsOnePair(cards string) (isOnePair bool) {
	for i := 0; i < 4; i++ {
		cards = strings.ReplaceAll(cards, string(cards[0]), "")
	}

	if len(cards) == 0 {
		isOnePair = true
	}

	return
}

func ConvertToInt(input []string) []int {
	result := make([]int, len(input))

	for i, v := range input {
		vi, err := strconv.Atoi(v)
		Check(err)

		result[i] = vi
	}

	return result
}

func Check(e error) {
	if e != nil {
		panic(e)
	}
}

func PrintMemUsage() {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	fmt.Printf("Alloc = %v MiB", m.Alloc/1024/1024)
	fmt.Printf(" | %v KiB", m.Alloc/1024)
	fmt.Printf("\tTotalAlloc = %v MiB", m.TotalAlloc/1024/1024)
	fmt.Printf(" | %v KiB", m.TotalAlloc/1024)
	fmt.Printf("\tSys = %v MiB", m.Sys/1024/1024)
	fmt.Printf(" | Sys = %v KiB", m.Sys/1024)
	fmt.Printf("\tNumGC = %v\n", m.NumGC)
}
