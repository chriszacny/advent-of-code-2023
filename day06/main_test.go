package main

import (
	"testing"
)

var in string = `Time:      7  15   30
Distance:  9  40  200`

func TestPartOne(t *testing.T) {
	p := NewInputParser(in)
	races := p.GetRaces()
	ways := []int{4, 8, 9}
	for i := 0; i < 3; i++ {
		actual := getWaysToWin(races[i])
		if ways[i] != actual {
			t.Fatalf("wanted %d, got %d", ways[i], actual)
		}
	}
	total := ways[0]
	for i := 1; i < len(ways); i++ {
		total = total * ways[i]
	}
	wanted := 288
	if wanted != total {
		t.Fatalf("wanted %d, got %d", wanted, total)
	}
}

func TestPartTwo(t *testing.T) {
	p := NewInputParser(in)
	race := p.GetRace()
	wanted := 71503
	actual := getWaysToWin(race)
	if wanted != actual {
		t.Fatalf("wanted %d, got %d", wanted, actual)
	}
}
