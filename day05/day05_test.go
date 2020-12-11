package main

import "testing"

func TestBinarySpacePartition(t *testing.T) {
	row, col := parseBoardingPass("BFFFBBFRRR")
	if row != 70 {
		t.Errorf("Row incorrect (expected 70): %d", row)
	}
	if col != 7 {
		t.Errorf("Col incorrect (expected 7): %d", col)
	}
}
