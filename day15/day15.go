package main

import "fmt"

func main() {
	fmt.Println(solveDay15Step1([]int{6, 4, 12, 1, 20, 0, 16}, 2020))
	fmt.Println(solveDay15Step1([]int{6, 4, 12, 1, 20, 0, 16}, 30000000))
}

func solveDay15Step1(startingNumbers []int, steps int) int {
	numbers := make([]int, steps)
	copy(numbers, startingNumbers)
	prevIndices := make(map[int]int)
	for i := 0; i < len(startingNumbers)-1; i++ {
		prevIndices[startingNumbers[i]] = i
	}
	lastValue := startingNumbers[len(startingNumbers)-1]
	for i := len(startingNumbers); i < steps; i++ {
		prevIndex, ok := prevIndices[lastValue]
		if !ok {
			numbers[i] = 0
		} else {
			numbers[i] = i - 1 - prevIndex
		}
		prevIndices[lastValue] = i - 1
		lastValue = numbers[i]
	}
	return numbers[steps-1]
}
