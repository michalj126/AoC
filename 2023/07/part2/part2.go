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
		return CompareHand(hands[i].Cards, hands[j].Cards, hands)
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

func CompareHand(a string, b string, hands []Hand) (result bool) {
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
	cardPower["A"] = 13
	cardPower["K"] = 12
	cardPower["Q"] = 11
	cardPower["T"] = 10
	cardPower["9"] = 9
	cardPower["8"] = 8
	cardPower["7"] = 7
	cardPower["6"] = 6
	cardPower["5"] = 5
	cardPower["4"] = 4
	cardPower["3"] = 3
	cardPower["2"] = 2
	cardPower["J"] = 1

	if a == "JJJJJ" {
		result = true
		return
	}

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

func IsFive(cards string) (isFive bool) {
	cards = strings.ReplaceAll(cards, "J", "")

	if len(cards) > 0 {
		cards = strings.ReplaceAll(cards, string(cards[0]), "")
	}

	if cards == "" {
		isFive = true
	}

	return
}

func IsFour(cards string) bool {
	joker := strings.Count(cards, "J")

	if joker == 3 {
		return true
	}

	cards = strings.ReplaceAll(cards, "J", "")

	ch1 := strings.Count(cards, string(cards[0]))
	ch2 := 0

	if len(cards) > 1 {
		ch2 = strings.Count(cards, string(cards[1]))
	}

	return ((ch1 == 2 || ch2 == 2) && joker == 2) || ((ch1 == 3 || ch2 == 3) && joker == 1) || (ch1 == 4 || ch2 == 4)
}

func IsFull(cards string) (isFull bool) {
	if strings.Count(cards, "J") == 1 {
		cards = strings.ReplaceAll(cards, "J", "")
	}

	for i := 0; i < 2; i++ {
		cards = strings.ReplaceAll(cards, string(cards[0]), "")
	}

	if cards == "" {
		isFull = true
	}

	return
}

func IsThree(cards string) (isThree bool) {
	joker := strings.Count(cards, "J")

	if joker == 2 {
		isThree = true
		return
	}

	cards = strings.ReplaceAll(cards, "J", "")

	ch1 := strings.Count(cards, string(cards[0])) + joker
	ch2 := strings.Count(cards, string(cards[1])) + joker
	ch3 := strings.Count(cards, string(cards[2])) + joker

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
	if strings.Count(cards, "J") == 1 {
		isOnePair = true
		return
	}

	length := len(cards)

	for i := 0; i < length-1; i++ {
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
