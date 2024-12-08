package aoc2024

import (
	"bufio"
	"bytes"
	"slices"

	"huchell/aoc/pkg/math"
)

type DayFour struct{}

func (DayFour) PartOne(input []byte) (any, error) {
	grid := parseInput(input)

	instances := 0
	for rowIdx := 0; rowIdx < grid.height; rowIdx++ {
		for colIdx := 0; colIdx < grid.width; colIdx++ {
			pos := math.NewPosition(colIdx, rowIdx)
			ch := grid.Cell(pos)
			switch ch {
			case 'X':
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
			pos := math.NewPosition(colIdx, rowIdx)
			ch := grid.Cell(pos)
			switch ch {
			case 'A':
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

func checkForX_MAS(grid DayFourGrid, pos math.Position) bool {
	positions := []math.Position{
		pos.Relative(-1, -1),
		pos.Relative(1, -1),
		pos,
		pos.Relative(-1, 1),
		pos.Relative(1, 1),
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

func (grid DayFourGrid) InstancesOf(pos math.Position, str []byte) int {
	instances := 0
	if grid.checkInstanceOf(pos, str, math.NewPosition(-1, 0)) {
		instances += 1
	}
	if grid.checkInstanceOf(pos, str, math.NewPosition(-1, -1)) {
		instances += 1
	}
	if grid.checkInstanceOf(pos, str, math.NewPosition(0, -1)) {
		instances += 1
	}
	if grid.checkInstanceOf(pos, str, math.NewPosition(1, -1)) {
		instances += 1
	}
	if grid.checkInstanceOf(pos, str, math.NewPosition(1, 0)) {
		instances += 1
	}
	if grid.checkInstanceOf(pos, str, math.NewPosition(1, 1)) {
		instances += 1
	}
	if grid.checkInstanceOf(pos, str, math.NewPosition(0, 1)) {
		instances += 1
	}
	if grid.checkInstanceOf(pos, str, math.NewPosition(-1, 1)) {
		instances += 1
	}

	return instances
}

func (grid DayFourGrid) checkInstanceOf(pos math.Position, str []byte, dir math.Position) bool {
	strLen := len(str)

	endPos := pos.Relative((strLen-1)*dir.X(), (strLen-1)*dir.Y())
	if !endPos.InBounds(0, grid.width-1, 0, grid.height-1) {
		return false
	}

	slice := make([]byte, strLen)
	for idx := 0; idx < strLen; idx++ {
		cellPos := pos.Relative(idx*dir.X(), idx*dir.Y())
		cell := grid.Cell(cellPos)
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

func (grid DayFourGrid) InBoundsSlice(positions []math.Position) bool {
	for _, pos := range positions {
		if !pos.InBounds(0, grid.width-1, 0, grid.height-1) {
			return false
		}
	}
	return true
}

func (grid DayFourGrid) Cell(pos math.Position) byte {
	return grid.cells[pos.Y()][pos.X()]
}
