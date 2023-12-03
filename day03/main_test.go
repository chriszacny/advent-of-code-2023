package main

import (
	"testing"
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

func TestBasic(t *testing.T) {
	testSimple := `
	..42*
	.....
	5....
	.$...
	....3`

	nums := getAllNums(testSimple)

	want := 2
	actual := len(nums)
	if want != actual {
		t.Fatalf("wanted %d, got %d", want, actual)
	}

	want = 42
	actual = nums[0]
	if want != actual {
		t.Fatalf("wanted %d, got %d", want, actual)
	}

	want = 5
	actual = nums[1]
	if want != actual {
		t.Fatalf("wanted %d, got %d", want, actual)
	}
}
