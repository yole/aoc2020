package main

import (
	"testing"
)

func TestSolveDay14Step1(t *testing.T) {
	distance := solveDay14Step1("day14testinput.txt")
	if distance != 165 {
		t.Errorf("Unexpected result %d", distance)
	}
}

func TestParseAndApplyMask(t *testing.T) {
	zeroMask, oneMask, _ := parseMask("XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX1XXXX0X")
	maskedValue := (11 | oneMask) & zeroMask
	if maskedValue != 73 {
		t.Errorf("Unexpected result %d", maskedValue)
	}
}

func TestSolveDay14Step2(t *testing.T) {
	distance := solveDay14Step2("day14testinput2.txt")
	if distance != 208 {
		t.Errorf("Unexpected result %d", distance)
	}
}
