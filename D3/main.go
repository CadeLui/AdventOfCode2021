package main

import (
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

func getFileLines(filename string) []string {
	data, error := os.ReadFile(filename)
	if error != nil {
		log.Fatal(error)
	}
	dataArray := strings.Split(string(data), "\n")
	return dataArray
}

func strToInt(str string) int {
	integer, err := strconv.Atoi((str))
	if err != nil {
		log.Fatal(err)
	}
	return integer
}

func convertStringArrayToIntArray(stringArray []string) []int {
	var intArray []int
	for i := 0; i < len(stringArray); i++ {
		intArray = append(intArray, strToInt(stringArray[i]))
	}
	return intArray
}

func common(numbersAsStr []string) []int {
	var common []int
	for col := 0; col < len(numbersAsStr[0]); col++ {
		// [1, 0]
		pair := []int{0, 0}
		for row := 0; row < len(numbersAsStr); row++ {
			if strToInt(string(numbersAsStr[row][col])) == 1 {
				pair[0]++
			} else {
				pair[1]++
			}
		}
		if pair[0] > pair[1] {
			common = append(common, 1)
		} else {
			common = append(common, 0)
		}
	}
	return common
}

func commonButWeird(numbersAsStr[]string) []int {
	var common []int
}

func invertBin(bin []int) []int {
	var inverse []int
	for i := 0; i < len(bin); i++ {
		switch bin[i] {
		case 1:
			inverse = append(inverse, 0)
		case 0:
			inverse = append(inverse, 1)
		}
	}
	return inverse
}

func binToInt(bin []int) int {
	num := 0
	for i := 0; i < len(bin); i++ {
		num += bin[i] * int(math.Pow(2, float64(len(bin)-i-1)))
	}
	return num
}

func main() {
	lines := getFileLines("input")
	common := common(lines)
	inverse := invertBin(common)
	fmt.Print(binToInt(common) * binToInt(inverse))
}
