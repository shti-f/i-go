package main

import (
	"fmt"
	"strconv"

	"./modules"
)

// Board は囲碁盤の様子です

func main() {
	var board = modules.New()
	board.Disp()

	var isFinish = finisher()
	var isPlayer1 = true
	for {
		fmt.Println("your turn")
		var arg1, arg2 string
		fmt.Scan(&arg1)
		fmt.Println(arg1)
		if isFinish(arg1) {
			fmt.Println("game finished")
			board.Disp()
			break
		}
		if arg1 != "p" {
			var a1, _ = strconv.Atoi(arg1)
			fmt.Scan(&arg2)
			var a2, _ = strconv.Atoi(arg2)
			if isPlayer1 {
				board.Place(a1, a2, modules.BLACK)
			} else {
				board.Place(a1, a2, modules.WHITE)
			}

			isPlayer1 = !isPlayer1
		}
		board.Disp()
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
