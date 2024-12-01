package aoc2024

import "testing"

func TestDayOnePartOne(t *testing.T) {
	input := []byte(`3   4
4   3
2   5
1   3
3   9
3   3
`)
	result, err := DayOne{}.PartOne(input)
	if err != nil {
		t.Fatal(err)
	}

	if result != 11 {
		t.Fatalf("%d != 11", result)
	}
}

func TestDayOnePartTwo(t *testing.T) {
	input := []byte(`3   4
4   3
2   5
1   3
3   9
3   3
`)

	result, err := DayOne{}.PartTwo(input)
	if err != nil {
		t.Fatal(err)
	}

	if result != 31 {
		t.Fatalf("%d != 31", result)
	}
}
