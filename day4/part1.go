package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type mark struct {
	row int
	col int
}

type board struct {
	matrix [][]string
	marks  []mark
}

func main() {
	input := readInput()

	var boards []board

	drawn := input[0]
	fmt.Println(drawn)

	input = input[2:]

	idx := 0

	for idx+1 < len(input) {
		matrixLines := input[idx : idx+5]
		var matrix [][]string
		for _, line := range matrixLines {
			// splitLine := strings.Split(line, " ")
			var splitLine []string
			i := 0
			for i < len(line) {
				splitLine = append(splitLine, strings.TrimSpace(string(line[i:i+2])))
				i += 3
			}
			matrix = append(matrix, splitLine)
		}
		boards = append(boards, board{matrix: matrix})
		idx += 6
	}

	fmt.Println(boards[0].matrix[0][0])
	fmt.Println(boards[0].matrix[4][4])
}

func readInput() []string {
	f, err := os.Open("./testinput")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	input := []string{}

	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		line := scanner.Text()
		fmt.Println(line)
		input = append(input, line)
	}
	return input
}
