package main

import (
	"fmt"
)

type Grid struct {
	Tiles [8][8]bool
}

func (g *Grid) PlaceCounterAt(x int, y int) {
	g.Tiles[x][y] = true
}

func (g *Grid) IsValidSolution() bool {
	for i := 0; i < 8; i++ {
		if g.isRowValid(i) == false {
			return false
		}

		if g.isColumnValid(i) == false {
			return false
		}

		if g.isLeftToTopDiagonalValid(i) == false {
			return false
		}

		if g.isLeftToBottomDiagonalValid(i) == false {
			return false
		}
	}

	return true
}

func (g *Grid) Print() {
	for row := 0; row < 8; row++ {
		fmt.Println()
		fmt.Print("[")

		for col := 0; col < 8; col++ {
			if g.Tiles[col][row] == true {
				fmt.Print(1)
			} else {
				fmt.Print(0)
			}

			if col < 7 {
				fmt.Print(" ")
			}
		}

		fmt.Print("]\r")
	}

	fmt.Println()
	fmt.Println()
}

func main() {

	solutions := make([]*Grid, 0)
	allGrids := GenerateAllPossibleGrids()

	for gridNumber, grid := range allGrids {

		if grid.IsValidSolution() == true {
			solutions = append(solutions, grid)
			fmt.Println("Solution found!!")
		}

		if gridNumber%100000 == 0 {
			fmt.Println("Solutions attempted: ", gridNumber)
		}
	}

    fmt.Println("Displaying solutions")
    fmt.Println()
    
	for i, solution := range solutions {
		fmt.Println("Solution", i)
		solution.Print()
	}
}
