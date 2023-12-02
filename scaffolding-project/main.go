package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

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
	fmt.Printf("%s\n", mls)
}
