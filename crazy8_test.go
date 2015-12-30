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

func Test_grid_with_two_counters_on_same_left_to_top_diagonal_is_invalid_solution(t *testing.T) {

	g := NewGrid()

	g.PlaceCounterAt(1, 3)
	g.PlaceCounterAt(3, 1)

	if g.IsValidSolution() == true {
		t.Fatal("Expected IsValidSolution false got true")
	}
}

func Test_grid_with_two_counters_on_same_left_to_bottom_diagonal_is_invalid_solution(t *testing.T) {

	g := NewGrid()

	g.PlaceCounterAt(0, 4)
	g.PlaceCounterAt(2, 6)

	if g.IsValidSolution() == true {
		t.Fatal("Expected IsValidSolution false got true")
	}
}
