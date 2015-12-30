package main

import (
	"fmt"
	"math/rand"
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
}

func (g *Grid) findFreeSpace() (int, int) {
	for {
		x := rand.Intn(8)
		y := rand.Intn(8)

		if g.Tiles[x][y] == false {
			return x, y
		}
	}
}

func (g *Grid) PlaceCountersRandomly() {
	for i := 0; i < 8; i++ {
		x, y := g.findFreeSpace()
		g.PlaceCounterAt(x, y)
	}
}

func NewGrid() *Grid {
	g := Grid{}
	return &g
}

func main() {
	// todo: multiple solutions. What if the same just rotated?
	var solutionsAttempted int64 = 0

	for {
		grid := NewGrid()
		grid.PlaceCountersRandomly()

		if grid.IsValidSolution() == true {
			fmt.Println("Solution found!!")
			grid.Print()
			return
		}

		solutionsAttempted++

		if solutionsAttempted%100000 == 0 {
			fmt.Println("Solutions attempted: ", solutionsAttempted)
		}
	}
}
