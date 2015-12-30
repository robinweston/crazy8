package main

func (g *Grid) isRowValid(rowNumber int) bool {
	countersOnRow := 0

	for col := 0; col < 8; col++ {
		if g.Tiles[col][rowNumber] == true {
			countersOnRow++

			if countersOnRow > 1 {
				return false
			}
		}
	}

	return true
}

func (g *Grid) isColumnValid(colNumber int) bool {
	countersOnColumn := 0

	for row := 0; row < 8; row++ {
		if g.Tiles[colNumber][row] == true {
			countersOnColumn++

			if countersOnColumn > 1 {
				return false
			}
		}
	}

	return true
}

func (g *Grid) isLeftToTopDiagonalValid(rowNumber int) bool {
	countersOnDiagonal := 0

	for row, col := rowNumber, 0; row >= 0 && row < 8 && col >= 0 && col < 8; row, col = row-1, col+1 {
		if g.Tiles[col][row] == true {
			countersOnDiagonal++

			if countersOnDiagonal > 1 {
				return false
			}
		}
	}

	return true
}

func (g *Grid) isLeftToBottomDiagonalValid(rowNumber int) bool {

	countersOnDiagonal := 0

	for row, col := rowNumber, 0; row >= 0 && row < 8 && col >= 0 && col < 8; row, col = row+1, col+1 {

		if g.Tiles[col][row] == true {
			countersOnDiagonal++

			if countersOnDiagonal > 1 {
				return false
			}
		}
	}

	return true
}
