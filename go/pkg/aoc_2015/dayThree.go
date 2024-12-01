package aoc2015

type DayThree struct{}

type position struct {
	x, y int
}

type visitMap = map[position]bool

func (day DayThree) PartOne(input []byte) (any, error) {
	housesVisited := make(visitMap, 1)
	pos := position{x: 0, y: 0}
	housesVisited[pos] = true

	for _, ch := range input {
		pos = movePosition(ch, pos)
		if housesVisited[pos] {
			continue
		}
		housesVisited[pos] = true
	}

	return len(housesVisited), nil
}

func (day DayThree) PartTwo(input []byte) (any, error) {
	housesVisited := []visitMap{
		map[position]bool{{x: 0, y: 0}: true},
		map[position]bool{{x: 0, y: 0}: true},
	}

	positions := []position{
		{x: 0, y: 0},
		{x: 0, y: 0},
	}
	for idx, ch := range input {
		mover := idx % 2
		visitMap := housesVisited[mover]
		pos := positions[mover]

		pos = movePosition(ch, pos)
		positions[mover] = pos

		visitMap[pos] = true
	}

	finalHousesVisited := make(map[position]bool)
	for _, visitMap := range housesVisited {
		for key := range visitMap {
			finalHousesVisited[key] = true
		}
	}

	return len(finalHousesVisited), nil
}

func movePosition(ch byte, pos position) position {
	switch ch {
	case '^':
		return position{x: pos.x, y: pos.y + 1}
	case '>':
		return position{x: pos.x + 1, y: pos.y}
	case 'v':
		return position{x: pos.x, y: pos.y - 1}
	case '<':
		return position{x: pos.x - 1, y: pos.y}
	}
	return pos
}
