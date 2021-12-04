package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func Part2(boards [][][]string, randoms []string) int {
	seenBoard := []int{}
	boolBoards := getBooleanBoard(boards)
	for _, num := range randoms {
		for b := 0; b < len(boards); b++ {
			board := boards[b]

			recent := 1
			for j := 0; j < len(board); j++ {
				for k := 0; k < len(board[j]); k++ {
					if board[j][k] == num {
						boolBoards[b][j][k] = true
						recent, _ = strconv.Atoi(board[j][k])
					}
				}

				if allTrue(boolBoards[b][j]) || allTrueRow(boolBoards[b], j) {
					if !arrIncludes(seenBoard, b) {
						seenBoard = append(seenBoard, b)
					}
					if len(seenBoard) == len(boards) {
						return sum(board, boolBoards[b]) * recent
					}
				}
			}

		}
	}
	return 0

}
func Part1(boards [][][]string, randoms []string) int {
	boolBoards := getBooleanBoard(boards)
	for _, num := range randoms {
		for b := 0; b < len(boards); b++ {
			board := boards[b]

			recent := 1
			for j := 0; j < len(board); j++ {
				for k := 0; k < len(board[j]); k++ {
					if board[j][k] == num {
						boolBoards[b][j][k] = true
						recent, _ = strconv.Atoi(board[j][k])
					}
				}

				if allTrue(boolBoards[b][j]) {
					return sum(board, boolBoards[b]) * recent
				}
			}

		}
	}
	return 0

}

func getBooleanBoard(boards [][][]string) [][][]bool {
	boolBoards := [][][]bool{}
	for b := 0; b < len(boards); b++ {
		board := boards[b]
		boolBoard := [][]bool{}
		for _, line := range board {

			boolLine := []bool{}
			for range line {
				boolLine = append(boolLine, false)
			}
			boolBoard = append(boolBoard, boolLine)
		}
		boolBoards = append(boolBoards, boolBoard)
	}
	return boolBoards
}

func allTrue(line []bool) bool {
	for _, b := range line {
		if !b {
			return false
		}
	}
	return true
}

func allTrueRow(board [][]bool, row int) bool {
	for _, b := range board {
		if !b[row] {
			return false
		}
	}
	return true
}

func sum(board [][]string, boolBoard [][]bool) int {
	sum := 0
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[i]); j++ {
			if !boolBoard[i][j] {
				num, _ := strconv.Atoi(board[i][j])
				sum += num
			}
		}
	}
	return sum
}

func arrIncludes(arr []int, str int) bool {
	for _, a := range arr {
		if a == str {
			return true
		}
	}
	return false
}

func main() {
	text, err := os.ReadFile("./../input/day4.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	input := strings.Split(string(text), "\n")

	randoms := strings.Split(input[0], ",")
	input = input[2:]
	boards := [][][]string{}
	board := [][]string{}
	// boardNum := 0
	for _, line := range input {
		if line == "" {
			boards = append(boards, board)
			board = [][]string{}
			continue
		}
		line = strings.TrimSpace(regexp.MustCompile(`\s\s+`).ReplaceAllString(line, " "))
		board = append(board, strings.Split(line, " "))
	}
	boards = append(boards, board)

	// for i := 0; i < len(boards); i++ {
	// 	// fmt.Println(boards[i])
	// 	fmt.Println(day4(boards[i], randoms))
	// }
	// fmt.Println(board, randoms)

	fmt.Println(Part1(boards, randoms))
	fmt.Println(Part2(boards, randoms))

}
