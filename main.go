package main

import (
	"fmt"
	"log"
	"os"

	"github.com/mocdaniel/advent-of-code/day01"
	"github.com/mocdaniel/advent-of-code/day02"
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
		},
		Name:  "Advent of Code 2022",
		Usage: "aoc dayXX INPUT",
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
