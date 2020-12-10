package main

import "testing"

func TestSolveDay8Step1(t *testing.T) {
	result := solveDay8Step1("day08testinput.txt")
	if result != 5 {
		t.Errorf("Unexpected result %d", result)
	}
}
