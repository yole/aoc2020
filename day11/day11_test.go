package main

import (
	"aoc2020/shared"
	"testing"
)

func TestSolveDay11Step1(t *testing.T) {
	seats := solveDay11Step1("day11testinput.txt")
	if seats != 37 {
		t.Errorf("Unexpected result %d", seats)
	}
}

func TestSolveDay11Step2(t *testing.T) {
	seats := solveDay11Step2("day11testinput.txt")
	if seats != 26 {
		t.Errorf("Unexpected result %d", seats)
	}
}

func TestSolveDay11Step1Round1(t *testing.T) {
	seatMap := shared.ReadLines("day11testinput.txt")
	seatMap = verifyNextRound(t, seatMap, "day11round2.txt")
	verifyNextRound(t, seatMap, "day11round3.txt")
}

func verifyNextRound(t *testing.T, seatMap []string, expectedFile string) []string {
	newMap, changed := advanceSeatMap(seatMap, updateSeatStateStep1)
	if !changed {
		t.Errorf("Expected the map to change")
	}
	expectedSeatMap := shared.ReadLines(expectedFile)
	for i := range seatMap {
		if newMap[i] != expectedSeatMap[i] {
			t.Errorf("Unexpected result for row %d: %s", i, newMap[i])
		}
	}
	return newMap
}

func TestCountAdjacentSeats(t *testing.T) {
	seatMap := shared.ReadLines("day11round2.txt")
	count90 := countAdjacentSeats(seatMap, 9, 0)
	if count90 != 3 {
		t.Errorf("Unexpected result %d at 9, 0", count90)
	}
}

func TestCountVisibleOccupiedSeats(t *testing.T) {
	seatMap := shared.ReadLines("day11visible1.txt")
	count := countVisibleOccupiedSeats(seatMap, 3, 4)
	if count != 8 {
		t.Errorf("Unexpected result %d at 3, 4", count)
	}
}
