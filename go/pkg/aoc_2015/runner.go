package aoc2015

import "huchell/aoc/pkg/aoc_core"

func NewRunner() aoc_core.Runner {
	return aoc_core.NewRunner(
		DayOne{},
	)
}
