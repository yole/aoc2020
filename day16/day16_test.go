package main

import "testing"

func TestSolveDay16Step1(t *testing.T) {
	task := parseTask("day16testinput.txt")
	result := task.scanningErrorRate()
	if result != 71 {
		t.Errorf("Unexpected result %d", result)
	}
}

func TestSolveDay16Step2(t *testing.T) {
	task := parseTask("day16testinput2.txt")
	result := task.calcPossibleFields()
	if result[0] != 2 {
		t.Errorf("Unexpected result %d", result)
	}
}
