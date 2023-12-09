package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var totalPartOne = 1
var totalPartTwo = 0
var timeData = []int{}
var distanceData = []int{}

func main() {
	file, err := os.Open("./input")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	timeStrSlice := []string{}
	distanceStrSlice := []string{}

	for scanner.Scan() {
		line := scanner.Text()
		if strings.HasPrefix(line, "Time:") {
			timeStrSlice = strings.Fields(strings.Split(line, ":")[1])
		}
		if strings.HasPrefix(line, "Distance:") {
			distanceStrSlice = strings.Fields(strings.Split(line, ":")[1])
		}
	}

	var timePart2Str, distancePart2Str string

	for _, s := range timeStrSlice {
		n, _ := strconv.Atoi(s)
		timeData = append(timeData, n)
		timePart2Str += s
	}

	for _, s := range distanceStrSlice {
		n, _ := strconv.Atoi(s)
		distanceData = append(distanceData, n)
		distancePart2Str += s
	}

	waysToWin := []int{}

	for i, v := range timeData {
		time := v
		distance := distanceData[i]
		var count int

		for j := 1; j < time; j++ {
			if (time-j)*j > distance {
				count++
			}
		}

		waysToWin = append(waysToWin, count)
	}

	for _, w := range waysToWin {
		totalPartOne *= w
	}

	// Part 2
	timePart2, _ := strconv.Atoi(timePart2Str)
	distancePart2, _ := strconv.Atoi(distancePart2Str)

	for i := 1; i < timePart2; i++ {
		if (timePart2-i)*i > distancePart2 {
			totalPartTwo++
		}
	}

	fmt.Println("Part 1:", totalPartOne)
	fmt.Println("Part 2:", totalPartTwo)
}
