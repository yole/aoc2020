package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("day2input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	validPasswordsStep1 := 0
	validPasswordsStep2 := 0
	for scanner.Scan() {
		passwordPolicy := strings.Split(scanner.Text(), " ")
		expectedNumbers := strings.Split(passwordPolicy[0], "-")
		minNumber, err := strconv.Atoi(expectedNumbers[0])
		if err != nil {
			log.Fatal("Minimum number is not a number", expectedNumbers[0])
		}
		maxNumber, err := strconv.Atoi(expectedNumbers[1])
		if err != nil {
			log.Fatal("Maximum number is not a number", expectedNumbers[1])
		}
		letter := passwordPolicy[1][0]

		password := passwordPolicy[2]
		cnt := 0
		for i := 0; i < len(password); i++ {
			if password[i] == letter {
				cnt++
			}
		}
		if cnt >= minNumber && cnt <= maxNumber {
			validPasswordsStep1++
		}

		pos1Correct := password[minNumber-1] == letter
		pos2Correct := password[maxNumber-1] == letter
		if pos1Correct != pos2Correct {
			validPasswordsStep2++
		}

	}
	fmt.Println(validPasswordsStep1)
	fmt.Println(validPasswordsStep2)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
