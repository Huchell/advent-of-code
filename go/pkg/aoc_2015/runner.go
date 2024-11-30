package aoc2015

import "huchell/aoc/pkg/aoc_core"

func FindProblem(day int) aoc_core.Problem {
	switch day {
	case 1:
		return DayOne{}
	default:
		return nil
	}
}
