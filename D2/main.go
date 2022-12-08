package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
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

func parseLine(line string) []int {
	splitLine := strings.Split(line, " ")
	switch splitLine[0] {
	case "forward":
		return []int{0, strToInt(splitLine[1])}
	case "up":
		return []int{-strToInt(splitLine[1]), 0}
	case "down":
		return []int{strToInt(splitLine[1]), 0}
	}
	return []int{0, 0}
}

func partOne() {
	position := []int{0, 0}
	lines := getFileLines("input")
	for i := 0; i < len(lines); i++ {
		position[0] += parseLine(lines[i])[0]
		position[1] += parseLine(lines[i])[1]
	}
	fmt.Println(position[0] * position[1])
}

// 0 = depth, 1 = forward

func partTwo() {
	position := []int{0, 0}
	aim := 0
	lines := getFileLines("input")
	for i := 0; i < len(lines); i++ {
		positionUpdate := parseLine(lines[i])
		position[1] += positionUpdate[1]
		aim += positionUpdate[0]
		position[0] += positionUpdate[1] * aim
	}
	fmt.Println(position[0] * position[1])
}

func main() {
	partTwo()
}
