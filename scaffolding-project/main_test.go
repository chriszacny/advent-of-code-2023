package main

import (
	"testing"
)

func TestBasic(t *testing.T) {
	actual := 1
	want := 2
	if actual != want {
		t.Fatalf("wanted %d, got %d", want, actual)
	}
}
