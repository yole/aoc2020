package main

import (
	"testing"
)

func TestSolveDay15Step1(t *testing.T) {
	result := solveDay15Step1([]int{0, 3, 6}, 9)
	if result != 4 {
		t.Errorf("Unexpected result %d", result)
	}
}

func TestSolveDay15Step1_2(t *testing.T) {
	result := solveDay15Step1([]int{3, 1, 2}, 2020)
	if result != 1836 {
		t.Errorf("Unexpected result %d", result)
	}
}
