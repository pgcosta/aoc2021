package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"strconv"
)

func checkError(e error) {
	if e != nil {
		panic(e)
	}
}

type binaryReport struct {
	zero int
	one  int
}

func calculateLargerByte(report [12]binaryReport, numberOfLines int) string {
	var buffer bytes.Buffer

	for _, reportByte := range report {
		if reportByte.one > reportByte.zero {
			buffer.WriteString("1")
		} else {
			buffer.WriteString("0")
		}
	}

	return buffer.String()
}

func flipBits(input string) string {
	var buffer bytes.Buffer

	for _, c := range input {
		if c == '1' {
			buffer.WriteString("0")
		} else {
			buffer.WriteString("1")
		}
	}

	return buffer.String()
}

func main() {
	file, err := os.Open("input.txt")
	checkError(err)
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var energyReport [12]binaryReport
	var numberOfLines int

	for scanner.Scan() {
		for i, bit := range scanner.Text() {
			if int(bit-48) == 0 {
				energyReport[i].zero += 1
			} else {
				energyReport[i].one += 1
			}
			numberOfLines++
		}
	}

	fmt.Printf("Energy report: %+v \n", energyReport)

	gamma := calculateLargerByte(energyReport, numberOfLines)

	gammaDec, _ := strconv.ParseInt(gamma, 2, 32)
	epsilonDec, _ := strconv.ParseInt(flipBits(gamma), 2, 32)

	fmt.Printf("gamma value: %d\n", gammaDec)
	fmt.Printf("epsilon byte: %d\n", epsilonDec)
	fmt.Printf("Total consumption: %d\n", gammaDec*epsilonDec)
}
