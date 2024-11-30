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

var runners = map[int]aoc_core.Runner{
	2015: aoc2015.NewRunner,
}

func main() {
	yearPtr := flag.Int("year", 2015, "Year to run")
	dayPtr := flag.Int("day", -1, "Day to run")

	runner, ok := runners[*yearPtr]
	if !ok {
		log.Fatalf("Could not find a year for %d", *yearPtr)
	}

	if *dayPtr == -1 {
		inputs, err := getInputs(*yearPtr)
		if err != nil {
			log.Fatalf("Could not load input %+v", err)
		}
		log.Fatal(runner.ExecuteAll(inputs))
	}

	input, err := getInput(*yearPtr, *dayPtr)
	if err != nil {
		log.Fatalf("Could not load input %+v", err)
	}
	log.Fatal(runner.ExecuteProblem(input, *dayPtr))
}

//go:embed inputs
var inputFS embed.FS

func getInput(year, day int) ([]byte, error) {
	inputPath := fmt.Sprintf("inputs/%d/day%d.txt", year, day)
	return fs.ReadFile(inputFS, inputPath)
}

func getInputs(year int) ([][]byte, error) {
	inputPath := fmt.Sprintf("inputs/%d", year)
	entries, err := fs.ReadDir(inputFS, inputPath)
	if err != nil {
		return nil, err
	}

	results := make([][]byte, len(entries))
	for idx, entry := range entries {
		dayInputPath := fmt.Sprintf("%s/%s", inputPath, entry.Name())
		result, err := fs.ReadFile(inputFS, dayInputPath)
		if err != nil {
			return nil, err
		}
		results[idx] = result
	}
	return results, nil
}
