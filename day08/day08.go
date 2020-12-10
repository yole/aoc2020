package main

import (
	"aoc2020/shared"
	"fmt"
	"log"
	"strconv"
	"strings"
)

const (
	NOP = iota
	ACC
	JMP
)

type instruction struct {
	opcode  int
	operand int
}

func main() {
	fmt.Println(solveDay8Step1("day08input.txt"))
}

func solveDay8Step1(input string) int {
	program := make([]instruction, 0)
	shared.ProcessInputLines(input, func(line string) {
		program = append(program, parseInstruction(line))
	})
	visited := make([]bool, len(program))
	acc := 0
	ip := 0
	for {
		if visited[ip] {
			return acc
		}
		visited[ip] = true
		acc, ip = executeInstruction(program, ip, acc)
	}
}

func executeInstruction(program []instruction, ip int, acc int) (int, int) {
	switch program[ip].opcode {
	case NOP:
		ip++
		break
	case ACC:
		acc += program[ip].operand
		ip++
		break
	case JMP:
		ip += program[ip].operand
		break
	}
	return acc, ip
}

func parseInstruction(line string) instruction {
	parts := strings.Split(line, " ")
	opcode := -1
	switch parts[0] {
	case "nop":
		opcode = NOP
		break
	case "acc":
		opcode = ACC
		break
	case "jmp":
		opcode = JMP
		break
	}
	if opcode == -1 {
		log.Fatal("Unknown opcode " + parts[0])
	}
	operand, err := strconv.Atoi(parts[1])
	if err != nil {
		log.Fatal("Can't parse operand " + parts[1])
	}
	return instruction{
		opcode:  opcode,
		operand: operand,
	}
}
