package chessboard

// Declare a type named File which stores if a square is occupied by a piece - this will be a slice of bools
type File []bool

// Declare a type named Chessboard which contains a map of eight Files, accessed with keys from "A" to "H"
type Chessboard map[string]File

// CountInFile returns how many squares are occupied in the chessboard,
// within the given file.
func CountInFile(cb Chessboard, file string) int {
	var count int = 0
	val, ok := cb[file]
	if ok {
		for _, f := range val {
			if f {
				count += 1
			}
		}
	}
	return count
}

// CountInRank returns how many squares are occupied in the chessboard,
// within the given rank.
func CountInRank(cb Chessboard, rank int) int {
	count := 0
	if rank > 8 || rank < 1 {
		return count
	}

	for _, f := range cb {
		if f[rank-1] {
			count += 1
		}
	}
	return count
}

// CountAll should count how many squares are present in the chessboard.
func CountAll(cb Chessboard) int {
	count := 0
	for _, v := range cb {
		for range v {
			count += 1
		}
	}
	return count
}

// CountOccupied returns how many squares are occupied in the chessboard.
func CountOccupied(cb Chessboard) int {
	count := 0
	for _, v := range cb {
		for _, f := range v {
			if f {
				count += 1
			}
		}
	}
	return count
}
