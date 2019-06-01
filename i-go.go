package main

import (
	"fmt"
	"strconv"
)

// Board は囲碁盤の様子です
type Board [][]string

func main() {
	var H, W = 5, 5
	var board = initBoard(H, W)
	board.disp()

	var isFinish = finisher()
	var isPlayer1 = true
	for {
		fmt.Println("your turn")
		var arg1, arg2 string
		fmt.Scan(&arg1)
		fmt.Println(arg1)
		if isFinish(arg1) {
			fmt.Println("game finished")
			board.disp()
			break
		}
		if arg1 != "p" {
			var a1, _ = strconv.Atoi(arg1)
			fmt.Scan(&arg2)
			var a2, _ = strconv.Atoi(arg2)
			board.turn(a1, a2, isPlayer1)
			isPlayer1 = !isPlayer1
		}
		board.disp()
	}
}

func initBoard(H, W int) (board Board) {
	for i := 0; i < H; i++ {
		var row []string
		for j := 0; j < W; j++ {
			row = append(row, ".")
		}
		board = append(board, row)
	}
	return board
}

func (board Board) disp() {
	for i := range board {
		var row string
		for _, v := range board[i] {
			row += v
		}
		fmt.Println(row)
	}
}

func (board *Board) turn(h, w int, isPlayer1 bool) {
	if isPlayer1 {
		(*board)[h][w] = "o"
	} else {
		(*board)[h][w] = "x"
	}
}

func finisher() func(arg string) bool {
	var passCount = 0
	return func(arg string) bool {
		if arg == "p" {
			passCount++
		} else {
			passCount = 0
		}
		if passCount >= 2 {
			return true
		}
		return false
	}
}
