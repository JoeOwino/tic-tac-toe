package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
)

func createGrid() [][]string {
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
	return i
}

func getCordinates(player string) (int, int) {
	i, j := 0, 0
	if player == "0" {
		i, j = rand.Intn(9), rand.Intn(9)
	} else {
		i, j = atoi(os.Args[1]), atoi(os.Args[2])
	}
	return i, j
}

func isValidMove(i, j int, grid [][]string) bool {
	if grid[i][j] == "" {
		return false
	} else {
		return true
	}
}

func switchPlayer(player string) string {
	if player == "X" {
		player = "0"
	} else {
		player = "X"
	}
}

func printGrid(grid [][]string) {
	for _, v := range grid {
		fmt.Println(v)
	}
}

func playGame() {
	grid := createGrid()
	player := "0"

	for {
		i, j := getCordinates(player)
		if !isValidMove(i, j, grid) {
			if player == "0" {
				continue
			} else {
				fmt.Print("Invalid Move, Try Again: ")
			}
		}
		grid[i][j] = player
		switchPlayer(player)
		if player == "X" {
			fmt.Println("Your turn: ")
		}

		gridStatus := checkGrid(grid)

		if gridStatus != "cont" {
			fmt.Println(gridStatus)
			return
		}
	}
}

func main() {
	player := "X"
	i, j := getCordinates(player)

	fmt.Println("(", i, ",", j, ")")
}
