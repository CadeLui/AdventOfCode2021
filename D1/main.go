package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
)

func getFileLines(filename string) []string {
	var a []string
	data, error := os.ReadFile(filename)
	if error != nil {
		log.Fatal(error)
	}
	grabline := ""
	for i := 0; i < len(data); i++ {
		character := string(data[i])
		if character == "\n" {
			a = append(a, grabline)
			grabline = ""
			continue
		}
		grabline += character
	}
	a = append(a, grabline)
	return a
}

func strToInt(str string) int {
	integer, err := strconv.Atoi((str))
	if err != nil {
		log.Fatal(err)
	}
	return integer
}

func findDecreasesFromStrings(depths []string) int {
	decreases := 0
	for i := 1; i < len(depths); i++ {
		currentDepth := strToInt(depths[i])
		prevDepth := strToInt(depths[i-1])
		if currentDepth > prevDepth {
			decreases += 1
		}
	}
	return decreases
}

func findDecreasesFromInts(depths []int) int {
	decreases := 0
	for i := 1; i < len(depths); i++ {
		currentDepth := depths[i]
		prevDepth := depths[i-1]
		if currentDepth > prevDepth {
			decreases += 1
		}
	}
	return decreases
}

func getGroups(depths []string) []int {
	var groups []int
	for i := 0; i < len(depths)-2; i++ {
		one := strToInt(depths[i])
		two := strToInt(depths[i+1])
		three := strToInt(depths[i+2])
		groups = append(groups, one+two+three)
	}
	return groups
}

func main() {
	fmt.Println("Part One: ", findDecreasesFromStrings(getFileLines("input")))
	fmt.Println("Part Two: ", findDecreasesFromInts(getGroups(getFileLines("input"))))
}
