package aoc2024

import (
	"huchell/aoc/pkg/assert"
	"testing"
)

func TestDayFivePartOne(t *testing.T) {
	input := []byte(`47|53
97|13
97|61
97|47
75|29
61|13
75|53
29|13
97|29
53|29
61|53
97|53
61|29
47|13
75|47
97|75
47|61
75|61
47|29
75|13
53|13

75,47,61,53,29
97,61,53,29,13
75,29,13
75,97,47,61,53
61,13,29
97,13,75,29,47`)

	result, err := DayFive{}.PartOne(input)
	assert.TestErrorIsNil(t, err)

	assert.TestIsEqual(t, result, 143)
}

func TestDayFivePartTwo(t *testing.T) {
	input := []byte(`47|53
97|13
97|61
97|47
75|29
61|13
75|53
29|13
97|29
53|29
61|53
97|53
61|29
47|13
75|47
97|75
47|61
75|61
47|29
75|13
53|13

75,47,61,53,29
97,61,53,29,13
75,29,13
75,97,47,61,53
61,13,29
97,13,75,29,47`)

	result, err := DayFive{}.PartTwo(input)
	assert.TestErrorIsNil(t, err)

	assert.TestIsEqual(t, result, 123)
}
