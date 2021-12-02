package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/pgcosta/aoc2021/2/submarine"
)

func check(err error) {
	if err != nil {
		panic("Kaboom!")
	}
}

func readInput() []string {
	file, err := os.Open("input.txt")
	check(err)
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var lines []string

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	return lines
}

func main() {
	commands := readInput()
	sub := submarine.Submarine{}

	for _, command := range commands {
		sub.Move(command)
	}

	fmt.Printf("Forward: %d\n", sub.Forward)
	fmt.Printf("Depth: %d\n", sub.Depth)
	fmt.Printf("Total with dive calculation: %d\n", sub.Total())
}
