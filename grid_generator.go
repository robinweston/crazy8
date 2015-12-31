package main

import "fmt"

func NewGrid() *Grid {
	g := Grid{}
	return &g
}

func GenerateAllPossibleGrids() []*Grid {
	fmt.Println("Generating all possible grids")

	grids := make([]*Grid, 0)

	for i0 := 0; i0 < 8; i0++ {
		for i1 := 0; i1 < 8; i1++ {
			for i2 := 0; i2 < 8; i2++ {
				for i3 := 0; i3 < 8; i3++ {
					for i4 := 0; i4 < 8; i4++ {
						for i5 := 0; i5 < 8; i5++ {
							for i6 := 0; i6 < 8; i6++ {
								for i7 := 0; i7 < 8; i7++ {
									grid := NewGrid()
									grid.PlaceCounterAt(0, i0)
									grid.PlaceCounterAt(1, i1)
									grid.PlaceCounterAt(2, i2)
									grid.PlaceCounterAt(3, i3)
									grid.PlaceCounterAt(4, i4)
									grid.PlaceCounterAt(5, i5)
									grid.PlaceCounterAt(6, i6)
									grid.PlaceCounterAt(7, i7)

									grids = append(grids, grid)

									if len(grids)%100000 == 0 {
										fmt.Println("Grids generated:", len(grids))
									}
								}
							}
						}
					}
				}
			}
		}
	}

	return grids
}
