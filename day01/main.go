package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func getCV(s string) int {
	var first int
	var second int

	for i := 0; i < len(s); i++ {
		val, err := strconv.Atoi(string(s[i]))
		if err == nil {
			first = val
			break
		}
	}
	for i := len(s) - 1; i >= 0; i-- {
		val, err := strconv.Atoi(string(s[i]))
		if err == nil {
			second = val
			break
		}
	}

	var b strings.Builder
	fmt.Fprintf(&b, "%d", first)
	fmt.Fprintf(&b, "%d", second)
	ns := b.String()
	val, err := strconv.Atoi(ns)
	if err != nil {
		panic("")
	}
	return val
}

func sumOfCVs(s string) int {
	split := strings.Split(s, "\n")
	total := 0
	for _, v := range split {
		total += getCV(v)
	}
	return total
}

func getMultiLineStringFromFile(file *os.File) string {
	var b strings.Builder
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fmt.Fprintf(&b, "%s\n", scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		panic("")
	}
	return b.String()
}

func main() {
	file, err := os.Open("in.dat")
	if err != nil {
		panic("")
	}
	defer file.Close()

	//fmt.Printf("allinput is: %s", allInput)
	mls := getMultiLineStringFromFile(file)
	sum := sumOfCVs(mls)
	fmt.Printf("The correct summed calibration is: %d\n", sum)
}

