package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/mocdaniel/advent-of-code/day01"
	"github.com/mocdaniel/advent-of-code/day02"
	"github.com/mocdaniel/advent-of-code/day03"
	"github.com/mocdaniel/advent-of-code/day04"
	"github.com/mocdaniel/advent-of-code/day05"
	"github.com/mocdaniel/advent-of-code/day06"
	"github.com/mocdaniel/advent-of-code/day07"
	"github.com/mocdaniel/advent-of-code/day08"
	"github.com/mocdaniel/advent-of-code/day09"
	"github.com/mocdaniel/advent-of-code/day10"
	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Commands: []*cli.Command{
			{
				Name:    "day01",
				Aliases: []string{"1"},
				Action: func(cCtx *cli.Context) error {

					input, err := os.ReadFile(cCtx.Args().First())
					if err != nil {
						return err
					}
					result1, err := day01.GetElfWithMostFood(string(input))
					if err != nil {
						return err
					}

					result2, err := day01.GetThreeElvesWithMostFood(string(input))
					if err != nil {
						return err
					}

					fmt.Printf("The most calories carried by a single elf are %v\n", result1)
					fmt.Printf("The most calories carried by three elves are %v\n", result2)
					return nil
				},
			},
			{
				Name:    "day02",
				Aliases: []string{"2"},
				Action: func(cCtx *cli.Context) error {

					input, err := os.ReadFile(cCtx.Args().First())
					if err != nil {
						return err
					}

					inputString := string(input)
					result1 := day02.GetTotalScore(inputString, day02.ResultMapA)
					result2 := day02.GetTotalScore(inputString, day02.ResultMapB)

					fmt.Printf("The total score according to the strategy guide would be %v\n", result1)
					fmt.Printf("The total score according to the decrypted strategy guide would be %v\n", result2)
					return nil
				},
			},
			{
				Name:    "day03",
				Aliases: []string{"3"},
				Action: func(cCtx *cli.Context) error {

					input, err := os.ReadFile(cCtx.Args().First())
					if err != nil {
						return err
					}

					inputStrings := strings.Split(strings.TrimSuffix(string(input), "\n"), "\n")
					result1 := 0

					for _, p := range inputStrings {
						result1 += day03.GetPriority(day03.FindDuplicatedItem(p))
					}

					result2 := 0
					tempStrings := make([]string, 0)

					for i, p := range inputStrings {
						tempStrings = append(tempStrings, p)

						if i%3 == 2 {
							result2 += day03.GetPriority(day03.FindGroupPriority(tempStrings))
							tempStrings = tempStrings[:0]
						}
					}

					fmt.Printf("The total priority score of all misplaced items is %v\n", result1)
					fmt.Printf("The total priority of group identifiers is %v\n", result2)
					return nil
				},
			},
			{
				Name:    "day04",
				Aliases: []string{"4"},
				Action: func(cCtx *cli.Context) error {

					input, err := os.ReadFile(cCtx.Args().First())
					if err != nil {
						return err
					}

					inputStrings := strings.Split(strings.TrimSuffix(string(input), "\n"), "\n")
					result1 := 0
					result2 := 0

					for _, i := range inputStrings {
						total_overlapping, err := day04.DetectOverlap(i, true)
						if err != nil {
							return err
						}
						if total_overlapping {
							result1 += 1
						}

						overlapping, err := day04.DetectOverlap(i, false)
						if err != nil {
							return err
						}
						if overlapping {
							result2 += 1
						}
					}

					fmt.Printf("The amount of entirely overlapping section pairs is %v\n", result1)
					fmt.Printf("The amount of overlapping section pairs is %v\n", result2)
					return nil
				},
			},
			{
				Name:    "day05",
				Aliases: []string{"5"},
				Action: func(cCtx *cli.Context) error {

					input, err := os.ReadFile(cCtx.Args().First())
					if err != nil {
						return err
					}

					result1, err := day05.ParseContainerStacks(string(input), false)
					if err != nil {
						return err
					}

					result2, err := day05.ParseContainerStacks(string(input), true)
					if err != nil {
						return err
					}

					fmt.Printf("The containers on top of each stack after operating with model 9000 spell %v\n", result1)
					fmt.Printf("The containers on top of each stack after operating with model 9001 spell %v\n", result2)
					return nil
				},
			},
			{
				Name:    "day06",
				Aliases: []string{"6"},
				Action: func(cCtx *cli.Context) error {

					input, err := os.ReadFile(cCtx.Args().First())
					if err != nil {
						return err
					}

					inputString := string(input)
					result1 := day06.GetStartOfSequence(inputString, 4)
					result2 := day06.GetStartOfSequence(inputString, 14)

					fmt.Printf("The number of characters needed to be processed before the first packet marker is %v\n", result1)
					fmt.Printf("The number of characters needed to be processed before the first message marker is %v\n", result2)
					return nil
				},
			},
			{
				Name:    "day07",
				Aliases: []string{"7"},
				Action: func(cCtx *cli.Context) error {

					input, err := os.ReadFile(cCtx.Args().First())
					if err != nil {
						return err
					}

					inputString := strings.Split(strings.TrimSuffix(string(input), "\n"), "\n")
					tree := day07.BuildFileSystemTree(inputString)

					result1 := tree.GetSizeBelowNumber(100000)
					fmt.Printf("The total size of all directories smaller than 100000 units is %v\n", result1)

					result2 := tree.GetSize()
					tree.GetSmallestAbove(30000000-(70000000-tree.GetSize()), &result2)
					fmt.Printf("The total size of the smallest directory needed to be cleaned is %v\n", result2)

					return nil
				},
			},
			{
				Name:    "day08",
				Aliases: []string{"8"},
				Action: func(cCtx *cli.Context) error {

					input, err := os.ReadFile(cCtx.Args().First())
					if err != nil {
						return err
					}

					grid, err := day08.ParseGrid(string(input))
					if err != nil {
						return err
					}

					result1 := day08.CountTrees(grid)
					result2 := day08.HighestScenicScore(grid)

					fmt.Printf("The amount of trees visible from outside the grid is %v\n", result1)
					fmt.Printf("The highest possible scenic score is %v\n", result2)

					return nil
				},
			},
			{
				Name:    "day09",
				Aliases: []string{"9"},
				Action: func(cCtx *cli.Context) error {

					input, err := os.ReadFile(cCtx.Args().First())
					if err != nil {
						return err
					}

					result1, err := day09.GetVisitedPositions(string(input), 2)
					if err != nil {
						return err
					}

					result2, err := day09.GetVisitedPositions(string(input), 10)
					if err != nil {
						return err
					}

					fmt.Printf("The amount of unique positions visited by the tail of a short rope within the grid is %v\n", result1)
					fmt.Printf("The amount of unique positions visited by the tail of a long rope within the grid is %v\n", result2)

					return nil
				},
			},
			{
				Name:    "day10",
				Aliases: []string{"10"},
				Action: func(cCtx *cli.Context) error {

					input, err := os.ReadFile(cCtx.Args().First())
					if err != nil {
						return err
					}

					instructions := strings.Split(strings.TrimSuffix(string(input), "\n"), "\n")

					cycles := []int{20, 60, 100, 140, 180, 220}
					result1 := 0
					for _, c := range cycles {
						s, err := day10.GetSignalStrength(instructions, c, false)
						if err != nil {
							return err
						}
						result1 += s
					}

					fmt.Printf("The added signal strength of the requested cycles is %v\n", result1)

					_, err = day10.GetSignalStrength(instructions, -1, true)
					if err != nil {
						return err
					}
					return nil
				},
			},
		},
		Name:  "Advent of Code 2022",
		Usage: "aoc dayXX INPUT",
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
