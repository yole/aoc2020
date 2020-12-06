package main

import (
	"aoc2020/shared"
	"fmt"
	"math/bits"
)

func main() {
	anyoneResult := 0
	everyoneResult := 0
	anyoneMask := uint(0)
	everyoneMask := uint(0)
	groupStart := true

	shared.ProcessInputLines("day6input.txt", func(line string) {
		if len(line) == 0 {
			anyoneResult += bits.OnesCount(anyoneMask)
			everyoneResult += bits.OnesCount(everyoneMask)
			anyoneMask = 0
			everyoneMask = 0
			groupStart = true
		} else {
			lineMask := uint(0)
			for _, c := range line {
				mask := uint(1 << (c - 'a'))
				anyoneMask |= mask
				lineMask |= mask
			}
			if groupStart {
				everyoneMask = lineMask
				groupStart = false
			} else {
				everyoneMask &= lineMask
			}
		}
	})
	anyoneResult += bits.OnesCount(anyoneMask)
	everyoneResult += bits.OnesCount(everyoneMask)

	fmt.Println(anyoneResult)
	fmt.Println(everyoneResult)
}
