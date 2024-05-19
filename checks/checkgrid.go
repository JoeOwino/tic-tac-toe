package tictactoe

// Function to check if there is a win
func isWin(grid [][]string, player string) bool {
	// Check rows and columns
	for i := 0; i < 3; i++ {
		if (grid[i][0] == player && grid[i][1] == player && grid[i][2] == player) ||
			(grid[0][i] == player && grid[1][i] == player && grid[2][i] == player) {
			return true
		}
	}

	// Check diagonals
	if (grid[0][0] == player && grid[1][1] == player && grid[2][2] == player) ||
		(grid[0][2] == player && grid[1][1] == player && grid[2][0] == player) {
		return true
	}

	return false
}

// Function to check if the grid is full
func isFull(grid [][]string) bool {
	for _, row := range grid {
		for _, cell := range row {
			if cell == "" {
				return false
			}
		}
	}
	return true
}

func CheckGrid(grid [][]string, player string) string {
	switch {
	case isWin(grid, player):
		if player == "X" {
			return "You won!!"
		} else {
			return "You wo!!"
		}
	case isFull(grid):
		return "Drow!!"
	}
	return "cont"
}
