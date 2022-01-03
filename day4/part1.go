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

	var boards []*board

	drawn := strings.Split(input[0], ",")

	input = input[2:]

	idx := 0

	for idx+1 < len(input) {
		matrixLines := input[idx : idx+5]
		var matrix [][]string
		for _, line := range matrixLines {
			var splitLine []string
			i := 0
			for i < len(line) {
				splitLine = append(splitLine, strings.TrimSpace(string(line[i:i+2])))
				i += 3
			}
			matrix = append(matrix, splitLine)
		}
		boards = append(boards, &board{matrix: matrix})
		idx += 6
	}

	var winnerBoard *board

	for _, number := range drawn {
		fmt.Println(number)
		for _, board := range boards {
			won := board.markBoard(number)
			if won == true {
				winnerBoard = board
				break
			}
		}
		if winnerBoard != nil {
			break
		}
	}

	fmt.Println(winnerBoard)
}

func (b *board) markBoard(n string) bool {
	for rowIdx, row := range b.matrix {
		for colIdx, num := range row {
			if num == n {
				b.marks = append(b.marks, mark{
					row: rowIdx,
					col: colIdx,
				})
			}
		}
	}

	return b.checkMarks()
}

func (b *board) checkMarks() bool {
	for row := 0; row < 5; row++ {
		marksInCol := 0
		for col := 0; col < 5; col++ {
			for _, mark := range b.marks {
				if mark.col == col && mark.row == row {
					marksInCol++
				}
			}
		}
		if marksInCol == 5 {
			return true
		}
	}
	return false
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
		input = append(input, line)
	}
	return input
}
