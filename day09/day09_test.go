package day09

import "testing"

func TestGetVisitedPositions(t *testing.T) {
	input := `R 4
U 4
L 3
D 1
R 4
D 1
L 5
R 2
`

	got, err := GetVisitedPositions(input, 2)
	if err != nil {
		t.Error(err)
	}

	if got != 13 {
		t.Errorf("Got: %v; Want 13", got)
	}
}
