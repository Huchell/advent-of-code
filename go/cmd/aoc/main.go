package main

import (
	"embed"
	"flag"
	"fmt"
	"io/fs"
	"log"

	aoc2015 "huchell/aoc/pkg/aoc_2015"
	"huchell/aoc/pkg/aoc_core"
)

type problemFunc = func(int) aoc_core.Problem

var runners = map[int]problemFunc{
	2015: aoc2015.FindProblem,
}

func main() {
	yearPtr := flag.Int("year", 2015, "Year to run")
	dayPtr := flag.Int("day", 1, "Day to run")

	runner, ok := runners[*yearPtr]
	if !ok {
		log.Fatalf("Could not find a year for %d", *yearPtr)
	}

	problem := runner(*dayPtr)
	if problem == nil {
		log.Fatalf("Could not find problem for day %d", *dayPtr)
	}

	input, err := input(*yearPtr, *dayPtr)
	if err != nil {
		log.Fatalf("Failed to load input: %+v", err)
	}

	partOneResult, err := problem.PartOne(input)
	if err != nil {
		log.Fatalf("[%d/%d] Part One failed: %+v", *yearPtr, *dayPtr, err)
	}
	log.Printf("[%d/%d] Part One result: %+v", *yearPtr, *dayPtr, partOneResult)

	partTwoResult, err := problem.PartTwo(input)
	if err != nil {
		log.Fatalf("[%d/%d] Part Two failed: %+v", *yearPtr, *dayPtr, err)
	}
	log.Printf("[%d/%d] Part Two result: %+v", *yearPtr, *dayPtr, partTwoResult)
}

//go:embed inputs
var inputFS embed.FS

func input(year, day int) ([]byte, error) {
	inputPath := fmt.Sprintf("inputs/%d/day%d.txt", year, day)
	return fs.ReadFile(inputFS, inputPath)
}
