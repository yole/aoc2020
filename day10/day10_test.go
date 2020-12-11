package main

import "testing"

func TestSolveDay10Step1(t *testing.T) {
	oneSteps, threeSteps := solveDay10Step1("day10testinput.txt")
	if oneSteps != 7 || threeSteps != 5 {
		t.Errorf("Unexpected result %d, %d", oneSteps, threeSteps)
	}
}

func TestSolveDay10Step1Input2(t *testing.T) {
	oneSteps, threeSteps := solveDay10Step1("day10testinput2.txt")
	if oneSteps != 22 || threeSteps != 10 {
		t.Errorf("Unexpected result %d, %d", oneSteps, threeSteps)
	}
}

func TestSolveDay10Step2(t *testing.T) {
	arrangements := solveDay10Step2("day10testinput.txt")
	if arrangements != 8 {
		t.Errorf("Unexpected result %d", arrangements)
	}
}

func TestSolveDay10Step2Input2(t *testing.T) {
	arrangements := solveDay10Step2("day10testinput2.txt")
	if arrangements != 19208 {
		t.Errorf("Unexpected result %d", arrangements)
	}
}
