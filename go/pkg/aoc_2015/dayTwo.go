package aoc2015

import (
	"bufio"
	"bytes"
	"strconv"
)

type DayTwo struct{}

// 2*l*w + 2*w*h + 2*h*l
func (day DayTwo) PartOne(input []byte) (any, error) {
	scanner := bufio.NewScanner(bytes.NewReader(input))

	sum := 0
	for scanner.Scan() {
		line := scanner.Bytes()
		if len(line) == 0 {
			continue
		}

		dim, err := parseLine(line)
		if err != nil {
			return -1, err
		}

		faceA := dim.length * dim.width
		faceB := dim.width * dim.height
		faceC := dim.height * dim.length
		minFace := min(faceA, faceB, faceC)
		area := 2*faceA + 2*faceB + 2*faceC

		sum += area + minFace
	}
	return sum, nil
}

func (day DayTwo) PartTwo(input []byte) (any, error) {
	scanner := bufio.NewScanner(bytes.NewReader(input))

	sum := 0
	for scanner.Scan() {
		line := scanner.Bytes()
		if len(line) == 0 {
			continue
		}
		dim, err := parseLine(line)
		if err != nil {
			return -1, err
		}

		volume := dim.length * dim.width * dim.height
		perimeterA := 2 * (dim.length + dim.width)
		perimeterB := 2 * (dim.width + dim.height)
		perimeterC := 2 * (dim.height + dim.length)
		minPerimeter := min(perimeterA, perimeterB, perimeterC)

		sum += volume + minPerimeter
	}
	return sum, nil
}

type Dim struct {
	length int
	width  int
	height int
}

func parseLine(line []byte) (Dim, error) {
	components := bytes.SplitN(line, []byte("x"), 3)
	length, err := strconv.Atoi(string(components[0]))
	if err != nil {
		return Dim{}, err
	}
	width, err := strconv.Atoi(string(components[1]))
	if err != nil {
		return Dim{}, err
	}
	height, err := strconv.Atoi(string(components[2]))
	if err != nil {
		return Dim{}, err
	}

	return Dim{length, width, height}, nil
}
