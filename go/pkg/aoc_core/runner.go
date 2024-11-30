package aoc_core

import (
	"errors"
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
	problem := runner.problems[day]
	if problem != nil {
		return errors.New("DayNotFound")
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
	return nil
}
