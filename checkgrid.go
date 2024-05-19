package main

import (
	"fmt"
)

// Function to check if there is a win
func checkWin(board [3][3]string, player string) bool {
	// Check rows and columns
	for i := 0; i < 3; i++ {
		if (board[i][0] == player && board[i][1] == player && board[i][2] == player) ||
			(board[0][i] == player && board[1][i] == player && board[2][i] == player) {
			return true
		}
	}

	// Check diagonals
	if (board[0][0] == player && board[1][1] == player && board[2][2] == player) ||
		(board[0][2] == player && board[1][1] == player && board[2][0] == player) {
		return true
	}

	return false
}

// Function to check if the grid is full
func isGridFull(board [3][3]string) bool {
	for _, row := range board {
		for _, cell := range row {
			if cell == "" {
				return false
			}
		}
	}
	return true
}

func main() {
	// Example Tic Tac Toe grid
	board := [3][3]string{
		{"X", "O", "X"},
		{"O", "X", "O"},
		{"X", "O", ""},
	}

	// Check if there is a win for X or O
	if checkWin(board, "X") {
		fmt.Println("Player X wins!")
	} else if checkWin(board, "O") {
		fmt.Println("Player O wins!")
	} else {
		fmt.Println("No winner yet.")
	}

	// Check if the grid is full
	if isGridFull(board) {
		fmt.Println("The grid is full.")
	} else {
		fmt.Println("The grid is not full yet.")
	}
}
