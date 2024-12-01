package aoc2024

import (
	"bufio"
	"bytes"
	"slices"
	"strconv"
)

type DayOne struct{}

func (DayOne) PartOne(input []byte) (any, error) {
	lhs, rhs, err := getLists(input)
	if err != nil {
		return -1, err
	}

	slices.Sort(lhs)
	slices.Sort(rhs)

	sum := 0
	for idx := 0; idx < len(lhs); idx++ {
		lhsVal := lhs[idx]
		rhsVal := rhs[idx]
		sum += max(lhsVal, rhsVal) - min(lhsVal, rhsVal)
	}

	return sum, nil
}

func (DayOne) PartTwo(input []byte) (any, error) {
	lhs, rhs, err := getLists(input)
	if err != nil {
		return -1, err
	}

	totalScore := 0
	scoreLookup := make(map[int]int, 0)
	for _, lhsVal := range lhs {
		score, ok := scoreLookup[lhsVal]
		if !ok {
			score = count(rhs, lhsVal)
			scoreLookup[lhsVal] = score
		}
		totalScore += score * lhsVal
	}

	return totalScore, nil
}

func getLists(input []byte) ([]int, []int, error) {
	scanner := bufio.NewScanner(bytes.NewReader(input))

	lhs := []int{}
	rhs := []int{}
	for scanner.Scan() {
		line := scanner.Bytes()
		if len(line) == 0 {
			continue
		}

		split := bytes.Split(line, []byte("   "))
		lhsVal, err := strconv.Atoi(string(split[0]))
		if err != nil {
			return nil, nil, err
		}
		rhsVal, err := strconv.Atoi(string(split[1]))
		if err != nil {
			return nil, nil, err
		}
		lhs = append(lhs, lhsVal)
		rhs = append(rhs, rhsVal)
	}

	return lhs, rhs, nil
}

func count(arr []int, val int) int {
	total := 0
	for _, arrVal := range arr {
		if arrVal != val {
			continue
		}

		total += 1
	}
	return total
}
