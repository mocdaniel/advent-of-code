package day01

import (
	"fmt"
	"strconv"
	"strings"
)

func findMinValue(source []int) (int, int) {
	index := 0
	min := source[0]

	for idx, number := range source {
		if number < min {
			index = idx
			min = number
		}
	}

	return min, index
}

func sumArray(source []int) int {
	sum := 0

	for _, number := range source {
		sum += number
	}

	return sum
}

func GetElfWithMostFood(input string) (int, error) {
	splitInventories := strings.Split(input, "\n\n")

	fmt.Printf("Found %v inventories\n\n", len(splitInventories))

	maxInventory := 0

	for _, inv := range splitInventories {
		currentInventory := 0
		items := strings.Split(inv, "\n")

		for _, item := range items {
			if item != "" {
				calories, err := strconv.ParseInt(item, 10, 64)
				if err != nil {
					return 0, err
				}

				currentInventory += int(calories)
			}
		}

		if currentInventory > maxInventory {
			maxInventory = currentInventory
		}
	}

	return maxInventory, nil
}

func GetThreeElvesWithMostFood(input string) (int, error) {
	splitInventories := strings.Split(input, "\n\n")

	fmt.Printf("Found %v inventories\n\n", len(splitInventories))

	maxInventories := []int{0, 0, 0}

	for _, inv := range splitInventories {
		currentInventory := 0
		items := strings.Split(inv, "\n")

		for _, item := range items {
			if item != "" {
				calories, err := strconv.ParseInt(item, 10, 64)
				if err != nil {
					return 0, err
				}

				currentInventory += int(calories)
			}
		}

		min, idx := findMinValue(maxInventories)
		if min < currentInventory {
			maxInventories[idx] = currentInventory
		}
	}

	return sumArray(maxInventories), nil
}
