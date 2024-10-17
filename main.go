package main

import "fmt"

type gameState struct {
	board      [3][3]string
	playerTurn bool
}

func (game *gameState) displayBoard() {
	for i, row := range game.board {
		for j, el := range row {
			if el == "" {
				el = " "
			}
			if j < 2 {
				fmt.Printf(" %s |", el)
			} else {
				fmt.Printf(" %s ", el)
			}
		}
		fmt.Println()
		if i < 2 {
			fmt.Println("---|---|---")
		}
	}
	fmt.Println()
}

func (game *gameState) placePiece(r, c int) bool {
	if r >= 3 || c >= 3 || r < 0 || c < 0 {
		fmt.Println("Out of Bounds")
		return false
	}
	if game.board[r][c] != "" {
		fmt.Println("Cell already taken")
		return false
	}
	if game.playerTurn {
		game.board[r][c] = "X"
	} else {
		game.board[r][c] = "O"
	}
	fmt.Println()
	game.displayBoard()
	game.playerTurn = !game.playerTurn
	return true
}

func (game *gameState) checkWin() bool {
	var winConditions = [][]int{
		{0, 1, 2},
		{3, 4, 5},
		{6, 7, 8},
		{0, 3, 6},
		{1, 4, 7},
		{2, 5, 8},
		{0, 4, 8},
		{2, 4, 6},
	}
	for _, condition := range winConditions {
		if game.board[condition[0]/3][condition[0]%3] != "" &&
			game.board[condition[0]/3][condition[0]%3] == game.board[condition[1]/3][condition[1]%3] &&
			game.board[condition[1]/3][condition[1]%3] == game.board[condition[2]/3][condition[2]%3] {
			return true
		}
	}
	return false
}

func (game *gameState) checkDraw() bool {
	for _, row := range game.board {
		for _, el := range row {
			if el == "" {
				return false
			}
		}
	}
	return true
}

func main() {
	game := gameState{}
	game.displayBoard()
	game.playerTurn = true

	for {
		if game.playerTurn {
			fmt.Println("Player 1's turn (X)")
		} else {
			fmt.Println("Player 2's turn (O)")
		}

		var row, col int
		fmt.Print("Enter row (0-2): ")
		fmt.Scan(&row)
		fmt.Print("Enter column (0-2): ")
		fmt.Scan(&col)

		if game.placePiece(row, col) {
			if game.checkWin() {
				if !game.playerTurn {
					fmt.Println("Player 1 (X) wins!")
				} else {
					fmt.Println("Player 2 (O) wins!")
				}
				break
			}

			if game.checkDraw() {
				fmt.Println("It's a draw!")
				break
			}

		} else {
			fmt.Println("Invalid move. Try again.")
		}
	}
}
