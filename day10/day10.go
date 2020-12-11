package main

import (
	"aoc2020/shared"
	"fmt"
	"log"
	"math"
	"sort"
)

func main() {
	oneSteps, threeSteps := solveDay10Step1("day10input.txt")
	fmt.Println(oneSteps * threeSteps)
	arrangements := solveDay10Step2("day10input.txt")
	fmt.Println(arrangements)
}

func solveDay10Step1(input string) (int, int) {
	numbers := shared.ReadNumbers(input)
	sort.Ints(numbers)
	oneSteps := 0
	threeSteps := 0
	for i := 0; i < len(numbers); i++ {
		var prevNumber int
		if i == 0 {
			prevNumber = 0
		} else {
			prevNumber = numbers[i-1]
		}
		if numbers[i]-prevNumber == 1 {
			oneSteps++
		} else if numbers[i]-prevNumber == 3 {
			threeSteps++
		} else {
			log.Fatal("not one or three steps")
		}
	}
	return oneSteps, threeSteps + 1 /* Finally, your device's built-in adapter is always 3 higher than the highest adapter */
}

func solveDay10Step2(input string) uint64 {
	numbers := shared.ReadNumbers(input)
	sort.Ints(numbers)

	sequencesOfOneLengths := make([]int, 6)
	currentSequenceLength := 1
	for i := range numbers {
		var prevNumber int
		if i == 0 {
			prevNumber = 0
		} else {
			prevNumber = numbers[i-1]
		}

		if numbers[i]-prevNumber == 1 {
			currentSequenceLength++
		} else if numbers[i]-prevNumber == 3 {
			for l := len(sequencesOfOneLengths); l <= currentSequenceLength; l++ {
				sequencesOfOneLengths = append(sequencesOfOneLengths, 0)
			}
			sequencesOfOneLengths[currentSequenceLength]++
			currentSequenceLength = 1
		}
	}
	for l := len(sequencesOfOneLengths); l <= currentSequenceLength; l++ {
		sequencesOfOneLengths = append(sequencesOfOneLengths, 0)
	}
	sequencesOfOneLengths[currentSequenceLength]++
	result := uint64(math.Pow(2, float64(sequencesOfOneLengths[3]))) *
		uint64(math.Pow(4, float64(sequencesOfOneLengths[4]))) *
		uint64(math.Pow(7, float64(sequencesOfOneLengths[5])))
	return result
}
