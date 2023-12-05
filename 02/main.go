package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var sumPartOne int
var sumPartTwo int

var gameRules = map[string]int{
	"red":   12,
	"green": 13,
	"blue":  14,
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

		gameID, err := strconv.Atoi(strings.Split(strings.Split(line, ":")[0], " ")[1])
		if err != nil {
			fmt.Println("Error parsing gameID")
		}

		gameInfo := strings.Split(strings.Split(line, ":")[1], ";")

		// Part 1
		possible := true

		// Part 2
		minReq := map[string]int{
			"red":   0,
			"green": 0,
			"blue":  0,
		}

		for _, round := range gameInfo {
			sets := strings.Split(round, ",")
			for _, set := range sets {
				num, err := strconv.Atoi(strings.Split(strings.Trim(set, " "), " ")[0])
				if err != nil {
					fmt.Println("Error parsing number of cubes")
				}
				color := strings.Split(strings.Trim(set, " "), " ")[1]

				// Part 2
				if num > minReq[color] {
					minReq[color] = num
				}

				// Part 1
				if num > gameRules[color] {
					possible = false
				}
			}
		}

		// Part 1
		if possible {
			sumPartOne += gameID
		}

		// Part 2
		sumPartTwo += minReq["red"] * minReq["green"] * minReq["blue"]
	}

	fmt.Println("Part 1:", sumPartOne)
	fmt.Println("Part 2:", sumPartTwo)
}
