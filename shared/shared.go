package shared

import (
	"bufio"
	"log"
	"os"
	"strconv"
)

func ProcessInputLines(filename string, callback func(string)) {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		callback(scanner.Text())
	}
}

func ReadNumbers(input string) []int {
	result := make([]int, 0)
	ProcessInputLines(input, func(line string) {
		n, err := strconv.Atoi(line)
		if err != nil {
			log.Fatal(err)
		}
		result = append(result, n)
	})
	return result
}

func ReadLines(input string) []string {
	result := make([]string, 0)
	ProcessInputLines(input, func(line string) {
		result = append(result, line)
	})
	return result
}
