package main

import (
	"testing"
)

func TestSolveDay17Step1(t *testing.T) {
	result := solveCube("day17testinput.txt", 3)
	if result != 112 {
		t.Errorf("Unexpected result %d", result)
	}
}

func TestSolveDay17Step2(t *testing.T) {
	result := solveCube("day17testinput.txt", 4)
	if result != 848 {
		t.Errorf("Unexpected result %d", result)
	}
}

func TestNeighbours(t *testing.T) {
	cube := loadCube("day17testinput.txt", 3)
	neighbors := cube.neighbors([]int{2, 2, -1})
	if neighbors != 3 {
		t.Errorf("Unexpected result %d", neighbors)
	}
}
