package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func getFileHandle() *os.File {
	filename := "in.dat"
	file, err := os.Open(filename)
	if err != nil {
		panic(fmt.Sprintf("err: error opening file %s\n", filename))
	}
	return file
	//defer file.Close()
}

func getMlsFromFile(file *os.File) string {
	var b strings.Builder
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fmt.Fprintf(&b, "%s\n", scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		panic(fmt.Sprintf("err: scanning file error\n"))
	}
	toReturn := cleanIn(b.String())
	return toReturn
}

func getInput() (string, *os.File) {
	file := getFileHandle()
	str := getMlsFromFile(file)
	return str, file
}

func cleanIn(s string) string {
	var toReturn string
	for _, v := range strings.Split(s, "\n") {
		if v != "" {
			s := fmt.Sprintf("%s\n", v)
			toReturn += strings.Trim(s, " \t")
		}
	}
	toReturn = strings.TrimSuffix(toReturn, "\n")
	return toReturn
}
