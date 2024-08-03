package notation

type (
	Type  int
	Side  int
	Piece struct {
		Type Type
		Side Side
	}
)

// Type
const (
	Pawn Type = iota
	Knight
	Bishop
	Rook
	Queen
	King
)

func (t Type) Points() int {
	switch t {
	case Pawn:
		return 1
	case Knight:
	case Bishop:
		return 3
	case Rook:
		return 5
	case Queen:
		return 9
	}
	return 0
}

func (t Type) Rune() rune {
	switch t {
	case Knight:
		return 'N'
	case Bishop:
		return 'B'
	case Rook:
		return 'R'
	case Queen:
		return 'Q'
	case King:
		return 'K'
	}
	return -1
}

func TypeFromRune(r rune) Type {
	switch r {
	case Knight.Rune():
		return Knight
	case Bishop.Rune():
		return Bishop
	case Rook.Rune():
		return Rook
	case Queen.Rune():
		return Queen
	case King.Rune():
		return King
	}
	return Pawn
}

// Side
const (
	White Side = iota // could be a bool, but what if I want to implement 4 player later?
	Black
)
