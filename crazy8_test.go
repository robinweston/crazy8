package main_test

import (
	. "github.com/robinweston/crazy8"
	"testing"
)

func Test_grid_with_two_counters_on_same_row_is_invalid_solution(t *testing.T) {

	g := NewGrid()

	g.PlaceCounterAt(0, 0)
	g.PlaceCounterAt(0, 1)

	if g.IsValidSolution() == true {
		t.Fatal("Expected IsValidSolution false got true")
	}
}

func Test_grid_with_two_counters_on_same_column_is_invalid_solution(t *testing.T) {

	g := NewGrid()

	g.PlaceCounterAt(0, 0)
	g.PlaceCounterAt(1, 0)

	if g.IsValidSolution() == true {
		t.Fatal("Expected IsValidSolution false got true")
	}
}

func Test_grid_with_two_counters_on_same_upwards_diagonal_is_invalid_solution(t *testing.T) {

	g := NewGrid()

	g.PlaceCounterAt(1, 3)
	g.PlaceCounterAt(3, 1)

	if g.IsValidSolution() == true {
		t.Fatal("Expected IsValidSolution false got true")
	}
}

func Test_grid_with_two_counters_on_same_upwards_diagonal_is_invalid_solution_2(t *testing.T) {

	g := NewGrid()

	g.PlaceCounterAt(7, 1)
	g.PlaceCounterAt(6, 2)

	if g.IsValidSolution() == true {
		t.Fatal("Expected IsValidSolution false got true")
	}
}

func Test_grid_with_two_counters_on_same_downwards_diagonal_is_invalid_solution(t *testing.T) {

	g := NewGrid()

	g.PlaceCounterAt(0, 4)
	g.PlaceCounterAt(2, 6)

	if g.IsValidSolution() == true {
		t.Fatal("Expected IsValidSolution false got true")
	}
}

func Test_grid_with_two_counters_on_same_downwards_diagonal_is_invalid_solution_2(t *testing.T) {

	g := NewGrid()
    
    g.PlaceCounterAt(5, 2)
	g.PlaceCounterAt(7, 4)

	if g.IsValidSolution() == true {
		t.Fatal("Expected IsValidSolution false got true")
	}
}