package main

import (
	"fmt"
	"slices"
	"testing"
)

var in string = `seeds: 79 14 55 13

seed-to-soil map:
50 98 2
52 50 48

soil-to-fertilizer map:
0 15 37
37 52 2
39 0 15

fertilizer-to-water map:
49 53 8
0 11 42
42 0 7
57 7 4

water-to-light map:
88 18 7
18 25 70

light-to-temperature map:
45 77 23
81 45 19
68 64 13

temperature-to-humidity map:
0 69 1
1 0 69

humidity-to-location map:
60 56 37
56 93 4`

func TestParseSeeds(t *testing.T) {
	p := NewInputParser(in)
	wanted := `temperature-to-humidity map:
0 69 1
1 0 69
`
	actual := p.GetSectionString("temperature-to-humidity map")
	if wanted != actual {
		t.Fatalf("wanted %s, got %s", wanted, actual)
	}
}

func TestGetSectionString(t *testing.T) {
	p := NewInputParser(in)
	wanted := []int{79, 14, 55, 13}
	actual := p.GetSeeds()
	if !slices.Equal(wanted, actual) {
		t.Fatalf("wanted %d, got %d", wanted, actual)
	}
}

func TestGetSectionMap(t *testing.T) {
	p := NewInputParser(in)
	m := p.GetSectionMap("temperature-to-humidity")
	wanted := "temperature-to-humidity"
	actual := m.name
	if wanted != actual {
		t.Fatalf("wanted %s, got %s", wanted, actual)
	}

	wantedInt := 0
	actualInt := m.data[0].destinationRangeStart
	if wantedInt != actualInt {
		t.Fatalf("wanted %d, got %d", wantedInt, actualInt)
	}

	wantedInt = 69
	actualInt = m.data[0].sourceRangeStart
	if wantedInt != actualInt {
		t.Fatalf("wanted %d, got %d", wantedInt, actualInt)
	}

	wantedInt = 1
	actualInt = m.data[0].rangeLength
	if wantedInt != actualInt {
		t.Fatalf("wanted %d, got %d", wantedInt, actualInt)
	}

	wantedInt = 1
	actualInt = m.data[1].destinationRangeStart
	if wantedInt != actualInt {
		t.Fatalf("wanted %d, got %d", wantedInt, actualInt)
	}

	wantedInt = 0
	actualInt = m.data[1].sourceRangeStart
	if wantedInt != actualInt {
		t.Fatalf("wanted %d, got %d", wantedInt, actualInt)
	}

	wantedInt = 69
	actualInt = m.data[1].rangeLength
	if wantedInt != actualInt {
		t.Fatalf("wanted %d, got %d", wantedInt, actualInt)
	}
}

func TestGetParsedSectionMap(t *testing.T) {
	p := NewInputParser(in)
	m := p.GetParsedSectionMap("temperature-to-humidity")

	wanted := "temperature"
	actual := m.data[0].sourceName
	if wanted != actual {
		t.Fatalf("wanted %s, got %s", wanted, actual)
	}

	wanted = "humidity"
	actual = m.data[0].destName
	if wanted != actual {
		t.Fatalf("wanted %s, got %s", wanted, actual)
	}

	wantedInt := 69
	actualInt := m.data[0].source.lowerBound
	if wantedInt != actualInt {
		t.Fatalf("wanted %d, got %d", wantedInt, actualInt)
	}

	wantedInt = 0
	actualInt = m.data[0].dest.lowerBound
	if wantedInt != actualInt {
		t.Fatalf("wanted %d, got %d", wantedInt, actualInt)
	}

	wantedInt = 69
	actualInt = m.data[0].source.upperBound
	if wantedInt != actualInt {
		t.Fatalf("wanted %d, got %d", wantedInt, actualInt)
	}

	wantedInt = 0
	actualInt = m.data[0].dest.lowerBound
	if wantedInt != actualInt {
		t.Fatalf("wanted %d, got %d", wantedInt, actualInt)
	}

	wantedInt = 0
	actualInt = m.data[1].source.lowerBound
	if wantedInt != actualInt {
		t.Fatalf("wanted %d, got %d", wantedInt, actualInt)
	}

	wantedInt = 1
	actualInt = m.data[1].source.offset
	if wantedInt != actualInt {
		t.Fatalf("wanted %d, got %d", wantedInt, actualInt)
	}

	wantedInt = 1
	actualInt = m.data[1].dest.lowerBound
	if wantedInt != actualInt {
		t.Fatalf("wanted %d, got %d", wantedInt, actualInt)
	}

	wantedInt = 69
	actualInt = m.data[1].dest.upperBound
	if wantedInt != actualInt {
		t.Fatalf("wanted %d, got %d", wantedInt, actualInt)
	}

	wantedInt = 68
	actualInt = m.data[1].source.upperBound
	if wantedInt != actualInt {
		t.Fatalf("wanted %d, got %d", wantedInt, actualInt)
	}
}

func TestGetDestId(t *testing.T) {
	p := NewInputParser(in)
	m := p.GetParsedSectionMap("seed-to-soil")

	wanted := 81
	actual := m.GetDestId(79)
	if wanted != actual {
		t.Fatalf("wanted %d, got %d", wanted, actual)
	}
}

func TestTraverse(t *testing.T) {
	p := NewInputParser(in)
	actual := p.Traverse()
	wanted := []int{82, 43, 86, 35}
	if !slices.Equal(wanted, actual) {
		t.Fatalf("wanted %v, got %v", wanted, actual)
	}
}

func TestSeedsPartTwo(t *testing.T) {
	p := NewInputParser(in)
	actual := p.GetSeedsPartTwo()
	fmt.Printf("ACTUAL: %v", actual)
}
