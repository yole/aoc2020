package main

import (
	"aoc2020/shared"
	"fmt"
	"log"
	"strconv"
)

func main() {
	fmt.Println(solveDay9Step1("day09input.txt", 25))
	fmt.Println(solveDay9Step2("day09input.txt", 25))
}

func solveDay9Step1(input string, preamble int) int {
	numbers := readNumbers(input)
	return findWrongNumber(preamble, numbers)
}

func findWrongNumber(preamble int, numbers []int) int {
	for i := preamble; i < len(numbers); i++ {
		if !isSumOfAnyTwo(numbers[i], numbers, i-preamble, i) {
			return numbers[i]
		}
	}
	return -1
}

func solveDay9Step2(input string, preamble int) int {
	numbers := readNumbers(input)
	wrongNumber := findWrongNumber(preamble, numbers)
	for i := 0; i < len(numbers); i++ {
		min, max := findContiguousSum(numbers, i, wrongNumber)
		if min != -1 {
			return min + max
		}
	}
	return 0
}

func findContiguousSum(numbers []int, start int, expect int) (int, int) {
	sum := 0
	min := numbers[start]
	max := numbers[start]
	for i := start; i < len(numbers); i++ {
		if min > numbers[i] {
			min = numbers[i]
		}
		if max < numbers[i] {
			max = numbers[i]
		}
		sum += numbers[i]
		if sum == expect {
			return min, max
		}
		if sum > expect {
			return -1, -1
		}
	}
	return -1, -1
}

func isSumOfAnyTwo(n int, numbers []int, first int, last int) bool {
	for i := first; i < last; i++ {
		for j := first; j < i; j++ {
			if n == numbers[i]+numbers[j] {
				return true
			}
		}
	}
	return false
}

func readNumbers(input string) []int {
	result := make([]int, 0)
	shared.ProcessInputLines(input, func(line string) {
		n, err := strconv.Atoi(line)
		if err != nil {
			log.Fatal(err)
		}
		result = append(result, n)
	})
	return result
}
