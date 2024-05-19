package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"

	t "tictactoe/checks"
)

func startGrid() [][]string {
	grid := [][]string{
		{"", "", ""},
		{"", "", ""},
		{"", "", ""},
	}
	return grid
}

func atoi(x string) int {
	i, errr := strconv.Atoi(x)
	if errr != nil {
		println(errr)
		os.Exit(0)
	}
	if i > 2 {
		fmt.Println("Error: Not in grid")
		os.Exit(0)
	}
	return i
}

func getUserCordinates() (int, int) {
	input := ""

	fmt.Print("Your turrn: ")
	fmt.Scanln(&input)
	arrInput := strings.Split(input, ",")
	if len(arrInput) != 2 {
		fmt.Println(arrInput)
		fmt.Println("Error: Wrong Input Format")
		os.Exit(0)
	}
	i, j := atoi(arrInput[0]), atoi(arrInput[1])
	return i, j
}

func getCordinates(player string) (int, int) {
	i, j := 0, 0
	if player == "X" {
		i, j = getUserCordinates()
	} else {
		i, j = rand.Intn(3), rand.Intn(3)
	}
	return i, j
}

func isValidMove(i, j int, grid [][]string) bool {
	if grid[i][j] == "" {
		return true
	} else {
		return false
	}
}

func switchPlayer(player string) string {
	if player == "X" {
		player = "0"
	} else {
		player = "X"
	}
	return player
}

func printGrid(grid [][]string) {
	for _, row := range grid {
		for j, v := range row {
			if v == "" {
				fmt. Print(".")
			} else {
				fmt.Print(v)
			}
			if j < 2 {
				fmt.Print(" ")
			}
		}
		fmt.Println()
	}
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

		gridStatus := t.CheckGrid(grid, player)

		if gridStatus != "cont" {
			fmt.Println(gridStatus)
			return
		}
	}
}

func main() {
	printGrid(startGrid())
	playGame()
}
