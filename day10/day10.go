package day10

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

func GetSignalStrength(input []string, cycle int, shallDraw bool) (int, error) {
	register := 1
	currentCycle := 0
	for _, instruction := range input {
		operation := strings.Split(instruction, " ")
		for _, op := range operation {
			currentCycle += 1
			if currentCycle == cycle {
				return cycle * register, nil
			}

			if shallDraw {
				if math.Abs(float64(((currentCycle-1)%40)-register)) < 2 {
					draw(true, currentCycle)
				} else {
					draw(false, currentCycle)
				}
			}

			if !(op == "noop" || op == "addx") {
				amount, err := strconv.ParseInt(op, 10, 64)
				if err != nil {
					return 0, err
				}
				register += int(amount)
			}
		}
	}

	return 0, nil
}

func draw(lit bool, cycle int) {
	if lit {
		fmt.Print("#")
	} else {
		fmt.Print(".")
	}

	if cycle%40 == 0 {
		fmt.Print("\n")
	}
}
