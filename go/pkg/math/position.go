package math

import "fmt"

type Position struct {
	x, y int
}

func NewPosition(x, y int) Position {
	return Position{x, y}
}

func (pos Position) X() int {
	return pos.x
}

func (pos Position) Y() int {
	return pos.y
}

func (pos Position) WithX(x int) Position {
	return Position{x: x, y: pos.y}
}

func (pos Position) WithY(x int) Position {
	return Position{x: x, y: pos.y}
}

func (pos Position) Relative(x, y int) Position {
	return Position{x: pos.x + x, y: pos.y + y}
}

func (pos Position) InBounds(xMin, xMax, yMin, yMax int) bool {
	return pos.x >= xMin && pos.x <= xMax && pos.y >= yMin && pos.y <= yMax
}

func (pos Position) String() string {
	return fmt.Sprintf("{ x:%d, y:%d }", pos.x, pos.y)
}
