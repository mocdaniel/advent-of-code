package day09

import (
	"math"
	"strconv"
	"strings"
)

type Coordinate struct {
	x int
	y int
}

func (c1 *Coordinate) isAdjacentTo(c2 *Coordinate) bool {
	// Overlap
	if c1.equals(c2) {
		return true
	}

	// above/below each other
	if c1.x == c2.x && math.Abs(float64(c1.y-c2.y)) < 2 {
		return true
	}

	// right/left from each other
	if c1.y == c2.y && math.Abs(float64(c1.x-c2.x)) < 2 {
		return true
	}

	// diagonally adjacent
	if math.Abs(float64(c1.x-c2.x)) == 1 && math.Abs(float64(c1.y-c2.y)) == 1 {
		return true
	}

	return false
}

func (c1 *Coordinate) followHead(c2 *Coordinate) {
	// In case of overlap, do nothing
	if c1.x == c2.x && c1.y == c2.y {
		return
	}

	if c2.x > c1.x {
		c1.x += 1
	} else if c2.x < c1.x {
		c1.x -= 1
	}

	if c2.y > c1.y {
		c1.y += 1
	} else if c2.y < c1.y {
		c1.y -= 1
	}
}

func (c1 *Coordinate) equals(c2 *Coordinate) bool {
	if c1.x == c2.x && c1.y == c2.y {
		return true
	}
	return false
}

// up: 0
// down: 1
// left: 2
// right: 3
func (c *Coordinate) move(axis int) {
	switch axis {
	case 0:
		c.y -= 1
	case 1:
		c.y += 1
	case 2:
		c.x -= 1
	case 3:
		c.x += 1
	}
}

type CoordinateList []Coordinate

func newCoordList() CoordinateList {
	return make([]Coordinate, 0)
}

func (cs CoordinateList) contains(c *Coordinate) bool {
	for i := range cs {
		if cs[i].equals(c) {
			return true
		}
	}
	return false
}

func GetVisitedPositions(input string, elements int) (int, error) {
	alreadyVisited := newCoordList()

	head := Coordinate{0, 0}
	tails := make([]Coordinate, elements-1)
	for i := range tails {
		tails[i] = Coordinate{0, 0}
	}

	instructions := strings.Split(strings.TrimSuffix(input, "\n"), "\n")

	for i := range instructions {
		info := strings.Split(instructions[i], " ")
		steps, err := strconv.ParseInt(info[1], 10, 64)
		if err != nil {
			return 0, err
		}
		for j := 0; j < int(steps); j++ {
			switch info[0] { // move head first
			case "U":
				head.move(0)
			case "D":
				head.move(1)
			case "L":
				head.move(2)
			case "R":
				head.move(3)
			}

			movingPart := head
			for k := 0; k < elements-1; k++ {
				if !tails[k].isAdjacentTo(&movingPart) {
					tails[k].followHead(&movingPart)
					movingPart = tails[k]
				} else {
					break
				}
			}

			if !alreadyVisited.contains(&tails[elements-2]) {
				alreadyVisited = append(alreadyVisited, Coordinate{tails[elements-2].x, tails[elements-2].y}) // if tail is on an unregistered coordinate, add it
			}
		}
	}

	return len(alreadyVisited), nil
}
