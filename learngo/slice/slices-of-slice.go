package main

import (
	"fmt"
	"strings"
)

func main() {
	//Create  a tic-tac-toe board
	board := [][]string {
		[]string{"_","_","_"},
		[]string{"_","_","_"},
		[]string{"_","_","_"},
	}

	//The player take turns
	board[0][0] = "X"
	board[2][2] = "O"
	board[2][0] = "X"
	board[1][0] = "O"
	board[0][2] = "X"

	for i := 0; i < len(board); i++ {
		fmt.Printf("%d, %s\n",i,strings.Join(board[i], " "))
	}

}
