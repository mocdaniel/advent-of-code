package day02

import (
	"fmt"
	"strings"
)

/* A -> Rock
*  B -> Paper
*  C -> Scissors
*  X -> Rock (1 pt)
*  Y -> Paper (2 pts)
*  Z -> Scissors (3 pts)
 */
var ResultMapA = map[string]int{
	"A X": 4,
	"A Y": 8,
	"A Z": 3,
	"B X": 1,
	"B Y": 5,
	"B Z": 9,
	"C X": 7,
	"C Y": 2,
	"C Z": 6,
}

/* A -> Rock
*  B -> Paper
*  C -> Scissors
*  X -> Lose
*  Y -> Draw
*  Z -> Win
 */
var ResultMapB = map[string]int{
	"A X": 3, // Scissors
	"A Y": 4, // Rock
	"A Z": 8, // Paper
	"B X": 1, // Rock
	"B Y": 5, // Paper
	"B Z": 9, // Scissors
	"C X": 2, // Paper
	"C Y": 6, // Scissors
	"C Z": 7, // Rock
}

func GetTotalScore(input string, resultMap map[string]int) int {
	splitGames := strings.Split(input, "\n")

	fmt.Printf("Found %v game records\n\n", len(splitGames))

	totalScore := 0
	for _, game := range splitGames {
		totalScore += resultMap[game]
	}

	return totalScore
}
