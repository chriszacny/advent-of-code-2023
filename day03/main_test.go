package main

import (
	"testing"
)

func TestA(t *testing.T) {
	want := 1
	actual := 2
	if want != actual {
		t.Fatalf("wanted %d, got %d", want, actual)
	}
}
