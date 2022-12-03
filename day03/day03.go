package day03

import "strings"

func FindDuplicatedItem(contents string) byte {
	length := len(contents)

	firstHalf := make([]byte, 0)

	for i := 0; i < length/2; i++ {
		firstHalf = append(firstHalf, contents[i])
	}

	for i := length / 2; i < length; i++ {
		if contains(firstHalf, contents[i]) {
			return contents[i]
		}
	}

	return 0
}

func FindGroupPriority(contents []string) byte {
	for _, c := range contents[0] {
		if strings.Contains(contents[1], string(c)) && strings.Contains(contents[2], string(c)) {
			return byte(c)
		}
	}

	return 0
}

func contains(slice []byte, item byte) bool {
	for _, i := range slice {
		if i == item {
			return true
		}
	}

	return false
}

func GetPriority(item byte) int {
	lowerCaseConst := 96
	upperCaseConst := 38

	if item > 96 {
		return int(item) - lowerCaseConst
	} else {
		return int(item) - upperCaseConst
	}
}
