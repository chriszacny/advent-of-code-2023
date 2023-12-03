package main

import (
	"fmt"
)

/***********************************************************************************
 * Struct Name: aStruct
 *
 ***********************************************************************************/
type aStruct struct {
	red   int
	green int
	blue  int
}

/***********************************************************************************
 * End of aStruct
 ***********************************************************************************/

func main() {
	// Boilerplate setup
	in, file := getInput()
	defer file.Close()

	// Part one
	fmt.Println(in)

}
