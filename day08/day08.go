package day08

import (
	"strconv"
	"strings"
)

func ParseGrid(input string) ([][]int, error) {
	rows := strings.Split(strings.TrimSuffix(input, "\n"), "\n")
	length := len(rows)
	grid := make([][]int, length)
	for i := range grid {
		grid[i] = make([]int, length)
	}

	for i, r := range rows {
		for j, c := range r {
			n, err := strconv.ParseInt(string(c), 10, 64)
			if err != nil {
				return nil, err
			}
			grid[i][j] = int(n)
		}
	}

	return grid, nil
}

func CountTrees(grid [][]int) int {
	totalCount := 0
	length := len(grid)

	countedTrees := make([][]bool, length)
	for i := range countedTrees {
		countedTrees[i] = make([]bool, length)
	}

	totalCount += fromNorth(grid, countedTrees)
	totalCount += fromEast(grid, countedTrees)
	totalCount += fromSouth(grid, countedTrees)
	totalCount += fromWest(grid, countedTrees)

	return totalCount
}

func fromNorth(grid [][]int, countedTrees [][]bool) int {
	length := len(grid)
	count := 00

	for c := 0; c < length; c++ {
		prev := -1
		for r := 0; r < length; r++ {
			if grid[r][c] > prev {
				prev = grid[r][c]
				if !countedTrees[r][c] {
					countedTrees[r][c] = true
					count += 1
				}
				if grid[r][c] == 9 {
					break
				}
			}
		}
	}
	return count
}

func fromEast(grid [][]int, countedTrees [][]bool) int {
	length := len(grid)
	count := 0

	for r := 0; r < length; r++ {
		prev := -1
		for c := length - 1; c > -1; c-- {
			if grid[r][c] > prev {
				prev = grid[r][c]
				if !countedTrees[r][c] {
					countedTrees[r][c] = true
					count += 1
				}
				if grid[r][c] == 9 {
					break
				}
			}
		}
	}
	return count
}

func fromSouth(grid [][]int, countedTrees [][]bool) int {
	length := len(grid)
	count := 0

	for c := 0; c < length; c++ {
		prev := -1
		for r := length - 1; r > -1; r-- {
			if grid[r][c] > prev {
				prev = grid[r][c]
				if !countedTrees[r][c] {
					countedTrees[r][c] = true
					count += 1
				}
				if grid[r][c] == 9 {
					break
				}
			}
		}
	}
	return count
}

func fromWest(grid [][]int, countedTrees [][]bool) int {
	length := len(grid)
	count := 0

	for r := 0; r < length; r++ {
		prev := -1
		for c := 0; c < length; c++ {
			if grid[r][c] > prev {
				prev = grid[r][c]
				if !countedTrees[r][c] {
					countedTrees[r][c] = true
					count += 1
				}
				if grid[r][c] == 9 {
					break
				}
			}
		}
	}
	return count
}

func HighestScenicScore(grid [][]int) int {
	highestScore := 0

	for r := range grid {
		for c := range grid[r] {

			// fromNorth
			fromNorth := viewFromNorth(grid, r, c)
			fromSouth := viewFromSouth(grid, r, c)
			fromEast := viewFromEast(grid, r, c)
			fromWest := viewFromWest(grid, r, c)
			total := fromNorth * fromEast * fromSouth * fromWest
			if total > highestScore {
				highestScore = total
			}
		}
	}

	return highestScore
}

func viewFromNorth(grid [][]int, r int, c int) int {
	totalCount := 0

	if r == len(grid)-1 {
		return 0
	}

	for i := r + 1; i < len(grid); i++ {
		totalCount += 1
		if grid[i][c] >= grid[r][c] {
			return totalCount
		}
	}

	return totalCount
}

func viewFromEast(grid [][]int, r int, c int) int {
	totalCount := 0

	if c == 0 {
		return 0
	}

	for i := c - 1; i > -1; i-- {
		totalCount += 1
		if grid[r][i] >= grid[r][c] {
			return totalCount
		}
	}

	return totalCount
}

func viewFromSouth(grid [][]int, r int, c int) int {
	totalCount := 0

	if r == 0 {
		return 0
	}

	for i := r - 1; i > -1; i-- {
		totalCount += 1
		if grid[i][c] >= grid[r][c] {
			return totalCount
		}
	}

	return totalCount
}

func viewFromWest(grid [][]int, r int, c int) int {
	totalCount := 0

	if c == len(grid)-1 {
		return 0
	}

	for i := c + 1; i < len(grid); i++ {
		totalCount += 1
		if grid[r][i] >= grid[r][c] {
			return totalCount
		}
	}

	return totalCount
}
