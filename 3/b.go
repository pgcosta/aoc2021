package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func readFile(filename string) []string {
	fileContents, _ := os.ReadFile("input.txt")
	lines := strings.Split(string(fileContents), "\n")
	// remove last string since it's a \0
	return lines[0:(len(lines) - 1)]
}

func mostCommon(column int, list []string, flip bool) string {
	if len(list) == 1 {
		return list[0]
	}

	var ones []string
	var zeros []string

	for _, word := range list {
		if string(word[column]) == "1" {
			ones = append(ones, word)
		} else {
			zeros = append(zeros, word)
		}
	}

	if len(ones) >= len(zeros) == flip {
		return mostCommon(column+1, ones, flip)
	} else {
		return mostCommon(column+1, zeros, flip)
	}
}

func main() {
	o2ReportList := readFile("input.txt")
	var co2ReportList = make([]string, len(o2ReportList))
	copy(co2ReportList, o2ReportList)

	o2GeneratorRating, _ := strconv.ParseInt(mostCommon(0, o2ReportList, true), 2, 32)
	co2ScrubberRating, _ := strconv.ParseInt(mostCommon(0, co2ReportList, false), 2, 32)

	fmt.Printf("Oxygen generator rating: %d\n", o2GeneratorRating)
	fmt.Printf("CO2 Scrubber rating: %d\n", co2ScrubberRating)
	fmt.Printf("Total: %d\n", o2GeneratorRating*co2ScrubberRating)
}
