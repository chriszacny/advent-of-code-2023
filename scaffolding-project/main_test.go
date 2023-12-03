package main

import (
	"testing"
	"fmt"
)

var testin string = `
	467..114..
	...*......
	..35..633.
	......#...
	617*......
	.....+.58.
	..592.....
	......755.
	...$.*....
	.664.598..`

func TestA(t *testing.T) {
	fmt.Printf("%s", cleanIn(testin))

	/*
	want := 1
	actual := 2
	if want != actual {
		t.Fatalf("wanted %d, got %d", want, actual)
	}
	*/
}
