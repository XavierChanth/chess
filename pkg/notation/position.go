package notation

import "fmt"

type (
	File     int // Letters a-h
	Rank     int // Numbers 1-8
	Position struct {
		f File
		r Rank
	}
)

const (
	fileOffset = 'a' - 1
	rankOffset = '1' - 1
)

// File
func FileFromRune(r rune) File {
	return File(r - fileOffset)
}

func (f File) Error() error {
	if f < 1 || f > 8 {
		return fmt.Errorf("invalid value for chess file %d", f)
	}
	return nil
}

func (f *File) Set(i int) error {
	*f = File(i)
	return f.Error()
}

func (f File) Value() int {
	return int(f)
}

func (f File) Rune() (rune, error) {
	err := f.Error()
	if err != nil {
		return -1, err
	}

	return fileOffset + rune(f.Value()), nil
}

// Rank
func RankFromRune(r rune) Rank {
	return Rank(r - rankOffset)
}

func (r Rank) Error() error {
	if r < 1 || r > 8 {
		return fmt.Errorf("invalid value for chess rank %d", r)
	}
	return nil
}

func (r *Rank) Set(i int) error {
	*r = Rank(i)
	return r.Error()
}

func (r Rank) Value() int {
	return int(r)
}

func (r Rank) Rune() (rune, error) {
	err := r.Error()
	if err != nil {
		return -1, err
	}

	return rankOffset + rune(r.Value()), nil
}

// Position
func (p Position) File() File {
	return p.f
}

func (p Position) Rank() Rank {
	return p.r
}

func (p Position) Error() error {
	ferr := p.f.Error()
	rerr := p.r.Error()
	if ferr == nil {
		return rerr
	}
	if rerr == nil {
		return ferr
	}
	return fmt.Errorf("invalid value for chess file %d & rank %d", p.f.Value(), p.r.Value())
}
