package tictactoe

import (
	"fmt"
)

func startGrid() [][]string {
	grid := [][]string{
		{"", "", ""},
		{"", "", ""},
		{"", "", ""},
	}
	return grid
}

func playGame() {
	grid := startGrid()
	player := "X"

	for {
		i, j := getCordinates(player)
		if !isValidMove(i, j, grid) {
			if player == "0" {
				continue
			} else {
				fmt.Print("Invalid Move, Try Again: ")
				continue
			}
		}
		grid[i][j] = player
		if player == "0" {
			printGrid(grid)
		}
		player = switchPlayer(player)

		gridStatus := checkGrid(grid, player)

		if gridStatus != "cont" {
			printGrid(grid)
			fmt.Println(gridStatus)
			return
		}
	}
}

func printGrid(grid [][]string) {
	fmt.Println("     0   1   2  ")
	fmt.Println("   _____________")

	for i, row := range grid {
		fmt.Print(i)
		fmt.Print("  | ")
		for j, v := range row {
			if v == "" {
				fmt.Print(".")
			} else {
				fmt.Print(v)
			}
			if j < 2 {
				fmt.Print(" | ")
			}
		}
		fmt.Println(" |")
	}
}

func TicTacToe() {
	printGrid(startGrid())
	playGame()
}
