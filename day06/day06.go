package day06

import "fmt"

func GetStartOfSequence(input string, amount int) int {
	history := make([]byte, 0)

	for i := 0; i < len(input); i++ {
		contains, index := in(input[i], history)
		if contains {
			fmt.Println("===CHANGE===")
			fmt.Printf("Before: %#v\n", history)
			history = resetHistory(history, index)
			fmt.Printf("After: %#v\n", history)
			fmt.Println("===END===")
		} else {
			if len(history) == amount-1 {
				return i + 1
			}
		}
		history = append(history, input[i])
		fmt.Printf("%#v\n", history)
	}

	return -1
}

func in(s byte, h []byte) (bool, int) {
	for i, c := range h {
		if c == s {
			return true, i
		}
	}

	return false, -1
}

func resetHistory(h []byte, index int) []byte {
	tempHist := make([]byte, len(h)-index-1)
	for i := 0; i < len(h)-index-1; i++ {
		tempHist[i] = h[index+i+1]
	}

	return tempHist
}
