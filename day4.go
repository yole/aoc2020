package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("day4input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	validPassportsStep1 := 0
	validPassportsStep2 := 0
	scanner := bufio.NewScanner(file)
	currentPassport := make(map[string]string)
	for scanner.Scan() {
		line := scanner.Text()
		if len(line) == 0 {
			step1, step2 := validatePassport(currentPassport)
			validPassportsStep1 += step1
			validPassportsStep2 += step2
			currentPassport = make(map[string]string)
		} else {
			parsePassportLine(line, currentPassport)
		}
	}
	step1, step2 := validatePassport(currentPassport)
	validPassportsStep1 += step1
	validPassportsStep2 += step2
	fmt.Println(validPassportsStep1)
	fmt.Println(validPassportsStep2)
}

func parsePassportLine(line string, passport map[string]string) {
	fields := strings.Split(line, " ")
	for _, field := range fields {
		keyValue := strings.Split(field, ":")
		passport[keyValue[0]] = keyValue[1]
	}
}

var requiredFields = []string{"byr", "iyr", "eyr", "hgt", "hcl", "ecl", "pid"}

func validatePassport(passport map[string]string) (int, int) {
	for _, field := range requiredFields {
		_, ok := passport[field]
		if !ok {
			return 0, 0
		}
	}

	validFields := validateNumber(passport["byr"], 4, 1920, 2002)
	validFields += validateNumber(passport["iyr"], 4, 2010, 2020)
	validFields += validateNumber(passport["eyr"], 4, 2020, 2030)
	validFields += validateHeight(passport["hgt"])
	validFields += validateRegexp(passport["hcl"], "#[0-9a-f]{6}")
	validFields += validateRegexp(passport["ecl"], "amb|blu|brn|gry|grn|hzl|oth")
	validFields += validateRegexp(passport["pid"], "[0-9]{9}")

	if validFields == len(requiredFields) {
		return 1, 1
	}
	return 1, 0
}

func validateNumber(value string, digits int, min int, max int) int {
	if len(value) != digits {
		return 0
	}
	i, err := strconv.Atoi(value)
	if err != nil {
		return 0
	}
	if i < min || i > max {
		return 0
	}
	return 1
}

func validateHeight(s string) int {
	if strings.HasSuffix(s, "cm") {
		return validateNumber(strings.TrimSuffix(s, "cm"), 3, 150, 193)
	}
	if strings.HasSuffix(s, "in") {
		return validateNumber(strings.TrimSuffix(s, "in"), 2, 59, 76)
	}
	return 0
}

func validateRegexp(s string, pattern string) int {
	matched, err := regexp.MatchString("^"+pattern+"$", s)
	if err != nil || !matched {
		return 0
	}
	return 1
}
