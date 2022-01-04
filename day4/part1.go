package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type mark struct {
	row int
	col int
}

type board struct {
	matrix      [][]string
	marks       []mark
	winnerMarks []mark
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
	var lastNumber string

	for _, number := range drawn {
		for _, board := range boards {
			won := board.markBoard(number)
			if won == true {
				winnerBoard = board
				lastNumber = number
				break
			}
		}
		if winnerBoard != nil {
			break
		}
	}

	score := winnerBoard.calculateScore(lastNumber)

	fmt.Println(score)
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
	if len(b.marks) < 5 {
		return false
	}
	for col := 0; col < 5; col++ {
		var sameCol []mark
		var sameRow []mark
		for _, mark := range b.marks {
			if mark.col == col {
				sameCol = append(sameCol, mark)
			}
			if mark.row == col {
				sameRow = append(sameRow, mark)
			}
		}
		if len(sameCol) == 5 {
			b.winnerMarks = sameCol
			return true
		}
		if len(sameRow) == 5 {
			b.winnerMarks = sameRow
			return true
		}
	}
	return false
}

func (b *board) calculateScore(ln string) int {
	lnI, _ := strconv.Atoi(ln)
	sum := 0

	for row := 0; row < 5; row++ {
		for col := 0; col < 5; col++ {
			unmarked := true
			for _, mark := range b.marks {
				if mark.col == col && mark.row == row {
					unmarked = false
				}
			}
			if unmarked {
				val, _ := strconv.Atoi(b.matrix[row][col])
				sum += val
			}
		}
	}

	return sum * lnI
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
