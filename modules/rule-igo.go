package modules

// New board
func New() (board Board) {
	return new(5, 5)
	// return new(19, 19)
}

// Place stone at x,y
func (board *Board) Place(x, y int, stone SquareStatus) Response {
	res := board.place(x, y, stone)
	if res == OK {

	}
	return res
}

// IsPrisoner checks 4 neighbors of (x,y)ここクロージャーにするのが良さげ
func (board Board) IsPrisoner(x, y int, stone SquareStatus) bool {
	if !board.Content[y][x+1].exist {
		return false
	}
	if !board.Content[y][x-1].exist {
		return false
	}
	if !board.Content[y+1][x].exist {
		return false
	}
	if !board.Content[y-1][x].exist {
		return false
	}

	if board.Content[y][x+1].piece.kind = stone.piece.kind {
		IsPrisoner(y, x+1)
	}
	return true
}
