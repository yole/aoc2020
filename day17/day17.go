package main

import (
	"aoc2020/shared"
	"fmt"
	"log"
)
import "github.com/willf/bitset"

func main() {
	fmt.Println(solveCube("day17input.txt", 3))
	fmt.Println(solveCube("day17input.txt", 4))
}

func solveCube(filename string, dimensions int) uint {
	cube := loadCube(filename, dimensions)
	for i := 0; i < 6; i++ {
		cube = cube.advance()
	}
	return cube.cells.Count()
}

func loadCube(filename string, dimensions int) *Cube {
	lines := shared.ReadLines(filename)
	cube := allocateCube(len(lines), 0, dimensions)
	cube.initCube(lines)
	return cube
}

type Cube struct {
	cells      bitset.BitSet
	offset     int
	size       int
	dimensions int
}

func (cube *Cube) set(point []int) {
	for c := range point {
		if !cube.isValid(point[c]) {
			log.Fatal("Invalid point", point)
		}
	}
	index := cube.index(point)
	cube.cells.Set(index)
}

func (cube *Cube) isValid(coord int) bool {
	return coord >= -cube.offset && coord < cube.size-cube.offset
}

func (cube *Cube) get(point []int) bool {
	for c := range point {
		if !cube.isValid(point[c]) {
			return false
		}
	}
	return cube.cells.Test(cube.index(point))
}

func (cube *Cube) index(point []int) uint {
	result := 0
	for c := range point {
		result = result*cube.size + point[c] + cube.offset
	}
	return uint(result)
}

func allocateCube(size int, offset int, dimensions int) *Cube {
	var cube Cube
	cube.offset = offset
	cube.size = size
	cube.dimensions = dimensions
	return &cube
}

func (cube *Cube) initCube(initData []string) {
	for l := range initData {
		line := initData[l]
		for c := range line {
			if line[c] == '#' {
				coord := make([]int, cube.dimensions)
				coord[0] = c
				coord[1] = l
				cube.set(coord)
			}
		}
	}
}

func (cube *Cube) advance() *Cube {
	newCube := allocateCube(cube.size+2, cube.offset+1, cube.dimensions)
	offset := newCube.offset
	for n := 0; n < pow(newCube.size, cube.dimensions); n++ {
		coord := sliceAddInt(indexToCoord(n, newCube.size, cube.dimensions), -offset)
		value := cube.get(coord)
		neighbors := cube.neighbors(coord)
		if (value && neighbors == 2) || neighbors == 3 {
			newCube.set(coord)
		}
	}

	return newCube
}

func (cube *Cube) neighbors(coord []int) int {
	result := 0
	for i := 0; i < pow(3, cube.dimensions); i++ {
		delta := sliceAddInt(indexToCoord(i, 3, cube.dimensions), -1)
		if !allZeros(delta) && cube.get(sliceAdd(coord, delta)) {
			result++
		}
	}
	return result
}

func (cube *Cube) print() {
	offset := cube.offset
	for z := 0; z < cube.size; z++ {
		fmt.Printf("z=%d\n", z-cube.offset)
		for y := 0; y < cube.size; y++ {
			for x := 0; x < cube.size; x++ {
				if cube.get([]int{x - offset, y - offset, z - offset}) {
					fmt.Print("#")
				} else {
					fmt.Print(".")
				}
			}
			fmt.Println("")
		}
		fmt.Println("")
	}
}

func pow(n int, power int) int {
	result := 1
	for i := 0; i < power; i++ {
		result *= n
	}
	return result
}

func sliceAdd(a []int, b []int) []int {
	result := make([]int, len(a))
	for i := range a {
		result[i] = a[i] + b[i]
	}
	return result
}

func indexToCoord(n int, size int, dimensions int) []int {
	result := make([]int, dimensions)
	for i := range result {
		result[i] = n % size
		n /= size
	}
	return result
}

func sliceAddInt(a []int, b int) []int {
	result := make([]int, len(a))
	for i := range a {
		result[i] = a[i] + b
	}
	return result
}

func allZeros(a []int) bool {
	for x := range a {
		if a[x] != 0 {
			return false
		}
	}
	return true
}
