package modules

import (
	"fmt"
	"strconv"
)

// Board is type for board
type Board struct {
	Content [][]SquareStatus
	LenX    int
	LenY    int
}

// Init returns initialized board
func new(LenX, LenY int) (board Board) {
	for y := 0; y < LenY+2; y++ {
		var row []SquareStatus
		for x := 0; x < LenX+2; x++ {
			if x == 0 || x == LenX+1 || y == 0 || y == LenY+1 {
				row = append(row, OUT)
			} else {
				row = append(row, SPACE)
			}
		}
		board.Content = append(board.Content, row)
	}
	board.LenX = LenX
	board.LenY = LenY
	return board
}

// Response type
type Response int

// responce status
const (
	OK          Response = 0
	EXIST       Response = 1
	OUTOFBOUNDS Response = 2
)

// Place x, y obj on SquareStatus
func (board *Board) place(x, y int, obj SquareStatus) Response {
	if x < 1 || x > board.LenX || y < 1 || y > board.LenY {
		return OUTOFBOUNDS
	}
	if board.Content[y][x].exist {
		return EXIST
	}
	board.Content[y][x] = obj
	return OK
}

// Remove obj
func (board *Board) remove(x, y int) {
	(board.Content)[y][x] = SPACE
}

// Disp board
func (board Board) Disp() {
	head := " "
	for x := 1; x <= board.LenX; x++ {
		head += strconv.Itoa(x)
	}
	fmt.Println(head)
	for y := 1; y <= board.LenY; y++ {
		row := strconv.Itoa(y)
		for x := 1; x <= board.LenX; x++ {
			row += board.Content[y][x].ToString()
		}
		fmt.Println(row)
	}
}
