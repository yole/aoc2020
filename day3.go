package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	file, err := os.Open("day3input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	lines := make([]string, 0)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	trees := countTrees(lines, 3, 1)
	fmt.Println(trees)
	trees11 := countTrees(lines, 1, 1)
	trees51 := countTrees(lines, 5, 1)
	trees71 := countTrees(lines, 7, 1)
	trees12 := countTrees(lines, 1, 2)
	fmt.Println(trees * trees11 * trees51 * trees71 * trees12)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

func countTrees(lines []string, right int, down int) int {
	trees := 0
	x := 0
	for y := 0; y < len(lines); y += down {
		index := x % len(lines[0])
		if lines[y][index] == '#' {
			trees++
		}
		x += right
	}
	return trees
}
