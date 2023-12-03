package main

import (
	"fmt"
	"testing"
)

func TestBasic(t *testing.T) {
	in := `..42.
...*.
5..8.
.$...
....3`

	fmt.Printf("%s", in)

	s1, adj := partOne(in)

	want := 55
	actual := s1
	if want != actual {
		t.Fatalf("wanted %d, got %d", want, actual)
	}

	s2 := partTwo(adj)
	want = 336
	actual = s2
	if want != actual {
		t.Fatalf("wanted %d, got %d", want, actual)
	}

}
