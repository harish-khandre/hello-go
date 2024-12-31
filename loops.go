package main

import (
	"fmt"
	"strings"
)

func ticTacToeBoard() {
	// Create a tic-tac-toe board.
	board := [][]string{
		{"_", "_", "_"},
		{"_", "_", "_"},
		{"_", "_", "_"},
	}

	// The players take turns.
	board[0][0] = "X"
	board[2][2] = "O"
	board[1][2] = "X"
	board[1][0] = "O"
	board[0][2] = "X"

	for i := 0; i < len(board); i++ {
		fmt.Printf("%s\n", strings.Join(board[i], " "))
	}
	pow := []int{1, 2, 4, 8, 16, 32, 64, 128}

	for i, v := range pow {
		fmt.Printf("2**%d = %d\n", i, v)
	}

	pew := make([]int, 10)
	for i := range pew {
		pew[i] = 1 << uint(i) // == 2**i
	}
	for _, value := range pew {
		fmt.Printf("%d\n", value)
	}
}
