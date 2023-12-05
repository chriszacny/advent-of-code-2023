package main

import (
	"fmt"
	"slices"
	"strconv"
	"strings"
)

/***********************************************************************************
 * Struct Name: Card
 *
 ***********************************************************************************/
type Card struct {
	id          int
	winningNums []int
	ourNums     []int
}

func NewCard(rawInput string) *Card {
	nc := Card{}
	left := strings.Split(rawInput, ":")[0]
	idlist := strings.Split(left, " ")
	idstr := idlist[len(idlist)-1]
	id, err := strconv.Atoi(idstr)
	if err != nil {
		panic("")
	}
	nc.id = id

	right := strings.Split(rawInput, ":")[1]
	winningNumsStr := strings.Split(right, "|")[0]
	ourNumsStr := strings.Split(right, "|")[1]

	for _, v := range strings.Split(winningNumsStr, " ") {
		if v == "" {
			continue
		}
		num, err := strconv.Atoi(v)
		if err != nil {
			panic("")
		}
		nc.winningNums = append(nc.winningNums, num)
	}

	for _, v := range strings.Split(ourNumsStr, " ") {
		if v == "" {
			continue
		}
		num, err := strconv.Atoi(v)
		if err != nil {
			panic("")
		}
		nc.ourNums = append(nc.ourNums, num)
	}

	return &nc
}

func (c *Card) Value() int {
	toReturn := 0
	for _, v := range c.winningNums {
		if slices.Contains(c.ourNums, v) {
			if toReturn == 0 {
				toReturn = 1
			} else {
				toReturn *= 2
			}
		}
	}
	return toReturn
}

func (c *Card) NumMatches() int {
	toReturn := 0
	for _, v := range c.winningNums {
		if slices.Contains(c.ourNums, v) {
			toReturn += 1
		}
	}
	return toReturn
}

/***********************************************************************************
 * End of Card
 ***********************************************************************************/

func main() {
	// Boilerplate setup
	in, file := getInput()
	defer file.Close()

	cards := []Card{}
	cardsWon := make(map[int]int)

	// Part one ///////////////////////
	sum := 0
	lines := strings.Split(in, "\n")
	for _, v := range lines {
		c := NewCard(v)
		cards = append(cards, *c)
		cardsWon[c.id] = 1
		sum += c.Value()
	}

	fmt.Printf("sum part 01: %d\n", sum)

	// Part two ///////////////////////
	countOfCards := 0
	for _, c := range cards {
		for i := 0; i < cardsWon[c.id]; i++ {
			matches := c.NumMatches()
			for j := c.id + 1; j < (c.id+1)+matches; j++ {
				cardsWon[j] += 1
			}
		}
	}
	for k := range cardsWon {
		countOfCards += cardsWon[k]
	}
	fmt.Printf("count from part 02: %d\n", countOfCards)
}
