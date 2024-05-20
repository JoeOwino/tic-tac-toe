package tictactoe

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
)

func switchPlayer(player string) string {
	if player == "X" {
		player = "0"
	} else {
		player = "X"
	}
	return player
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
