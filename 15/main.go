package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type box struct {
	contents []string
}

var totalPartOne int
var totalPartTwo int
var boxes = make([]box, 256)

func hashStr(s string) int {
	var hashed int
	for _, c := range s {
		hashed = (hashed + int(c)) * 17 % 256
	}
	return hashed
}

func findLabelInBox(b *box, l string) int {
	for i := 0; i < len(b.contents); i++ {
		if strings.HasPrefix(b.contents[i], l) {
			return i
		}
	}
	return -1
}

func main() {
	file, err := os.ReadFile("./input")
	if err != nil {
		panic(err)
	}
	input := strings.Split(string(file[:len(file)-1]), ",")

	for _, s := range input {
		totalPartOne += hashStr(s)

		var op, label string
		var lens int

		if strings.Contains(s, "-") {
			op = "-"
			label = strings.Split(s, "-")[0]
			lens, _ = strconv.Atoi(strings.Split(s, "-")[1])
		} else if strings.Contains(s, "=") {
			op = "="
			label = strings.Split(s, "=")[0]
			lens, _ = strconv.Atoi(strings.Split(s, "=")[1])
		}

		theBox := &boxes[hashStr(label)]
		labelIdx := findLabelInBox(theBox, label)

		if op == "-" && labelIdx != -1 {
			theBox.contents = append(theBox.contents[:labelIdx], theBox.contents[labelIdx+1:]...)
		} else if op == "=" {
			if labelIdx == -1 {
				theBox.contents = append(theBox.contents, label+"-"+strconv.Itoa(lens))
			} else {
				theBox.contents[labelIdx] = label + "-" + strconv.Itoa(lens)
			}
		}
	}

	for i, box := range boxes {
		for j := 0; j < len(box.contents); j++ {
			focalLength, _ := strconv.Atoi(strings.Split(box.contents[j], "-")[1])
			totalPartTwo += (1 + i) * (j + 1) * focalLength
		}
	}

	fmt.Println("Part 1:", totalPartOne)
	fmt.Println("Part 2:", totalPartTwo)
}
