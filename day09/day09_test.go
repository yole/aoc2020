package main

import "testing"

func TestSolveDay9Step1(t *testing.T) {
	result := solveDay9Step1("day09testinput.txt", 5)
	if result != 127 {
		t.Errorf("Unexpected result %d", result)
	}
}

func TestSolveDay9Step2(t *testing.T) {
	result := solveDay9Step2("day09testinput.txt", 5)
	if result != 62 {
		t.Errorf("Unexpected result %d", result)
	}
}
