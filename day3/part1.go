package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	// input := []string{
	// 	"00100",
	// 	"11110",
	// 	"10110",
	// 	"10111",
	// 	"10101",
	// 	"01111",
	// 	"00111",
	// 	"11100",
	// 	"10000",
	// 	"11001",
	// 	"00010",
	// 	"01010",
	// }

	input := readInput()

	var gamma string
	var epsilon string

	b := true
	idx := 0

	for {
		ones := 0
		zeroes := 0
		for _, binary := range input {
			if idx+1 > len(binary) {
				b = false
				break
			}
			numeric, err := strconv.Atoi(string(binary[idx]))
			if err != nil {
				panic(err)
			}
			if numeric == 1 {
				ones++
			} else {
				zeroes++
			}
		}
		if b == false {
			break
		}
		if ones > zeroes {
			gamma += "1"
		} else {
			gamma += "0"
		}
		idx++
	}

	for _, num := range gamma {
		if string(num) == "0" {
			epsilon += "1"
		} else {
			epsilon += "0"
		}
	}

	gammaDec, err := strconv.ParseInt(gamma, 2, 64)
	if err != nil {
		panic(err)
	}

	epsilonDec, err := strconv.ParseInt(epsilon, 2, 64)
	if err != nil {
		panic(err)
	}

	// fmt.Printf("Gamma in binary: %s\n", gamma)
	// fmt.Printf("Gamma in decimal: %v\n", gammaDec)
	// fmt.Printf("Epsilon in binary: %s\n", epsilon)
	// fmt.Printf("Epsilon in decimal: %v\n", epsilonDec)
	fmt.Println(gammaDec * epsilonDec)
}

func readInput() []string {
	f, err := os.Open("./input")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	input := []string{}

	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		line := scanner.Text()
		input = append(input, line)
	}
	return input
}
