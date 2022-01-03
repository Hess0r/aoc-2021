package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	input := readInput()

	idx := 0

	o2Reduce := input
	co2Reduce := input

	for len(o2Reduce) > 1 {
		o2Reduce = filterAtPosition(o2Reduce, idx, true)
		idx++
	}

	idx = 0

	for len(co2Reduce) > 1 {
		co2Reduce = filterAtPosition(co2Reduce, idx, false)
		idx++
	}

	o2Dec, err := strconv.ParseInt(o2Reduce[0], 2, 64)
	if err != nil {
		panic(err)
	}

	co2Dec, err := strconv.ParseInt(co2Reduce[0], 2, 64)
	if err != nil {
		panic(err)
	}

	fmt.Println(o2Dec * co2Dec)
}

func filterAtPosition(i []string, index int, mcb bool) []string {
	var ones []int
	var zeroes []int
	var out []string

	for idx, binary := range i {
		if string(binary[index]) == "1" {
			ones = append(ones, idx)
		} else {
			zeroes = append(zeroes, idx)
		}
	}

	if mcb == true {
		if len(ones) >= len(zeroes) {
			for _, idx := range ones {
				out = append(out, i[idx])
			}
		} else {
			for _, idx := range zeroes {
				out = append(out, i[idx])
			}
		}
	} else {
		if len(zeroes) <= len(ones) {
			for _, idx := range zeroes {
				out = append(out, i[idx])
			}
		} else {
			for _, idx := range ones {
				out = append(out, i[idx])
			}
		}
	}

	return out
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
