package main

import (
	"bufio"
	"fmt"
	"github.com/willf/bitset"
	"log"
	"os"
)

func main() {
	file, err := os.Open("day5input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	maxSeatId := uint(0)
	var takenSeats bitset.BitSet
	for scanner.Scan() {
		row, column := parseBoardingPass(scanner.Text())
		seatId := row*8 + column
		if seatId > maxSeatId {
			maxSeatId = seatId
		}
		takenSeats.Set(seatId)
	}

	freeSeat := uint(0)
	for i := uint(0); i < maxSeatId; i++ {
		freeSeat, _ = takenSeats.NextClear(i)
		if freeSeat != i {
			break
		}
	}

	fmt.Println(maxSeatId)
	fmt.Println(freeSeat)
}

func parseBoardingPass(text string) (uint, uint) {
	row := binarySpacePartition(text[0:7], 'F')
	col := binarySpacePartition(text[7:10], 'L')
	return row, col
}

func binarySpacePartition(text string, lowerHalfDeterminant uint8) uint {
	minRow := 0
	maxRow := 1 << len(text)
	for i := 0; i < len(text); i++ {
		if text[i] == lowerHalfDeterminant {
			maxRow = minRow + (maxRow-minRow)/2
		} else {
			minRow = minRow + (maxRow-minRow)/2
		}
	}
	return uint(minRow)
}
