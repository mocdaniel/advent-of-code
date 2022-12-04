package day04

import (
	"strconv"
	"strings"
)

func DetectOverlap(input string, total bool) (bool, error) {
	sections := strings.Split(input, ",")

	section1Split := strings.Split(sections[0], "-")
	section2Split := strings.Split(sections[1], "-")

	a, err := strconv.ParseInt(section1Split[0], 10, 64)
	if err != nil {
		return false, err
	}

	b, err := strconv.ParseInt(section1Split[1], 10, 64)
	if err != nil {
		return false, err
	}

	c, err := strconv.ParseInt(section2Split[0], 10, 64)
	if err != nil {
		return false, err
	}

	d, err := strconv.ParseInt(section2Split[1], 10, 64)
	if err != nil {
		return false, err
	}

	if total {
		return ((a <= c && b >= d) || (a >= c && b <= d)), nil
	} else {
		return ((a >= c && a < d) || (b >= c && b < d) || (c >= a && c < b) || (d >= a && d < b) || (a > c && a <= d) || (b > c && b <= d) || (c > a && c <= b) || (d > a && d <= b)), nil
	}
}
