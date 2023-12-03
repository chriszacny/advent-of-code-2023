package main

import (
	"fmt"
	"strconv"
	"strings"
)

type CellValueType int

const (
	Number CellValueType = iota + 1
	Symbol
	Other
)

/***********************************************************************************
 * Struct Name: Cell
 *
 ***********************************************************************************/
type Cell struct {
	id        int
	value     string
	valueType CellValueType
	edges     []Cell
}

/***********************************************************************************
 * End of Cell
 ***********************************************************************************/

/***********************************************************************************
 * Struct Name: NumberCandidate
 *
 ***********************************************************************************/
type NumberCandidate struct {
	value []Cell
}

func (n *NumberCandidate) Int() int {
	num := ""
	for _, n := range n.value {
		num += n.value
	}
	fullint, err := strconv.Atoi(num)
	//fmt.Printf("%d\n", fullint)
	if err != nil {
		panic("")
	}
	return fullint
}

/***********************************************************************************
 * End of NumberCandidate
 ***********************************************************************************/

/***********************************************************************************
 * Struct Name: GearData
 *
 ***********************************************************************************/
type GearData struct {
	value  Cell
	numOne int
	numTwo int
}

/***********************************************************************************
 * End of GearData
 ***********************************************************************************/

func getValueType(s string) CellValueType {
	// check if it is a number or a period
	// else it is a symbol
	_, err := strconv.Atoi(s)
	if err != nil {
		if s == "." {
			return Other
		} else {
			return Symbol
		}
	}
	return Number
}

func getNumberCands(matrix [][]Cell) []NumberCandidate {
	cands := []NumberCandidate{}
	for i := 0; i < len(matrix); i++ {
		j := 0
		for j < len(matrix[i]) {
			cand := NumberCandidate{}
			foundNumber := true
			for k := j; foundNumber == true && k < len(matrix[i]); k++ {
				if matrix[i][k].valueType == Number {
					cand.value = append(cand.value, matrix[i][k])
				} else {
					foundNumber = false
					j = k
				}
			}
			if len(cand.value) > 0 {
				cands = append(cands, cand)
			}
			j++
		}
	}
	return cands
}

func getGearData(matrix [][]Cell) []GearData {
	data := []GearData{}
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			if matrix[i][j].value == "*" {
				// we now need to determine if it is bordered by two numbers AND only two numbers
			}
		}
	}
	return data
}

func main() {
	// Boilerplate setup
	in, file := getInput()
	defer file.Close()

	// Part one ///////////////////////
	lines := strings.Split(in, "\n")

	// build the matrix
	rows := len(lines)
	cols := len(lines[0])
	matrix := make([][]Cell, rows)
	for i := range matrix {
		matrix[i] = make([]Cell, cols)
	}

	// Setup values
	idct := 1
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			c := Cell{}
			c.id = idct
			c.value = string(lines[i][j])
			c.valueType = getValueType(c.value)
			matrix[i][j] = c
			idct++
		}
	}

	// Assign Edges
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			// Top Left
			if i-1 >= 0 && j-1 >= 0 {
				matrix[i][j].edges = append(matrix[i][j].edges, matrix[i-1][j-1])
			}

			// Top
			if i-1 >= 0 {
				matrix[i][j].edges = append(matrix[i][j].edges, matrix[i-1][j])
			}

			// Top Right
			if i-1 >= 0 && j+1 < cols {
				matrix[i][j].edges = append(matrix[i][j].edges, matrix[i-1][j+1])
			}

			// Left
			if j-1 >= 0 {
				matrix[i][j].edges = append(matrix[i][j].edges, matrix[i][j-1])
			}

			// Right
			if j+1 < cols {
				matrix[i][j].edges = append(matrix[i][j].edges, matrix[i][j+1])
			}

			// Bottom Left
			if i+1 < rows && j-1 >= 0 {
				matrix[i][j].edges = append(matrix[i][j].edges, matrix[i+1][j-1])
			}

			// Bottom
			if i+1 < rows {
				matrix[i][j].edges = append(matrix[i][j].edges, matrix[i+1][j])
			}

			// Right
			if i+1 < rows && j+1 < cols {
				matrix[i][j].edges = append(matrix[i][j].edges, matrix[i+1][j+1])
			}
		}
	}

	numberCands := getNumberCands(matrix)

	validNumbers := []NumberCandidate{}
	for _, c := range numberCands {
		isValid := false
		for _, v := range c.value {
			for _, e := range v.edges {
				if e.valueType == Symbol {
					isValid = true
					break
				}
			}
			if isValid {
				break
			}
		}
		if isValid {
			validNumbers = append(validNumbers, c)
		}
	}

	sum := 0

	for _, vn := range validNumbers {
		sum += vn.Int()
	}

	fmt.Printf("sum of part 1 is: %d\n", sum)

	// Part two ///////////////////////

	gearPositions := make(map[int][]NumberCandidate)

	for _, n := range validNumbers {
		found := false
		for _, v := range n.value {
			for _, e := range v.edges {
				if e.valueType == Symbol && e.value == "*" {
					found = true
					gearPositions[e.id] = append(gearPositions[e.id], n)
					break
				}
			}
			if found {
				break
			}
		}
	}

	sumGearRatios := 0
	for k := range gearPositions {
		if len(gearPositions[k]) == 2 {
			fmt.Printf("%d\n", gearPositions[k][0].Int())
			fmt.Printf("%d\n", gearPositions[k][1].Int())
			gearRatio := gearPositions[k][0].Int() * gearPositions[k][1].Int()
			sumGearRatios += gearRatio
		}
	}
	fmt.Printf("sum of part 2 is: %d\n", sumGearRatios)
}
