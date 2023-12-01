package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

var spellToDigitMap = make(map[string]int)

func init() {
	spellToDigitMap["one"] = 1
	spellToDigitMap["two"] = 2
	spellToDigitMap["three"] = 3
	spellToDigitMap["four"] = 4
	spellToDigitMap["five"] = 5
	spellToDigitMap["six"] = 6
	spellToDigitMap["seven"] = 7
	spellToDigitMap["eight"] = 8
	spellToDigitMap["nine"] = 9
	spellToDigitMap["zero"] = 0
}

func getCV(s string) int {
	if s == "" {
		return 0
	}

	results := make(map[int]int)
	// Get all single string digits into a map, with their indicies
	for i, v := range s {
		val, err := strconv.Atoi(string(v))
		if err == nil {
			results[i] = val
		}
	}

	// Get all spelled out digits into a map, with their indicies
	for k, v := range spellToDigitMap {
		i := strings.Index(s, k)
		if i != -1 {
			results[i] = v
		}
		i = strings.LastIndex(s, k)
		if i != -1 {
			results[i] = v
		}
	}

	// Sort the map keys, return the first value from first and last keys
	keys := []int{}
	for k := range results {
		keys = append(keys, k)
	}
	sort.Ints(keys)
	first := results[keys[0]]
	second := results[keys[len(keys)-1]]

	// Prepare and return results
	var b strings.Builder
	fmt.Fprintf(&b, "%d", first)
	fmt.Fprintf(&b, "%d", second)
	ns := b.String()
	val, err := strconv.Atoi(ns)
	if err != nil {
		panic(fmt.Sprintf("err: could not convert %v to an int", ns))
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
		panic(fmt.Sprintf("err: scanning file error\n"))
	}
	return b.String()
}

func main() {
	filename := "in.dat"
	file, err := os.Open(filename)
	if err != nil {
		panic(fmt.Sprintf("err: error opening file %s\n", filename))
	}
	defer file.Close()

	mls := getMultiLineStringFromFile(file)
	sum := sumOfCVs(mls)
	fmt.Printf("The correct summed calibration is: %d\n", sum)
}
