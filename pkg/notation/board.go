package notation

type BoardEntry struct {
	Piece    Piece
	Position Position
}

type OptionalPiece interface{}

type (
	// Representing a board like this is preferred for:
	// - Smaller memory footprint
	// - Less iterations to process all pieces in the board
	BoardByEntries []BoardEntry

	// Representing a board like this is preferred for:
	// - Consistent size and order
	// - Easier to lookup if a piece exists at a position
	BoardByPosition map[Position]OptionalPiece
)

// So which type should be preferred?
// - Probably BoardByEntries:
//   - If we want a consistent encoding, sort the array by position, and
//     (optionally) pad it with empty pieces
//   - There are probably other optimizations (depending on the context) which
//     can made to circumvent the slow(er) position lookup time.
//   - And even though position lookup time is slower, it is equal to piece
//     lookup time, where as you would have to process the whole map with
//     BoardByPosition to lookup an individual piece
