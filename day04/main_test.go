package main

import (
	"slices"
	"testing"
)

func TestParser(t *testing.T) {
	in := "Card 214: 86  8  2 92 | 97 54 53"
	c := NewCard(in)
	want := 214
	actual := c.id
	if want != actual {
		t.Fatalf("wanted %d, got %d", want, actual)
	}

	wantSlice := []int{86, 8, 2, 92}
	actualSlice := c.winningNums
	if !slices.Equal(wantSlice, actualSlice) {
		t.Fatalf("wanted %d, got %d", wantSlice, actualSlice)
	}

	wantSlice = []int{97, 54, 53}
	actualSlice = c.ourNums
	if !slices.Equal(wantSlice, actualSlice) {
		t.Fatalf("wanted %d, got %d", wantSlice, actualSlice)
	}
}

func TestBasic(t *testing.T) {
	in := "Card 1: 1 2 | 1 2 3"
	c := NewCard(in)

	want := 2
	actual := c.Value()
	if want != actual {
		t.Fatalf("wanted %d, got %d", want, actual)
	}
}
