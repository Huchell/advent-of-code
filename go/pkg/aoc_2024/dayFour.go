package aoc2024

import (
	"bufio"
	"bytes"
	"fmt"
	"slices"
)

type DayFour struct{}

func (DayFour) PartOne(input []byte) (any, error) {
	grid := parseInput(input)

	instances := 0
	for rowIdx := 0; rowIdx < grid.height; rowIdx++ {
		for colIdx := 0; colIdx < grid.width; colIdx++ {
			ch := grid.cells[rowIdx][colIdx]
			switch ch {
			case 'X':
				pos := Position{X: colIdx, Y: rowIdx}
				cellInstances := grid.InstancesOf(pos, []byte("XMAS"))
				instances += cellInstances
				break

			default:
				continue
			}
		}
	}
	return instances, nil
}

func (DayFour) PartTwo(input []byte) (any, error) {
	grid := parseInput(input)

	instances := 0
	for rowIdx := 0; rowIdx < grid.height; rowIdx++ {
		for colIdx := 0; colIdx < grid.width; colIdx++ {
			ch := grid.cells[rowIdx][colIdx]
			switch ch {
			case 'A':
				pos := Position{X: colIdx, Y: rowIdx}
				if checkForX_MAS(grid, pos) {
					instances += 1
				}
				break

			default:
				continue
			}
		}
	}
	return instances, nil
}

func checkForX_MAS(grid DayFourGrid, pos Position) bool {
	positions := []Position{
		{X: pos.X - 1, Y: pos.Y - 1},
		{X: pos.X + 1, Y: pos.Y - 1},
		pos,
		{X: pos.X - 1, Y: pos.Y + 1},
		{X: pos.X + 1, Y: pos.Y + 1},
	}

	if !grid.InBoundsSlice(positions) {
		return false
	}

	p1 := []byte{grid.Cell(positions[0]), grid.Cell(positions[2]), grid.Cell(positions[4])}
	p2 := []byte{grid.Cell(positions[1]), grid.Cell(positions[2]), grid.Cell(positions[3])}
	if !slices.Equal(p1, []byte("MAS")) && !slices.Equal(p1, []byte("SAM")) {
		return false
	}
	if !slices.Equal(p2, []byte("MAS")) && !slices.Equal(p2, []byte("SAM")) {
		return false
	}
	return true
}

func parseInput(input []byte) DayFourGrid {
	scanner := bufio.NewScanner(bytes.NewReader(input))

	grid := [][]byte{}
	width := 0
	for scanner.Scan() {
		line := scanner.Bytes()
		if len(line) == 0 {
			continue
		}

		runes := parseLine(line)
		width = len(line)
		grid = append(grid, runes)
	}
	return DayFourGrid{
		width:  width,
		height: len(grid),
		cells:  grid,
	}
}

func parseLine(line []byte) []byte {
	result := make([]byte, len(line))
	for idx, r := range line {
		result[idx] = r
	}
	return result
}

type DayFourGrid struct {
	width  int
	height int
	cells  [][]byte
}

type Position struct {
	X int
	Y int
}

func (pos Position) String() string {
	return fmt.Sprintf("{ X:%d, Y:%d }", pos.X, pos.Y)
}

func (grid DayFourGrid) InstancesOf(pos Position, str []byte) int {
	instances := 0
	if grid.checkInstanceOf(pos, str, Position{X: -1, Y: 0}) {
		instances += 1
	}
	if grid.checkInstanceOf(pos, str, Position{X: -1, Y: -1}) {
		instances += 1
	}
	if grid.checkInstanceOf(pos, str, Position{X: 0, Y: -1}) {
		instances += 1
	}
	if grid.checkInstanceOf(pos, str, Position{X: 1, Y: -1}) {
		instances += 1
	}
	if grid.checkInstanceOf(pos, str, Position{X: 1, Y: 0}) {
		instances += 1
	}
	if grid.checkInstanceOf(pos, str, Position{X: 1, Y: 1}) {
		instances += 1
	}
	if grid.checkInstanceOf(pos, str, Position{X: 0, Y: 1}) {
		instances += 1
	}
	if grid.checkInstanceOf(pos, str, Position{X: -1, Y: 1}) {
		instances += 1
	}

	return instances
}

func (grid DayFourGrid) checkInstanceOf(pos Position, str []byte, dir Position) bool {
	strLen := len(str)

	endX := pos.X + (strLen-1)*dir.X
	if endX < 0 || endX >= grid.width {
		return false
	}
	endY := pos.Y + (strLen-1)*dir.Y
	if endY < 0 || endY >= grid.height {
		return false
	}

	slice := make([]byte, strLen)
	for idx := 0; idx < strLen; idx++ {
		cell := grid.cells[pos.Y+idx*dir.Y][pos.X+idx*dir.X]
		slice[idx] = cell
	}

	if slices.Equal(slice, str) {
		return true
	}

	slices.Reverse(slice)
	if slices.Equal(slice, str) {
		return true
	}
	return false
}

func (grid DayFourGrid) InBoundsSlice(positions []Position) bool {
	for _, pos := range positions {
		if pos.X < 0 || pos.X >= grid.width {
			return false
		}
		if pos.Y < 0 || pos.Y >= grid.height {
			return false
		}
	}
	return true
}

func (grid DayFourGrid) Cell(pos Position) byte {
	return grid.cells[pos.Y][pos.X]
}
