package main

import (
	"testing"
)

func TestSolveDay12Step1(t *testing.T) {
	distance := solveDay12Step1("day12testinput.txt")
	if distance != 25 {
		t.Errorf("Unexpected result %d", distance)
	}
}

func TestSolveDay12Step2(t *testing.T) {
	distance := solveDay12Step2("day12testinput.txt")
	if distance != 286 {
		t.Errorf("Unexpected result %d", distance)
	}
}
