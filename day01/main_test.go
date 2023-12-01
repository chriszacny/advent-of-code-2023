package main

import (
	"testing"
)

func TestGetCalibrationValue(t *testing.T) {
	actual := getCV("1abc2")

	want := 12
	if actual != want {
		t.Fatalf("wanted %d, got %d", want, actual)
	}
	actual = getCV("pqr3stu8vwx")
	want = 38
	if actual != want {
		t.Fatalf("wanted %d, got %d", want, actual)
	}
	actual = getCV("a1b2c3d4e5f")
	want = 15
	if actual != want {
		t.Fatalf("wanted %d, got %d", want, actual)
	}
	actual = getCV("treb7uchet")
	want = 77
	if actual != want {
		t.Fatalf("wanted %d, got %d", want, actual)
	}
	actual = getCV("42")
	want = 42
	if actual != want {
		t.Fatalf("wanted %d, got %d", want, actual)
	}
	actual = getCV("8j")
	want = 88
	if actual != want {
		t.Fatalf("wanted %d, got %d", want, actual)
	}
	actual = getCV("1")
	want = 11
	if actual != want {
		t.Fatalf("wanted %d, got %d", want, actual)
	}
}


func TestSumOfCalibrationValues(t *testing.T) {
	input := `1abc2
	pqr3stu8vwx
	a1b2c3d4e5f
	treb7uchet`

	want := 142
	actual := sumOfCVs(input)
	if actual != want {
		t.Fatalf("wanted %d, got %d", want, actual)
	}
}
