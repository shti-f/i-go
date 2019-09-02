package modules

// Piece on square
type Piece struct {
	kind int
	str  string
}

// SquareStatus for placing sth
type SquareStatus struct {
	exist bool
	piece Piece
}

// Contents of a SquareStatus. (ゲームに合わせて設定を変える)
var (
	OUT   = SquareStatus{true, Piece{-1, " "}}
	SPACE = SquareStatus{false, Piece{0, "."}}
	WHITE = SquareStatus{true, Piece{1, "o"}}
	BLACK = SquareStatus{true, Piece{2, "*"}}
)

// ToString square
func (ss SquareStatus) ToString() string {
	return ss.piece.str
}
