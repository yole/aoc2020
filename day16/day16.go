package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"math/bits"
	"os"
	"strconv"
	"strings"
)

type Range struct {
	min int
	max int
}

type Field struct {
	name   string
	ranges []Range
}

type Task struct {
	fields        []Field
	yourTicket    []int
	nearbyTickets [][]int
}

func main() {
	task := parseTask("day16input.txt")
	fmt.Println(task.scanningErrorRate())
	fmt.Println(task.solveStep2())
}

func parseTask(filename string) Task {
	var t Task
	t.fields = make([]Field, 0)
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if len(line) == 0 {
			break
		}
		t.fields = append(t.fields, parseField(line))
	}

	scanner.Scan()
	section := scanner.Text()
	if section != "your ticket:" {
		log.Fatalf("Unexpected line %s", section)
	}
	scanner.Scan()
	t.yourTicket = parseTicket(scanner.Text())
	scanner.Scan()
	scanner.Scan()

	section = scanner.Text()
	if section != "nearby tickets:" {
		log.Fatalf("Unexpected line %s", section)
	}
	t.nearbyTickets = make([][]int, 0)
	for scanner.Scan() {
		t.nearbyTickets = append(t.nearbyTickets, parseTicket(scanner.Text()))
	}

	return t
}

func parseField(line string) Field {
	f := strings.Split(line, ": ")
	r := strings.Split(f[1], " or ")
	var result Field
	result.name = f[0]
	result.ranges = parseRanges(r)
	return result
}

func parseRanges(r []string) []Range {
	result := make([]Range, len(r))
	for i := range r {
		ends := strings.Split(r[i], "-")
		min, err := strconv.Atoi(ends[0])
		if err != nil {
			log.Fatal("Unparseable range start")
		}
		max, err := strconv.Atoi(ends[1])
		if err != nil {
			log.Fatal("Unparseable range end")
		}
		result[i] = Range{min, max}
	}
	return result
}

func parseTicket(line string) []int {
	fields := strings.Split(line, ",")
	result := make([]int, len(fields))
	for i := range fields {
		field, err := strconv.Atoi(fields[i])
		if err != nil {
			log.Fatalf("Unparseable ticket value %s", fields[i])
		}
		result[i] = field
	}
	return result
}

func (task *Task) scanningErrorRate() int {
	result := 0
	for t := range task.nearbyTickets {
		errorValue, _ := task.validateTicket(task.nearbyTickets[t])
		result += errorValue
	}
	return result

}

func (task *Task) validateTicket(ticket []int) (int, bool) {
	result := 0
	valid := true
	for n := range ticket {
		if !task.isValidForAnyField(ticket[n]) {
			result += ticket[n]
			valid = false
		}
	}
	return result, valid
}

func (task *Task) isValidForAnyField(n int) bool {
	for i := range task.fields {
		if task.fields[i].isValidValue(n) {
			return true
		}
	}
	return false
}

func (task *Task) solveStep2() int {
	possibleFields := task.calcPossibleFields()
	result := 1
	for f := range task.fields {
		if strings.HasPrefix(task.fields[f].name, "departure") {
			index := indexOf(possibleFields, 1<<f)
			if index == -1 {
				log.Fatalf("Did not converge on single possible index for field %d", f)
			}
			result *= task.yourTicket[index]
		}
	}
	return result
}

func (task *Task) calcPossibleFields() []uint {
	possibleFields := make([]uint, len(task.fields))
	for i := range possibleFields {
		possibleFields[i] = 1<<len(task.fields) - 1
	}

	for t := range task.nearbyTickets {
		ticket := task.nearbyTickets[t]
		_, valid := task.validateTicket(ticket)
		if !valid {
			continue
		}
		fmt.Printf("Processing ticket %d\n", t)
		task.updatePossibleFields(ticket, possibleFields)
	}
	return possibleFields
}

func (task *Task) updatePossibleFields(ticket []int, possibleFields []uint) {
	for i := range ticket {
		for f := range possibleFields {
			if possibleFields[i]&(1<<f) != 0 && !task.fields[f].isValidValue(ticket[i]) {
				fmt.Printf("Field %d is not %s: unsuitable value %d\n", i, task.fields[f].name, ticket[i])
				possibleFields[i] &= not(1 << f)
				task.checkUniquePossibility(possibleFields, i)
			}
		}
	}
}

func (task *Task) checkUniquePossibility(possibleFields []uint, i int) {
	if bits.OnesCount(possibleFields[i]) == 1 {
		fmt.Printf("Found unique possibility for field %d\n", i)
		for c := range possibleFields {
			if c != i && bits.OnesCount(possibleFields[c]) != 1 {
				possibleFields[c] &= not(possibleFields[i])
				task.checkUniquePossibility(possibleFields, c)
			}
		}
	} else if possibleFields[i] == 0 {
		fmt.Printf("No valid possibilities for field %d\n", i)
	}
}

func (field *Field) isValidValue(n int) bool {
	for i := range field.ranges {
		if n >= field.ranges[i].min && n <= field.ranges[i].max {
			return true
		}
	}
	return false
}

func not(mask uint) uint {
	return math.MaxUint32 ^ mask
}

func indexOf(slice []uint, value uint) int {
	for i := range slice {
		if slice[i] == value {
			return i
		}
	}
	return -1
}
