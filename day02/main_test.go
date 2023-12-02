package main

import (
	"testing"
)

func TestFewestNumPossibleCubesPerGame(t *testing.T) {
	gr := gameRecord{id: 1}
	gr.rounds = append(gr.rounds, gameRound{red: 2})
	gr.rounds = append(gr.rounds, gameRound{red: 1, green: 3, blue: 2})
	red, green, blue := fewestNumPossibleCubesPerGame(gr)

	want := 2
	actual := red
	if want != actual {
		t.Fatalf("wanted %d, got %d", want, actual)
	}
	want = 3
	actual = green
	if want != actual {
		t.Fatalf("wanted %d, got %d", want, actual)
	}
	want = 2
	actual = blue
	if want != actual {
		t.Fatalf("wanted %d, got %d", want, actual)
	}
}

func TestGamesPossible(t *testing.T) {
	puzzleInput := `1:3b,4r;1r,2g,6b;2g
2:1b,2g;3g,4b,1r;1g,1b
3:8g,6b,20r;5b,4r,13g;5g,1r
4:1g,3r,6b;3g,6r;3g,15b,14r
5:6r,1b,3g;2b,1r,2g`

	gp := gamesPossible(puzzleInput, 12, 13, 14)

	want := 3
	actual := len(gp)
	if want != actual {
		t.Fatalf("wanted %d, got %d", want, actual)
	}

	want = 1
	actual = gp[0].id
	if want != actual {
		t.Fatalf("wanted %d, got %d", want, actual)
	}

	want = 2
	actual = gp[1].id
	if want != actual {
		t.Fatalf("wanted %d, got %d", want, actual)
	}

	want = 5
	actual = gp[2].id
	if want != actual {
		t.Fatalf("wanted %d, got %d", want, actual)
	}
}

func TestParseGamesFileLine(t *testing.T) {
	line := "1234:2r,2g;6r,3g;2r,1g,2b;1r"
	parsedline := parseGamesFileLine(line)

	want := 1234
	actual := parsedline.id
	if want != actual {
		t.Fatalf("wanted %d, got %d", want, actual)
	}

	want = 4
	actual = len(parsedline.rounds)
	if want != actual {
		t.Fatalf("wanted %d, got %d", want, actual)
	}

	wantStr := "&main.gameRecord{id:1234, rounds:[]main.gameRound{main.gameRound{red:2, green:2, blue:0}, main.gameRound{red:6, green:3, blue:0}, main.gameRound{red:2, green:1, blue:2}, main.gameRound{red:1, green:0, blue:0}}}"
	actualStr := parsedline.String()
	if wantStr != actualStr {
		t.Fatalf("wanted %s, got %s", wantStr, actualStr)
	}
}
