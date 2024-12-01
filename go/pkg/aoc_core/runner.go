package aoc_core

import (
	"errors"
	"fmt"
	"log"
)

type Runner struct {
	problems []Problem
}

func NewRunner(problems ...Problem) Runner {
	return Runner{
		problems,
	}
}

func (runner *Runner) ExecuteAll(input [][]byte) error {
	for dayIdx, problem := range runner.problems {
		err := executeProblem(input[dayIdx], problem)
		if err != nil {
			return err
		}
	}
	return nil
}

func (runner *Runner) ExecuteProblem(input []byte, day int) error {
	if len(runner.problems) <= day {
		return errors.New(fmt.Sprintf("DayToLarge: %d", day))
	}

	problem := runner.problems[day]
	if problem == nil {
		return errors.New(fmt.Sprintf("DayNotFound: %d", day))
	}
	executeProblem(input, problem)
	return nil
}

func executeProblem(input []byte, problem Problem) error {
	result, err := problem.PartOne(input)
	if err != nil {
		return err
	}

	log.Printf("Part One result: %+v", result)

	result, err = problem.PartTwo(input)
	if err != nil {
		return err
	}

	log.Printf("Part Two result: %+v", result)
	return nil
}
