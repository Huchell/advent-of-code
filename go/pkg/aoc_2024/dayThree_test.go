package aoc2024

import "testing"

func TestDayThreePartOne(t *testing.T) {
	input := []byte(`xmul(2,4)%&mul[3,7]!@^do_not_mul(5,5)+mul(32,64]then(mul(11,8)mul(8,5))`)

	result, err := DayThree{}.PartOne(input)
	if err != nil {
		t.Fatal(err)
	}

	if result != 161 {
		t.Fatalf("%d != 161", result)
	}
}

func TestDayThreePartTwo(t *testing.T) {
	input := []byte(`xmul(2,4)&mul[3,7]!^don't()_mul(5,5)+mul(32,64](mul(11,8)undo()?mul(8,5))`)

	result, err := DayThree{}.PartTwo(input)
	if err != nil {
		t.Fatal(err)
	}

	if result != 48 {
		t.Fatalf("%d != 48", result)
	}
}
