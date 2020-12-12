package main

import (
	"aoc2020/shared"
	"fmt"
	"log"
)

func main() {
	fmt.Println(solveDay11Step1("day11input.txt"))
	fmt.Println(solveDay11Step2("day11input.txt"))
}

func solveDay11Step1(input string) int {
	return iterateSeatMap(input, updateSeatStateStep1)
}

func solveDay11Step2(input string) int {
	return iterateSeatMap(input, updateSeatStateStep2)
}

func iterateSeatMap(input string, updateFunc func([]string, int, int) uint8) int {
	seatMap := shared.ReadLines(input)
	for {
		newMap, changed := advanceSeatMap(seatMap, updateFunc)
		if !changed {
			break
		}
		seatMap = newMap
	}
	return countSeats(seatMap)
}

func advanceSeatMap(seatMap []string, updateFunc func([]string, int, int) uint8) ([]string, bool) {
	newSeatMap := make([]string, len(seatMap))
	changed := false
	for y := 0; y < len(seatMap); y++ {
		newSeatRow := ""
		for x := 0; x < len(seatMap[y]); x++ {
			oldState := getSeatAt(seatMap, x, y)
			newState := updateFunc(seatMap, x, y)
			newSeatRow += string(newState)
			if oldState != newState {
				changed = true
			}
		}
		newSeatMap[y] = newSeatRow
	}
	return newSeatMap, changed
}

func countAdjacentSeats(seatMap []string, x int, y int) int {
	result := 0
	for dx := x - 1; dx <= x+1; dx++ {
		for dy := y - 1; dy <= y+1; dy++ {
			if (dx != x || dy != y) && getSeatAt(seatMap, dx, dy) == '#' {
				result++
			}
		}
	}
	return result
}

func updateSeatStateStep1(seatMap []string, x int, y int) uint8 {
	adjacentSeats := countAdjacentSeats(seatMap, x, y)
	return updateSeatState(seatMap[y][x], adjacentSeats, 4)
}

func updateSeatState(oldState uint8, adjacentSeats int, adjacentSeatsThreshold int) uint8 {
	if oldState == '.' {
		return '.'
	}
	if oldState == 'L' {
		if adjacentSeats == 0 {
			return '#'
		}
		return 'L'
	}
	if oldState == '#' {
		if adjacentSeats >= adjacentSeatsThreshold {
			return 'L'
		}
		return '#'
	}
	log.Fatalf("Unknown seat state %c", oldState)
	return '.'
}

func updateSeatStateStep2(seatMap []string, x int, y int) uint8 {
	visibleSeats := countVisibleOccupiedSeats(seatMap, x, y)
	return updateSeatState(seatMap[y][x], visibleSeats, 5)
}

func countVisibleOccupiedSeats(seatMap []string, x int, y int) int {
	result := 0
	for dx := -1; dx <= 1; dx++ {
		for dy := -1; dy <= 1; dy++ {
			if (dx != 0 || dy != 0) && canSeeOccupiedSeat(seatMap, x, y, dx, dy) {
				result++
			}
		}
	}
	return result
}

func canSeeOccupiedSeat(seatMap []string, x int, y int, dx int, dy int) bool {
	for step := 1; ; step++ {
		cx := x + dx*step
		cy := y + dy*step
		if cx < 0 || cy < 0 || cy >= len(seatMap) || cx >= len(seatMap[cy]) {
			return false
		}
		seat := seatMap[cy][cx]
		if seat == '#' {
			return true
		}
		if seat == 'L' {
			return false
		}
	}
}

func countSeats(seatMap []string) int {
	result := 0
	for y := 0; y < len(seatMap); y++ {
		for x := 0; x < len(seatMap[y]); x++ {
			if seatMap[y][x] == '#' {
				result++
			}
		}
	}
	return result
}

func getSeatAt(seatMap []string, x int, y int) uint8 {
	if y < 0 || y >= len(seatMap) || x < 0 || x >= len(seatMap[y]) {
		return '.'
	}
	return seatMap[y][x]
}
