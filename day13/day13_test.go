package main

import (
	"testing"
)

func TestSolveDay13Step1(t *testing.T) {
	distance := solveDay13Step1("day13testinput.txt")
	if distance != 295 {
		t.Errorf("Unexpected result %d", distance)
	}
}

func TestBezoutCoefficients(t *testing.T) {
	m1, m2 := bezoutCoefficients(3, 4)
	if m1 != -1 || m2 != 1 {
		t.Errorf("Unexpected result (%d, %d)", m1, m2)
	}
}

func TestChineseRemainder(t *testing.T) {
	result := chineseRemainder(0, 3, 3, 4)
	if result != 3 {
		t.Errorf("Unexpected result %d", result)
	}
}

func TestMultipleChineseRemainder(t *testing.T) {
	result := multipleChineseRemainder([]int64{3, 4, 5}, []int64{0, 3, 4})
	if result != 39 {
		t.Errorf("Unexpected result %d", result)
	}
}

func TestMultipleChineseRemainderExample1(t *testing.T) {
	result := multipleChineseRemainder([]int64{17, 13, 19}, []int64{0, 13 - 2, 19 - 3})
	if result != 3417 {
		t.Errorf("Unexpected result %d", result)
	}
}

func TestSolveDay13Step2(t *testing.T) {
	result := multipleChineseRemainderFromString("17,x,13,19")
	if result != 3417 {
		t.Errorf("Unexpected result %d", result)
	}
}

func TestSolveDay13Step2_2(t *testing.T) {
	result := multipleChineseRemainderFromString("7,13,x,x,59,x,31,19")
	if result != 1068781 {
		t.Errorf("Unexpected result %d", result)
	}
}

func TestSolveDay13Step2_3(t *testing.T) {
	result := multipleChineseRemainderFromString("67,7,59,61")
	if result != 754018 {
		t.Errorf("Unexpected result %d", result)
	}
}

func TestSolveDay13Step2_4(t *testing.T) {
	result := multipleChineseRemainderFromString("1789,37,47,1889")
	if result != 1202161486 {
		t.Errorf("Unexpected result %d", result)
	}
}

func TestSolveDay13Step2_5(t *testing.T) {
	result := multipleChineseRemainderFromString("67,x,7,59,61")
	if result != 779210 {
		t.Errorf("Unexpected result %d", result)
	}
}

func TestSolveDay13Step2_6(t *testing.T) {
	result := multipleChineseRemainderFromString("67,7,x,59,61")
	if result != 1261476 {
		t.Errorf("Unexpected result %d", result)
	}
}
