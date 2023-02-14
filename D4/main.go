package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func getFileLines(filename string) []string {
	data, _ := os.ReadFile(filename)
	dataArray := strings.Split(string(data), "\n")
	return dataArray
}

func strToInt(str string) int {
	str = strings.TrimSpace(str)
	integer, _ := strconv.Atoi((str))
	return integer
}

func testWin(board [][]bool) bool {
	for row := 0; row < len(board); row++ {
		hits := 0
		for col := 0; col < len(board[0]); col++ {
			if board[row][col] {
				hits += 1
			}
		}
		if hits == 5 {
			return true
		}
	}
	for col := 0; col < len(board[0]); col++ {
		hits := 0
		for row := 0; row > len(board); row++ {
			if board[row][col] {
				hits += 1
			}
		}
		if hits == 5 {
			return true
		}
	}
	return false
}

func playNum(board [][]int, boolBoard [][]bool, space int) [][]bool {
	for row := 0; row < len(board); row++ {
		for col := 0; col < len(board[0]); col++ {
			if board[row][col] == space {
				boolBoard[row][col] = true
			}
		}
	}
	return boolBoard
}

func grabBoards(lines []string) [][][]int {
	var boards [][][]int
	for iter := 1; iter < len(lines); iter++ {
		if lines[iter] == "" {
			boards = append(boards, [][]int{})
			continue
		}
		boards[len(boards)-1] = append(boards[len(boards)-1], []int{})
		line := strings.Split(lines[iter], " ")
		for iter2 := 0; iter2 < len(line); iter2++ {
			if line[iter2] == "" {
				continue
			}
			boards[len(boards)-1][len(boards[len(boards)-1])-1] = append(boards[len(boards)-1][len(boards[len(boards)-1])-1], strToInt(line[iter2]))
		}
	}
	return boards
}

func convertInstructions(rawInstructions string) []int {
	var instructions []int
	translation := strings.Split(rawInstructions, ",")
	for iter := 0; iter < len(translation); iter++ {
		instructions = append(instructions, strToInt(translation[iter]))
	}
	return instructions
}

func genBoolBoard() [][]bool {
	return [][]bool{
		{false, false, false, false, false},
		{false, false, false, false, false},
		{false, false, false, false, false},
		{false, false, false, false, false},
		{false, false, false, false, false}}
}

func winner(boards [][][]int, instructions []int) [][]int {
	var boolBoards [][][]bool
	for iter := 0; iter < len(boards); iter++ {
		boolBoards = append(boolBoards, genBoolBoard())
	}
	for iter := 0; iter < len(instructions); iter++ {
		for b := 0; b < len(boards); b++ {
			boolBoards[b] = playNum(boards[b], boolBoards[b], instructions[iter])
			if testWin(boolBoards[b]) {
				return boards[b]
			}
		}
	}
	return nil
}

func winnerBool(boards [][][]int, instructions []int) [][]bool {
	var boolBoards [][][]bool
	for iter := 0; iter < len(boards); iter++ {
		boolBoards = append(boolBoards, genBoolBoard())
	}
	for iter := 0; iter < len(instructions); iter++ {
		for b := 0; b < len(boards); b++ {
			boolBoards[b] = playNum(boards[b], boolBoards[b], instructions[iter])
			if testWin(boolBoards[b]) {
				return boolBoards[b]
			}
		}
	}
	return nil
}

func BoolArrayEquals(a [][]bool, b [][]bool) bool {
	if len(a) != len(b) {
		return false
	}
	for r := 0; r < len(a); r++ {
		if len(a[r]) != len(b[r]) {
			return false
		}
		for c := 0; c < len(a[0]); c++ {
			if a[r][c] != b[r][c] {
				return false
			}
		}
	}
	return true
}

func assignBoolMatrix(newMatrix [][]bool) [][]bool {
	// I think this is meant for abstraction?
	var matrix [][]bool
	for r, r_d := range newMatrix {
		matrix = append(matrix, []bool{})
		for _, d := range r_d {
			matrix[r] = append(matrix[r], d)
		}
	}
	return matrix
}

func testIfAnyWin(boards [][][]bool) bool {
	for _, d := range boards {
		if !testWin(d) {
			return false
		}
	}
	return true
}

func getPlayedMoves(boards [][][]int, instructions []int) []int {
	var moves []int
	var boolBoards [][][]bool
	for iter := 0; iter < len(boards); iter++ {
		boolBoards = append(boolBoards, genBoolBoard())
	}
	for iter := 0; iter < len(instructions); iter++ {
		moves = append(moves, instructions[iter])
		for b := 0; b < len(boards); b++ {
			boolBoards[b] = playNum(boards[b], boolBoards[b], instructions[iter])
		}
		if testIfAnyWin(boolBoards) {
			return moves
		}
	}
	return nil
}

func winningMoves(board [][]int, instructions []int) []int {
	// Iterate through the instructions. If the move applied to the board, add it.
	var moves []int
	boolBoard := genBoolBoard()
	for iter := 0; iter < len(instructions); iter++ {
		if testWin(boolBoard) {
			break
		}
		var oldBoard [][]bool = assignBoolMatrix(boolBoard)       // clone of board
		boolBoard = playNum(board, boolBoard, instructions[iter]) // play move
		if BoolArrayEquals(boolBoard, oldBoard) == false {
			moves = append(moves, instructions[iter])
		}
	}
	return moves
}

func markedMoves(board [][]int, instructions []int) []int {
	// Iterate through the instructions. If the move applied to the board, add it.
	var moves []int
	boolBoard := genBoolBoard()
	for iter := 0; iter < len(instructions); iter++ {
		var oldBoard [][]bool = assignBoolMatrix(boolBoard)       // clone of board
		boolBoard = playNum(board, boolBoard, instructions[iter]) // play move
		if !BoolArrayEquals(boolBoard, oldBoard) {
			moves = append(moves, instructions[iter])
		}
	}
	return moves
}

func unmarkedMoves(board [][]int, boolBoard [][]bool) []int {
	var moves []int
	for row, row_dat := range boolBoard {
		for col, col_dat := range row_dat {
			if col_dat == false {
				moves = append(moves, board[row][col])
			}
		}
	}
	return moves
}

func main() {
	lines := getFileLines("input")
	// Convert the first line of moves into an int array
	instructions := convertInstructions(lines[0])
	// Get each game board and convert them to two dimensional arrays
	// Then stuff each of those arrays into another array
	boards := grabBoards(lines)

	// Finds how many moves are played until a board wins
	instructions = getPlayedMoves(boards, instructions)
	// Finds the first winner from the condensed moves
	win := winner(boards, instructions)
	winBool := winnerBool(boards, instructions)
	fmt.Println(win)
	for _, line := range winBool {
		fmt.Println(line)
	}
	// Shows which moves won the game
	fmt.Println(unmarkedMoves(win, winBool))
	fmt.Println(winningMoves(win, instructions))
	unmarked := unmarkedMoves(win, winBool)
	moves := winningMoves(win, instructions)
	unmarkedSum := 0
	for _, number := range unmarked {
		unmarkedSum += number
	}
	fmt.Println(unmarkedSum * moves[len(moves)-1])
}
