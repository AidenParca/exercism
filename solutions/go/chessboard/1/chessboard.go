package chessboard

// Declare a type named File which stores if a square is occupied by a piece - this will be a slice of bools
type File []bool
// Declare a type named Chessboard which contains a map of eight Files, accessed with keys from "A" to "H"
type Chessboard map[string]File
// {A:File,B:File,C:File,D:File,E:File,G:File,H:File}
// CountInFile returns how many squares are occupied in the chessboard,
// within the given file.
func CountInFile(cb Chessboard, file string) int {
    targetFile, ok := cb[file]
	if !ok {
		return 0
	}
    count := 0
    for _, isOccupied := range targetFile {
		if isOccupied {
			count++
		}
	}
	return count
}

// CountInRank returns how many squares are occupied in the chessboard,
// within the given rank.
func CountInRank(cb Chessboard, rank int) int {
    if rank > 8 || rank < 1 {
        return 0
    }
	count := 0
    for _, fileData := range cb {
		// A rank number (1-8) corresponds to a slice index (0-7).
		// We check if the square at the given rank's index is occupied.
		if fileData[rank-1] {
			count++
		}
	}
	return count
}


// CountAll should count how many squares are present in the chessboard.
func CountAll(cb Chessboard) int {
	return 64
}

// CountOccupied returns how many squares are occupied in the chessboard.
func CountOccupied(cb Chessboard) int {
	total := 0
    for i := 1 ; i <= 8 ; i ++ {
        total = total + CountInRank(cb,i)
    }
    return total 
}
