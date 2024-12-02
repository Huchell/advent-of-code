package aoc2024

import (
	"bufio"
	"bytes"
	"math"
	"strconv"
)

type DayTwo struct{}

func (DayTwo) PartOne(input []byte) (any, error) {
	scanner := bufio.NewScanner(bytes.NewReader(input))

	safeCount := 0
	for scanner.Scan() {
		line := scanner.Bytes()
		if len(line) == 0 {
			continue
		}

		levelsBytes := bytes.Split(line, []byte(" "))
		levels, err := charsToIntArray(levelsBytes)
		if err != nil {
			return -1, err
		}
		if !isSafe(levels) {
			continue
		}

		safeCount += 1
	}
	return safeCount, nil
}

func (DayTwo) PartTwo(input []byte) (any, error) {
	scanner := bufio.NewScanner(bytes.NewReader(input))

	safeCount := 0
	for scanner.Scan() {
		line := scanner.Bytes()
		if len(line) == 0 {
			continue
		}

		levelsBytes := bytes.Split(line, []byte(" "))
		levels, err := charsToIntArray(levelsBytes)
		if err != nil {
			return -1, err
		}
		if !problemDampenedIsSafe(levels) {
			continue
		}

		safeCount += 1
	}
	return safeCount, nil
}

func isSafe(levels []int) bool {
	isDecsending := math.Signbit(float64(levels[1] - levels[0]))
	lastVal := levels[0]
	for idx := 1; idx < len(levels); idx++ {
		val := levels[idx]
		diff := float64(val - lastVal)
		lastVal = val

		if diff == 0 {
			return false
		}
		if diff < -3 {
			return false
		}
		if diff > 3 {
			return false
		}
		if math.Signbit(diff) != isDecsending {
			return false
		}
	}
	return true
}

func problemDampenedIsSafe(levels []int) bool {
	if isSafe(levels) {
		return true
	}

	for idx := range levels {
		dampenedLevels := []int{}
		dampenedLevels = append(dampenedLevels, levels[:idx]...)
		dampenedLevels = append(dampenedLevels, levels[idx+1:]...)
		if isSafe(dampenedLevels) {
			return true
		}
	}
	return false
}

func charsToIntArray(src [][]byte) ([]int, error) {
	dst := make([]int, len(src))
	for idx, str := range src {
		num, err := strconv.Atoi(string(str))
		if err != nil {
			return nil, err
		}
		dst[idx] = num
	}
	return dst, nil
}
