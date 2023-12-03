package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
	"unicode"
)

var sumPartOne int
var sumPartTwo int

func partOne(line string) {
	var calVal string

	for i := 0; i < len(line); i++ {
		if unicode.IsDigit(rune(line[i])) {
			calVal += string(line[i])
			break
		}
	}

	for i := len(line) - 1; i >= 0; i-- {
		if unicode.IsDigit(rune(line[i])) {
			calVal += string(line[i])
			break
		}
	}

	lineCalVal, _ := strconv.Atoi(calVal)
	sumPartOne += lineCalVal
}

func partTwo(line string) {
	digits := []string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9"}
	words := []string{"zero", "one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}

	minIdx := len(line)
	maxIdx := 0

	var firstDigit, lastDigit string

	for _, digit := range digits {
		firstDigitIndex := strings.Index(line, digit)
		if firstDigitIndex != -1 && firstDigitIndex < minIdx {
			minIdx = firstDigitIndex
			firstDigit = string(line[minIdx])
		}
		lastDigitIndex := strings.LastIndex(line, digit)
		if lastDigitIndex != -1 && lastDigitIndex >= maxIdx {
			maxIdx = lastDigitIndex
			lastDigit = string(line[maxIdx])
		}
	}

	for _, word := range words {
		firstDigitIndex := strings.Index(line, word)
		if firstDigitIndex != -1 && firstDigitIndex < minIdx {
			minIdx = firstDigitIndex
			firstDigit = string(line[minIdx : minIdx+len(word)])
		}
		lastDigitIndex := strings.LastIndex(line, word)
		if lastDigitIndex != -1 && lastDigitIndex > maxIdx {
			maxIdx = lastDigitIndex
			lastDigit = string(line[maxIdx : maxIdx+len(word)])
		}
	}

	var calVal string

	if slices.Contains(words, firstDigit) {
		calVal += fmt.Sprint(slices.Index(words, firstDigit))
	} else {
		calVal += firstDigit
	}

	if slices.Contains(words, lastDigit) {
		calVal += fmt.Sprint(slices.Index(words, lastDigit))
	} else {
		calVal += lastDigit
	}

	lineCalVal, _ := strconv.Atoi(calVal)
	sumPartTwo += lineCalVal
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
		partOne(line)
		partTwo(line)
	}

	fmt.Println("Part 1:", sumPartOne)
	fmt.Println("Part 2:", sumPartTwo)
}
