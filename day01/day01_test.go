package day01

import "testing"

func TestGetElfWithMostFood(t *testing.T) {
	testString := `1000
2000
3000

4000

5000
6000

7000
8000
9000

10000`

	result, err := GetElfWithMostFood(testString)
	if err != nil {
		t.Error(err)
	}

	if result != 24000 {
		t.Errorf("Got: %v; Wanted: %v", result, 24000)
	}
}

func TestGetThreeElvesWithMostFood(t *testing.T) {
	testString := `1000
2000
3000

4000

5000
6000

7000
8000
9000

10000`

	result, err := GetThreeElvesWithMostFood(testString)
	if err != nil {
		t.Error(err)
	}

	if result != 45000 {
		t.Errorf("Got: %v; Wanted: %v", result, 45000)
	}
}

func TestSumArray(t *testing.T) {
	cases := []struct {
		name string
		got  []int
		want int
	}{
		{
			name: "Empty array",
			got:  []int{},
			want: 0,
		},
		{
			name: "Normal array",
			got:  []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
			want: 55,
		},
		{
			name: "Array with single entry",
			got:  []int{42},
			want: 42,
		},
		{
			name: "Array with duplicated entries",
			got:  []int{1, 3, 3, 4, 5, 5, 5, 6},
			want: 32,
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			sum := sumArray(tc.got)
			if sum != tc.want {
				t.Errorf("Got: %v; Want: %v", sum, tc.want)
			}
		})
	}
}

func TestFindMinValue(t *testing.T) {
	cases := []struct {
		name      string
		got       []int
		idxWant   int
		valueWant int
	}{
		{
			name:      "All entries equal",
			got:       []int{0, 0, 0},
			idxWant:   0,
			valueWant: 0,
		},
		{
			name:      "One smallest entry",
			got:       []int{1, 0, 2},
			idxWant:   1,
			valueWant: 0,
		},
		{
			name:      "Two smallest entries",
			got:       []int{3, 1, 1},
			idxWant:   1,
			valueWant: 1,
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			value, idx := findMinValue(tc.got)
			if value != tc.valueWant && idx != tc.idxWant {
				t.Errorf("Got: (%v, %v); Want: (%v, %v)", value, idx, tc.idxWant, tc.valueWant)
			}
		})
	}
}
