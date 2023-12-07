package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"regexp"
	"slices"
	"strconv"
	"strings"
)

var totalPartOne int
var totalPartTwo int
var cardMap = make(map[int]int)

func removeEmptyElements(s []string) []string {
	for i, v := range s {
		if v == "" {
			s = append(s[:i], s[i+1:]...)
		}
	}
	return s
}

func main() {
	file, err := os.Open("./input")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		line := scanner.Text()

		// Part 2
		cardNum := strings.Split(line, ":")[0]
		var cardIdStr string
		for _, s := range strings.Split(cardNum, "") {
			if matched, _ := regexp.MatchString("[0-9]", s); matched {
				cardIdStr += s
			}
		}
		cardId, err := strconv.Atoi(cardIdStr)
		if err != nil {
			fmt.Println("Error parsing card ID", cardNum)
		}
		cardMap[cardId]++
		////

		winningNumbers := strings.Trim(strings.Split(strings.Split(line, ":")[1], "|")[0], " ")
		hadNumbers := strings.Trim(strings.Split(line, "|")[1], " ")

		winningNumbersSlice := strings.Split(winningNumbers, " ")
		winningNumbersSlice = removeEmptyElements(winningNumbersSlice)

		hadNumbersSlice := strings.Split(hadNumbers, " ")
		hadNumbersSlice = removeEmptyElements(hadNumbersSlice)

		var count int
		for _, number := range hadNumbersSlice {
			if slices.Contains(winningNumbersSlice, number) {
				count++
			}
		}

		totalPartOne += int(math.Pow(2, float64(count-1)))

		// Part 2
		totalPartTwo += cardMap[cardId]

		if count > 0 {
			for i := cardId + 1; i <= cardId+count; i++ {
				cardMap[i] += cardMap[cardId]
			}
		}
	}

	fmt.Println("Part 1:", totalPartOne)
	fmt.Println("Part 2:", totalPartTwo)
}
