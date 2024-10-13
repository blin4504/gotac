package main

import "fmt"

type gameState struct {
	board      [3][3]string
	playerTurn string
}

func (game *gameState) displayBoard() {
	for i, row := range game.board {
		for j, el := range row {
			if j < 2 {
				fmt.Printf(" %s | ", el)
			} else {
				fmt.Printf(" %s ", el)
			}
		}
		if i < 2 {
			fmt.Println("\n--|---|--")
		}
	}
	fmt.Println()
}

func main() {
	game := gameState{}
	game.displayBoard()

}
