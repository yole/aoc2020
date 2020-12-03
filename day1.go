package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	file, err := os.Open("day1input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	values := make([]int, 0)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		i, err := strconv.Atoi(scanner.Text())
		if err != nil {
			log.Fatal("Not a number in input file", i)
		}
		values = append(values, i)
	}
	for i := 0; i < len(values); i++ {
		for j := 0; j < i; j++ {
			for k := 0; k < j; k++ {
				if values[i]+values[j]+values[k] == 2020 {
					fmt.Println(values[i] * values[j] * values[k])
				}
			}
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
