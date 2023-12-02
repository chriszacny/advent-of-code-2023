package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

/***********************************************************************************
 * Struct Name: gameRound
 *
 * Description:
 *   todo
 *
 * Usage:
 *   todo
 *
 * Notes:
 *   - todo
 *
 ***********************************************************************************/
type gameRound struct {
	red   int
	green int
	blue  int
}

/***********************************************************************************
 * End of gameRound
 ***********************************************************************************/

/***********************************************************************************
 * Struct Name: gameRecord
 *
 * Description:
 *   todo
 *
 * Usage:
 *   todo
 *
 * Notes:
 *   - todo
 *
 ***********************************************************************************/
type gameRecord struct {
	id     int
	rounds []gameRound
}

func (g *gameRecord) String() string {
	return fmt.Sprintf("%#v", g)
}

/***********************************************************************************
 * End of gameRecord
 ***********************************************************************************/

func gamesPossible(puzzleInput string, red int, green int, blue int) []gameRecord {
	toReturn := []gameRecord{}
	gamesStr := strings.Split(puzzleInput, "\n")
	for _, line := range gamesStr {
		gr := parseGamesFileLine(line)
		possible := true
		for _, round := range gr.rounds {
			if round.red > red || round.green > green || round.blue > blue {
				possible = false
				break
			}
		}
		if possible == true {
			toReturn = append(toReturn, gr)
		}
	}
	return toReturn
}

func parseGamesFileLine(line string) gameRecord {
	idstr := strings.Split(line, ":")[0]
	id, err := strconv.Atoi(idstr)
	if err != nil {
		fmt.Println(idstr)
		panic("todo")
	}
	// 1:2r,2g;6r,3g;2r,1g,2b;1r
	// 2r,2g;6r,3g;2r,1g,2b;1r
	roundsstr := strings.Split(line, ":")[1]
	rounds := strings.Split(roundsstr, ";")

	gr := gameRecord{}
	gr.id = id

	for _, v := range rounds {
		// ["2r,2g" ...]
		round := gameRound{}
		colorsStr := strings.Split(v, ",")
		for _, c := range colorsStr {
			// "2r"
			colorCountStr := c[0 : len(c)-1]
			colorCount, err := strconv.Atoi(colorCountStr)
			if err != nil {
				panic("todo")
			}
			colorValue := string(c[len(c)-1])

			if colorValue == "r" {
				round.red = colorCount
			} else if colorValue == "b" {
				round.blue = colorCount
			} else {
				round.green = colorCount
			}
		}
		gr.rounds = append(gr.rounds, round)
	}

	return gr
}

func getMultiLineStringFromFile(file *os.File) string {
	var b strings.Builder
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fmt.Fprintf(&b, "%s\n", scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		panic(fmt.Sprintf("err: scanning file error\n"))
	}
	toReturn := b.String()
	toReturn = strings.TrimSuffix(toReturn, "\n")
	return toReturn
}

func main() {
	filename := "in.dat"
	file, err := os.Open(filename)
	if err != nil {
		panic(fmt.Sprintf("err: error opening file %s\n", filename))
	}
	defer file.Close()

	mls := getMultiLineStringFromFile(file)
	gp := gamesPossible(mls, 12, 13, 14)
	sum := 0
	for _, v := range gp {
		sum += v.id
	}
	fmt.Printf("sum is: %d\n", sum)
}
