package day05

import (
	"regexp"
	"strconv"
	"strings"
)

type Heap []string

func (h Heap) Put(s string) Heap {
	return append(h, s)
}

func (h Heap) PutStack(s []string) Heap {
	for _, i := range s {
		h = h.Put(i)
	}
	return h
}

// This pop is based on the assumption, that the riddle will never try to take a container off an empty stack.
func (h Heap) Pop() (Heap, string) {
	return h[:len(h)-1], h[len(h)-1]
}

func ParseContainerStacks(stackString string, enhanced bool) (string, error) {
	exp := regexp.MustCompile(`((?:(?:\[[A-Z]\]\s+)+\[[A-Z]\]\n)+)(\s(?:\d+\s+)+)\n((?:move \d+ from \d+ to \d+\n)+)`)

	// If all went okay, we will get 4 results in here
	matches := exp.FindStringSubmatch(stackString)

	// 2nd group contains numbers of stacks, get last number for amount of stacks
	numStacks, err := strconv.ParseInt(matches[2][len(matches[2])-3:len(matches[2])-2], 10, 64)
	if err != nil {
		return "", err
	}

	// Create number of stacks
	stacks := make([]Heap, numStacks)

	// Container layout is in 1st group
	containerRows := strings.Split(strings.TrimSuffix(matches[1], "\n"), "\n")

	// Populate stacks, iterate from bottom up
	for i := len(containerRows) - 1; i >= 0; i-- {
		for j := 0; j < int(numStacks); j++ {

			// iterate over container slots in row
			payload := string(containerRows[i][1+j*4])

			// if there is a container, push to its respective stack
			if payload != " " {
				stacks[j] = stacks[j].Put(payload)
			}
		}
	}

	if !enhanced {
		return ParseCrateMover9000(matches[3], stacks)
	} else {
		return ParseCrateMover9001(matches[3], stacks)
	}

}

func ParseCrateMover9000(instructions string, stacks []Heap) (string, error) {
	// Parse move instructions in 3rd group
	for _, instruction := range strings.Split(strings.TrimSuffix(instructions, "\n"), "\n") {
		exp := regexp.MustCompile(`move (\d+) from (\d+) to (\d+)`)
		numbers := exp.FindStringSubmatch(instruction)

		amount, err := strconv.ParseInt(numbers[1], 10, 64)
		if err != nil {
			return "", err
		}

		from, err := strconv.ParseInt(numbers[2], 10, 64)
		if err != nil {
			return "", err
		}

		to, err := strconv.ParseInt(numbers[3], 10, 64)
		if err != nil {
			return "", err
		}

		for k := 0; k < int(amount); k++ {
			// pop from stack
			poppedStack, toPut := stacks[from-1].Pop()
			// save old stack
			stacks[from-1] = poppedStack

			// put onto new stack and save
			stacks[to-1] = stacks[to-1].Put(toPut)
		}
	}

	result := ""
	for l := 0; l < int(len(stacks)); l++ {
		_, partialResult := stacks[l].Pop()

		result += partialResult
	}

	return result, nil
}

func ParseCrateMover9001(instructions string, stacks []Heap) (string, error) {
	// Parse move instructions in 3rd group
	for _, instruction := range strings.Split(strings.TrimSuffix(instructions, "\n"), "\n") {
		exp := regexp.MustCompile(`move (\d+) from (\d+) to (\d+)`)
		numbers := exp.FindStringSubmatch(instruction)

		amount, err := strconv.ParseInt(numbers[1], 10, 64)
		if err != nil {
			return "", err
		}

		from, err := strconv.ParseInt(numbers[2], 10, 64)
		if err != nil {
			return "", err
		}

		to, err := strconv.ParseInt(numbers[3], 10, 64)
		if err != nil {
			return "", err
		}

		stacks[to-1] = stacks[to-1].PutStack(stacks[from-1][len(stacks[from-1])-int(amount):])
		stacks[from-1] = stacks[from-1][:len(stacks[from-1])-int(amount)]
	}

	result := ""
	for l := 0; l < int(len(stacks)); l++ {
		_, partialResult := stacks[l].Pop()

		result += partialResult
	}

	return result, nil
}
