package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	data, err := os.ReadFile("input.txt")
	check(err)

	numbers := strings.Fields(string(data))

	previous_number, err := strconv.Atoi(numbers[0])
	check(err)

	number_of_increasings := 0
	for _, number := range numbers {
		int_number, _ := strconv.Atoi(number)
		if int_number > previous_number {
			number_of_increasings = number_of_increasings + 1
		}
		previous_number = int_number
	}
	fmt.Println(number_of_increasings)
}
