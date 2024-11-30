package aoc_core

type ProblemResult interface {
	String() string
}

type Problem interface {
	PartOne(input []byte) (any, error)
	PartTwo(input []byte) (any, error)
}
