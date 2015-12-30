package main

import "fmt"

type Grid struct {
    Tiles [8][8]bool
}

func (g *Grid) PlaceCounterAt(x int, y int) {
    g.Tiles[x][y] = true
}

func (g *Grid) isRowValid(rowNumber int) bool {
    countersOnRow := 0
    
    for col := 0; col < 8; col++ {
        if g.Tiles[rowNumber][col] == true {
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
        if g.Tiles[row][colNumber] == true {
            countersOnColumn++
            
            if countersOnColumn > 1 {
                return false
            }
        }
    }
    
    return true
}

func (g *Grid) IsValidSolution() bool {
    for i := 0; i < 8; i++ {
        if g.isRowValid(i) == false {
            return false
        }
        
        if g.isColumnValid(i) == false {
            return false
        }
    }
    
    return true
}

func (g *Grid) Print() {
    fmt.Println(g.Tiles)
}

func NewGrid() *Grid {
    g := Grid{}
    return &g
}

func main() {
	fmt.Printf("Hello, world.\n")
}