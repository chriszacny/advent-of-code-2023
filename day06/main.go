package main

import (
	"fmt"
	"strconv"
	"strings"
)

type Race struct {
	totaltime    int
	bestdistance int
}

/***********************************************************************************
 * Struct Name: InputParser
 *
 ***********************************************************************************/

type InputParser interface {
	GetRaces() []Race
	GetRace() Race
}

type Parser struct {
	rawinput string
}

func NewInputParser(in string) InputParser {
	return &Parser{rawinput: in}
}

func (p *Parser) parseline(s string, existing []int) []int {
	strs := strings.Split(s, " ")
	sanitizedstrs := []string{}
	for _, v := range strs {
		if v != "" && !strings.HasPrefix(v, "Time") && !strings.HasPrefix(v, "Distance") {
			sanitizedstrs = append(sanitizedstrs, v)
		}
	}
	for _, v := range sanitizedstrs {
		vint, err := strconv.Atoi(v)
		if err != nil {
			fmt.Printf("%v\n", err.Error())
		} else {
			existing = append(existing, vint)
		}
	}
	return existing
}

func (p *Parser) GetRaces() []Race {
	ret := []Race{}
	times := []int{}
	distances := []int{}
	for _, v := range strings.Split(p.rawinput, "\n") {
		if strings.HasPrefix(v, "Time:") {
			times = p.parseline(v, times)
		} else if strings.HasPrefix(v, "Distance") {
			distances = p.parseline(v, distances)
		}
	}
	for i := 0; i < len(times); i++ {
		r := Race{totaltime: times[i], bestdistance: distances[i]}
		ret = append(ret, r)
	}
	return ret
}

func (p *Parser) GetRace() Race {
	in := p.GetRaces()
	timeStr := ""
	distStr := ""
	for _, r := range in {
		timeStr += fmt.Sprint(r.totaltime)
		distStr += fmt.Sprint(r.bestdistance)
	}
	time, err := strconv.Atoi(timeStr)
	if err != nil {
		panic("")
	}
	dist, err := strconv.Atoi(distStr)
	if err != nil {
		panic("")
	}
	return Race{totaltime: time, bestdistance: dist}
}

/***********************************************************************************
 * End of InputParser
 ***********************************************************************************/

func getWaysToWin(r Race) int {
	wins := 0
	for i := 0; i <= r.totaltime; i++ {
		speed := i
		dist := speed * (r.totaltime - i)
		if dist > r.bestdistance {
			wins++
		}
	}
	return wins
}

func main() {
	// Boilerplate setup
	in, file := getInput()
	defer file.Close()

	// part one
	p := NewInputParser(in)
	races := p.GetRaces()
	wins := []int{}
	for _, v := range races {
		wins = append(wins, getWaysToWin(v))
	}
	total := wins[0]
	if len(wins) == 0 {
		total = 0
	} else if len(wins) == 1 {
		total = wins[0]
	} else {
		for i := 1; i < len(wins); i++ {
			total = total * wins[i]
		}
	}
	fmt.Printf("part one: %d\n", total)

	// part two - 475,213,810,151,650
	raceP2 := p.GetRace()
	totalP2 := getWaysToWin(raceP2)
	fmt.Printf("part one: %d\n", totalP2)
}
