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
	fmt.Println(solveDay8Step2("day08input.txt"))
}

func solveDay8Step1(input string) int {
	program := loadProgram(input)
	_, acc := checkTerminates(program, 0, 0)
	return acc
}

func solveDay8Step2(input string) int {
	program := loadProgram(input)
	accSequence := make([]int, len(program))
	ipSequence := make([]int, len(program))
	visited := make([]bool, len(program))
	ip := 0
	acc := 0
	step := 0
	for {
		if visited[ip] {
			break
		}
		visited[ip] = true
		acc, ip = executeInstruction(program, ip, acc)
		accSequence[step] = acc
		ipSequence[step] = ip
		step++
	}

	for i := step - 1; i >= 0; i-- {
		if program[ipSequence[i]].opcode != ACC {
			terminates, acc := checkTerminates(patchProgram(program, ipSequence[i]), accSequence[i], ipSequence[i])
			if terminates {
				return acc
			}
		}
	}

	return -1
}

func loadProgram(input string) []instruction {
	program := make([]instruction, 0)
	shared.ProcessInputLines(input, func(line string) {
		program = append(program, parseInstruction(line))
	})
	return program
}

func checkTerminates(program []instruction, acc int, ip int) (bool, int) {
	visited := make([]bool, len(program))
	for {
		if ip == len(program) {
			return true, acc
		}
		if visited[ip] {
			return false, acc
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

func patchProgram(program []instruction, offset int) []instruction {
	patchedProgram := make([]instruction, len(program))
	copy(patchedProgram, program)
	if patchedProgram[offset].opcode == NOP {
		patchedProgram[offset].opcode = JMP
	} else {
		patchedProgram[offset].opcode = NOP
	}
	return patchedProgram
}
