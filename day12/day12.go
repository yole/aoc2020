package main

import (
	"aoc2020/shared"
	"fmt"
	"log"
	"strconv"
)

func main() {
	fmt.Println(solveDay12Step1("day12input.txt"))
	fmt.Println(solveDay12Step2("day12input.txt"))
}

func solveDay12Step1(input string) int {
	x := 0
	y := 0
	dir := 90

	shared.ProcessInputLines(input, func(line string) {
		operand, err := strconv.Atoi(line[1:])
		if err != nil {
			log.Fatal("Not a number in commands list")
		}
		switch line[0] {
		case 'N':
			y += operand
			break
		case 'S':
			y -= operand
			break
		case 'E':
			x += operand
			break
		case 'W':
			x -= operand
			break
		case 'L':
			dir -= operand
			if dir < 0 {
				dir += 360
			}
			break
		case 'R':
			dir += operand
			if dir >= 360 {
				dir -= 360
			}
			break
		case 'F':
			switch dir {
			case 0:
				y += operand
				break
			case 90:
				x += operand
				break
			case 180:
				y -= operand
				break
			case 270:
				x -= operand
				break
			default:
				log.Fatalf("Unhandled direction %d", dir)
			}
			break
		}
	})
	return abs(x) + abs(y)
}

func solveDay12Step2(input string) int {
	x := 0
	y := 0
	wpdx := 10
	wpdy := 1

	shared.ProcessInputLines(input, func(line string) {
		operand, err := strconv.Atoi(line[1:])
		if err != nil {
			log.Fatal("Not a number in commands list")
		}
		switch line[0] {
		case 'N':
			wpdy += operand
			break
		case 'S':
			wpdy -= operand
			break
		case 'E':
			wpdx += operand
			break
		case 'W':
			wpdx -= operand
			break
		case 'L':
			for i := 0; i < operand/90; i++ {
				wpdx, wpdy = -wpdy, wpdx
			}
			break
		case 'R':
			for i := 0; i < operand/90; i++ {
				wpdx, wpdy = wpdy, -wpdx
			}
			break
		case 'F':
			x += wpdx * operand
			y += wpdy * operand
			break
		}
	})
	return abs(x) + abs(y)
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
