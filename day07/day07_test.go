package main

import "testing"

func TestParseBagLine(t *testing.T) {
	holder, content := parseBagRule("light red bags contain 1 bright white bag, 2 muted yellow bags.")
	if holder != "light red" {
		t.Errorf("Holder incorrect (expected light red): %s", holder)
	}
	if len(content) != 2 {
		t.Errorf("unexpected length of content")
	}
	if content[0].count != 1 {
		t.Errorf("unexpected count[0]")
	}
	if content[0].color != "bright white" {
		t.Errorf("unexpected color[0]")
	}
	if content[1].count != 2 {
		t.Errorf("unexpected count[1]")
	}
	if content[1].color != "muted yellow" {
		t.Errorf("unexpected color[1]")
	}
}

func TestParseBagLine7Empty(t *testing.T) {
	holder, content := parseBagRule("faded blue bags contain no other bags.")
	if holder != "faded blue" {
		t.Errorf("Holder incorrect (expected light red): %s", holder)
	}
	if len(content) != 0 {
		t.Errorf("unexpected length of content")
	}
}

func TestSolveDay7Step1(t *testing.T) {
	result := solveDay7Step1("day7testinput.txt")
	if result != 4 {
		t.Errorf("Unexpected result %d", result)
	}
}

func TestSolveDay7Step2(t *testing.T) {
	result := solveDay7Step2("day7testinput.txt")
	if result != 32 {
		t.Errorf("Unexpected result %d", result)
	}
}
