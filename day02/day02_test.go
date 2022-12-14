package day02

import "testing"

func TestCalculateScore(t *testing.T) {
	input := `A Y
B X
C Z`
	score := GetTotalScore(input, ResultMapA)

	if score != 15 {
		t.Errorf("Got: %v; Want: %v", score, 15)
	}

	score = GetTotalScore(input, ResultMapB)

	if score != 12 {
		t.Errorf("Got: %v; Want: %v", score, 12)
	}
}
