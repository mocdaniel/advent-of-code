package day03

import (
	"strings"
	"testing"
)

func TestFindDuplicatedItem(t *testing.T) {
	scenarios := []struct {
		name string
		got  string
		want byte
	}{
		{
			name: "Find p",
			got:  "vJrwpWtwJgWrhcsFMMfFFhFp",
			want: 'p',
		},
		{
			name: "Find L",
			got:  "jqHRNqRjqzjGDLGLrsFMfFZSrLrFZsSL",
			want: 'L',
		},
		{
			name: "Find P",
			got:  "PmmdzqPrVvPwwTWBwg",
			want: 'P',
		},
		{
			name: "Find v",
			got:  "wMqvLMZHhHMvwLHjbvcjnnSBnvTQFn",
			want: 'v',
		},
		{
			name: "Find t",
			got:  "ttgJtRGJQctTZtZT",
			want: 't',
		},
		{
			name: "Find s",
			got:  "CrZsJsPPZsGzwwsLwLmpwMDw",
			want: 's',
		},
	}

	for _, s := range scenarios {
		t.Run(s.name, func(t *testing.T) {
			found := FindDuplicatedItem(s.got)
			if found != s.want {
				t.Errorf("Got: %v; Want: %v", found, s.want)
			}
		})
	}
}

func TestGetPriority(t *testing.T) {
	scenarios := []struct {
		name string
		got  byte
		want int
	}{
		{
			name: "Want 16",
			got:  'p',
			want: 16,
		},
		{
			name: "Want 38",
			got:  'L',
			want: 38,
		},
		{
			name: "Want 42",
			got:  'P',
			want: 42,
		},
		{
			name: "Want 22",
			got:  'v',
			want: 22,
		},
		{
			name: "Want 20",
			got:  't',
			want: 20,
		},
		{
			name: "Want 19",
			got:  's',
			want: 19,
		},
		{
			name: "Want 52",
			got:  'Z',
			want: 52,
		},
		{
			name: "Want 27",
			got:  'A',
			want: 27,
		},
		{
			name: "Want 1",
			got:  'a',
			want: 1,
		},
		{
			name: "Want 26",
			got:  'z',
			want: 26,
		},
	}

	for _, s := range scenarios {
		t.Run(s.name, func(t *testing.T) {
			got := GetPriority(s.got)
			if got != s.want {
				t.Errorf("Got: %v; Want: %v", got, s.want)
			}
		})
	}
}

func TestFindGroupPriority(t *testing.T) {
	scenarios := []struct {
		name string
		got  []string
		want byte
	}{
		{
			name: "Want 16",
			got: strings.Split(`vJrwpWtwJgWrhcsFMMfFFhFp
jqHRNqRjqzjGDLGLrsFMfFZSrLrFZsSL
PmmdzqPrVvPwwTWBwg`, "\n"),
			want: 'r',
		},
		{
			name: "Want 38",
			got: strings.Split(`wMqvLMZHhHMvwLHjbvcjnnSBnvTQFn
ttgJtRGJQctTZtZT
CrZsJsPPZsGzwwsLwLmpwMDw`, "\n"),
			want: 'Z',
		},
	}

	for _, s := range scenarios {
		t.Run(s.name, func(t *testing.T) {
			got := FindGroupPriority(s.got)
			if got != s.want {
				t.Errorf("Want: %v; Got: %v", s.want, got)
			}
		})
	}
}
