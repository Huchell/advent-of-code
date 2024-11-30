package aoc2015

import "errors"

type DayOne struct{}

func (day DayOne) PartOne(input []byte) (any, error) {
	sum := 0
	for _, ch := range input {
		switch ch {
		case '(':
			sum += 1
			break

		case ')':
			sum -= 1
			break
		}
	}
	return sum, nil
}

func (day DayOne) PartTwo(input []byte) (any, error) {
	sum := 0
	for idx, ch := range input {
		switch ch {
		case '(':
			sum += 1
			break

		case ')':
			sum -= 1
			break
		}

		if sum < 0 {
			return idx + 1, nil
		}
	}
	return -1, errors.New("NoIndexFound")
}
