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
	fmt.Println(solveDay13Step1("day13input.txt"))
	fmt.Println(solveDay13Step2("day13input.txt"))
}

func solveDay13Step1(filename string) int {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Scan()
	timestamp, err := strconv.Atoi(scanner.Text())
	if err != nil {
		log.Fatal(err)
	}
	scanner.Scan()
	busIds := strings.Split(scanner.Text(), ",")

	minBusId := -1
	minWaitTime := -1
	for i := range busIds {
		if busIds[i] == "x" {
			continue
		}
		busId, err := strconv.Atoi(busIds[i])
		if err != nil {
			log.Fatal(err)
		}
		waitTime := busId - (timestamp % busId)
		if minWaitTime == -1 || waitTime < minWaitTime {
			minBusId = busId
			minWaitTime = waitTime
		}
	}
	return minBusId * minWaitTime
}

func solveDay13Step2(filename string) int64 {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Scan()
	scanner.Scan()
	return multipleChineseRemainderFromString(scanner.Text())
}

func bezoutCoefficients(a int64, b int64) (int64, int64) {
	oldR, r := a, b
	oldS, s := int64(1), int64(0)
	oldT, t := int64(0), int64(1)

	for r != 0 {
		quotient := oldR / r
		oldR, r = r, oldR-quotient*r
		oldS, s = s, oldS-quotient*s
		oldT, t = t, oldT-quotient*t
	}
	return oldS, oldT
}

func chineseRemainder(a1 int64, n1 int64, a2 int64, n2 int64) int64 {
	m1, m2 := bezoutCoefficients(n1, n2)
	result := a1*m2*n2 + a2*m1*n1
	for result < 0 {
		result += n1 * n2
	}
	return result % (n1 * n2)
}

func multipleChineseRemainder(moduli []int64, remainders []int64) int64 {
	modulus, remainder := moduli[0], remainders[0]
	for i := 1; i < len(moduli); i++ {
		modulus, remainder = modulus*moduli[i], chineseRemainder(remainder, modulus, remainders[i], moduli[i])
	}
	return remainder
}

func multipleChineseRemainderFromString(line string) int64 {
	busIds := strings.Split(line, ",")
	moduli := make([]int64, 0)
	remainders := make([]int64, 0)
	for i := range busIds {
		if busIds[i] == "x" {
			continue
		}
		busId, err := strconv.Atoi(busIds[i])
		if err != nil {
			log.Fatal(err)
		}
		moduli = append(moduli, int64(busId))
		if i == 0 {
			remainders = append(remainders, 0)
		} else {
			remainders = append(remainders, int64(busId-(i%busId)))
		}
	}
	for i := range moduli {
		fmt.Printf("%d %d\n", remainders[i], moduli[i])
	}
	return multipleChineseRemainder(moduli, remainders)
}
