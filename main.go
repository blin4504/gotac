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

func (game *gameState) placePiece(r, c int) {
	if r >= 3 || c >= 3 || r < 0 || c < 0 {
		fmt.Println("Out of Bounds")
		return
	}
	if game.playerTurn {
		game.board[r][c] = "O"
	} else {
		game.board[r][c] = "X"
	}
	fmt.Println()
	game.displayBoard()
	game.playerTurn = !game.playerTurn
}

func main() {
	game := gameState{}
	game.displayBoard()
	// game.placePiece(0, 2)

	//
	// game.placePiece(1, 2)

	game.playerTurn = !game.playerTurn
	game.placePiece(1, 2)

	// game.playerTurn = !game.playerTurn
	// game.placePiece(0, 1)
}
