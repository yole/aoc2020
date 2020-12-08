package main

import (
	"aoc2020/shared"
	"fmt"
	"log"
	"strconv"
	"strings"
)

type BagCount struct {
	color string
	count int
}

func main() {
	fmt.Println(solveDay7Step1("day7input.txt"))
	fmt.Println(solveDay7Step2("day7input.txt"))
}

func solveDay7Step1(input string) int {
	containers := make(map[string][]string)
	shared.ProcessInputLines(input, func(line string) {
		holder, content := parseBagRule(line)
		for i := range content {
			bagCount := content[i]
			list, ok := containers[bagCount.color]
			if ok {
				containers[bagCount.color] = append(list, holder)
			} else {
				containers[bagCount.color] = []string{holder}
			}
		}
	})
	processed := make(map[string]bool)
	countContainers(containers, "shiny gold", processed)
	return len(processed)
}

func solveDay7Step2(input string) int {
	rulesets := make(map[string][]BagCount)
	shared.ProcessInputLines(input, func(line string) {
		holder, content := parseBagRule(line)
		rulesets[holder] = content
	})
	return countContents(rulesets, "shiny gold")
}

func parseBagRule(line string) (string, []BagCount) {
	bagIndex := strings.Index(line, "bag")
	holder := strings.Trim(line[0:bagIndex], " ")
	contentIndex := strings.Index(line, "contain") + 8
	contentString := strings.Split(line[contentIndex:], ", ")
	if contentString[0][0:2] == "no" {
		return holder, make([]BagCount, 0)
	}
	content := make([]BagCount, len(contentString))
	for i := range contentString {
		colorWithCount := strings.Split(contentString[i], " ")
		count, err := strconv.Atoi(colorWithCount[0])
		if err != nil {
			log.Fatal(err)
		}
		content[i] = BagCount{
			color: colorWithCount[1] + " " + colorWithCount[2],
			count: count,
		}
	}
	return holder, content
}

func countContainers(containers map[string][]string, color string, processed map[string]bool) {
	containerColors, ok := containers[color]
	if !ok {
		return
	}
	for i := range containerColors {
		containerColor := containerColors[i]
		_, ok := processed[containerColor]
		if !ok {
			processed[containerColor] = true
			countContainers(containers, containerColor, processed)
		}
	}
}

func countContents(rulesets map[string][]BagCount, color string) int {
	contents, ok := rulesets[color]
	if !ok {
		return 0
	}
	result := 0
	for i := range contents {
		result += (1 + countContents(rulesets, contents[i].color)) * contents[i].count
	}
	return result
}
