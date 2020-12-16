package main

import (
	"aoc2020/shared"
	"fmt"
	"log"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	fmt.Println(solveDay14Step1("day14input.txt"))
	fmt.Println(solveDay14Step2("day14input.txt"))
}

func solveDay14Step1(input string) int64 {
	return processInstructions(input, step1Executor)
}

func processInstructions(input string, executor func(map[int64]int64, int64, int64, int64, int64, []int)) int64 {
	pattern := regexp.MustCompile("mem\\[(\\d+)] = (\\d+)")
	memory := make(map[int64]int64)
	zeroMask := int64(0)
	oneMask := int64(0)
	floatingBits := make([]int, 0)
	shared.ProcessInputLines(input, func(line string) {
		if strings.HasPrefix(line, "mask") {
			zeroMask, oneMask, floatingBits = parseMask(line)
		} else {
			address, value := parseMem(line, pattern)
			executor(memory, address, value, zeroMask, oneMask, floatingBits)
		}
	})
	return sumOfMemory(memory)
}

func step1Executor(memory map[int64]int64, address int64, value int64, zeroMask int64, oneMask int64, _ []int) {
	memory[address] = (value | oneMask) & zeroMask
}

func solveDay14Step2(input string) int64 {
	return processInstructions(input, step2Executor)
}

func step2Executor(memory map[int64]int64, address int64, value int64, _ int64, oneMask int64, floatingBits []int) {
	baseAddress := address | oneMask
	for i := 0; i < 1<<len(floatingBits); i++ {
		address := baseAddress
		for k := range floatingBits {
			bit := floatingBits[k]
			if i&(1<<k) == 0 {
				address &= not(1 << bit)
			} else {
				address |= 1 << bit
			}
		}
		memory[address] = value
	}
}

func sumOfMemory(memory map[int64]int64) int64 {
	result := int64(0)
	for _, value := range memory {
		result += value
	}
	return result
}

func parseMask(line string) (int64, int64, []int) {
	mask := line[7:]
	if len(mask) != 36 {
		log.Fatalf("Unexpected mask length %d", len(mask))
	}
	zeroMask := int64(0)
	oneMask := int64(0)
	floatingBits := make([]int, 0)
	for i := 0; i < 36; i++ {
		if mask[36-i-1] == '0' {
			zeroMask |= 1 << i
		}
		if mask[36-i-1] == '1' {
			oneMask |= 1 << i
		}
		if mask[36-i-1] == 'X' {
			floatingBits = append(floatingBits, i)
		}
	}
	return not(zeroMask), oneMask, floatingBits
}

func not(zeroMask int64) int64 {
	return (int64(1)<<37 - 1) ^ zeroMask
}

func parseMem(line string, pattern *regexp.Regexp) (int64, int64) {
	matches := pattern.FindStringSubmatch(line)
	address, err := strconv.ParseInt(matches[1], 10, 64)
	if err != nil {
		log.Fatal(err)
	}
	value, err := strconv.ParseInt(matches[2], 10, 64)
	if err != nil {
		log.Fatal(err)
	}
	return address, value
}
