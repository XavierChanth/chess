package notation

import (
	"fmt"
	"strconv"
	"strings"
)

const (
	CounterRune   rune = '.'
	CaptureRune   rune = 'x'
	CheckRune     rune = '+'
	CheckmateRune rune = '#'
)

type (
	None          struct{}
	Notation      interface{}
	Disambiguator interface{} // Could be None, Rank, File or Position
)

type BasicMove struct {
	From        Disambiguator
	Piece       Piece
	To          Position
	IsCapture   bool
	IsCheck     bool
	IsCheckmate bool
}

type MoveSequence struct {
	Number int
}

type KingsideCastle struct {
	Side Side
}

type QueensideCastle struct {
	Side Side
}

type Win struct {
	Side Side
}

func ParseMove(str string, side Side) (Notation, error) {
	str = strings.Trim(str, " \t")
	switch str {
	case "O-O":
	case "0-0":
		return KingsideCastle{side}, nil
	case "O-O-O":
	case "0-0-0":
		return QueensideCastle{side}, nil
	case "1-0":
		return Win{White}, nil
	case "0-1":
		return Win{Black}, nil
	}

	runes := []rune(str)
	end := len(runes) - 1

	switch {
	case len(runes) < 1:
		return nil, fmt.Errorf("empty move sequence")
	case runes[0] >= '1' && runes[0] <= '9' && runes[end] == '.':
		n, err := strconv.Atoi(string(runes[0 : end-1]))
		if err != nil {
			return nil, err
		}
		return MoveSequence{n}, nil
	}

	m := BasicMove{}
	m.Piece.Side = side

	fs, rs := []File{}, []Rank{}

	for i, r := range runes {
		// The first character will always tell us the piece type
		if i == 0 {
			t := TypeFromRune(r)
			m.Piece.Type = t
			if t != Pawn {
				continue
			}
		}

		// Parse augment symbols (capture, check, checkmate)
		cont := true
		switch r {
		case CaptureRune:
			m.IsCapture = true
		case CheckRune:
			m.IsCheck = true
		case CheckmateRune:
			m.IsCheckmate = true
		default:
			cont = false
		}
		if cont {
			continue
		}

		switch {
		case r >= 'a' && r <= 'h':
			fs = append(fs, FileFromRune(r))
		case r >= '1' && r <= '8':
			rs = append(rs, RankFromRune(r))
		default:
			return nil, fmt.Errorf("unrecognized character at index %d, character %c", i, r)
		}
	}

	df := len(fs) > 1 // Disambiguator for file
	dr := len(rs) > 1 // Disambiguator for rank

	if !df && !dr {
		m.From = None{}
		m.To = Position{fs[0], rs[0]}
	} else if !dr {
		m.From = File(fs[0])
		m.To = Position{fs[1], rs[0]}
	} else if !df {
		m.From = Rank(rs[0])
		m.To = Position{fs[0], rs[1]}
	} else {
		m.From = Position{fs[0], rs[0]}
		m.To = Position{fs[1], rs[1]}
	}

	return m, nil
}
