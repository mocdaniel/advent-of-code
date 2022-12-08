package day08

import "testing"

func TestScenery(t *testing.T) {
	got := [][]int{
		{3, 0, 3, 7, 3}, {2, 5, 5, 1, 2}, {6, 5, 3, 3, 2}, {3, 3, 5, 4, 9}, {3, 5, 3, 9, 0},
	}

	if HighestScenicScore(got) != 8 {
		t.Errorf("Not correct")
	}
}
