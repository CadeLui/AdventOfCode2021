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
	fixThisShit := stringArray
	var intArray []int
	for i := 0; i < len(fixThisShit); i++ {
		intArray = append(intArray, strToInt(fixThisShit[i]))
	}
	return intArray
}

func common(numbersAsStr []string) []int {
	var common []int
	for col := 0; col < len(numbersAsStr[0]); col++ {
		zeroes := 0
		ones := 1
		for row := 0; row < len(numbersAsStr); row++ {
			if strToInt(string(numbersAsStr[row][col])) == 1 {
				ones++
			} else {
				zeroes++
			}
		}
		if ones > zeroes {
			common = append(common, 1)
		} else {
			common = append(common, 0)
		}
	}
	return common
}

func remove(s []string, i int) []string {
    s[i] = s[len(s)-1]
    return s[:len(s)-1]
}

func commonWithReduction(numbersAsStr[]string) []int {
	common := numbersAsStr
	for col := 0; col < len(common[0]); col++ {
		zeroes := 0
		ones := 0
		for row := 0; row < len(common); row++ {
			if strToInt(string(common[row][col])) == 1 {
				ones++
			} else {
				zeroes++
			}
		}
		if ones >= zeroes {
			for row := 0; row < len(common); row++ {
				if common[row][col] != '1' {
					common = remove(common, row)
					row--
				}
			}
		} else {
			for row := 0; row < len(common); row++ {
				if common[row][col] != '0' {
					common = remove(common, row)
					row--
				}
			}
		}
	}
	var intCommon []int
	for col := 0; col < len(common[0]); col++ {
		if common[0][col] == '1' { intCommon = append(intCommon, 1) }
		if common[0][col] == '0' { intCommon = append(intCommon, 0) }
	}

	return intCommon
}

func uncommonWithReduction(numbersAsStr[]string) []int {
	uncommon := numbersAsStr
	for col := 0; col < len(uncommon[0]); col++ {
		zeroes := 0
		ones := 0
		if len(uncommon) == 1 { break }
		for row := 0; row < len(uncommon); row++ {
			if strToInt(string(uncommon[row][col])) == 1 {
				ones++
			} else {
				zeroes++
			}
		}
		if ones >= zeroes {
			for row := 0; row < len(uncommon); row++ {
				if uncommon[row][col] != '0' {
					uncommon = remove(uncommon, row)
					row--
				}
			}
		} else {
			for row := 0; row < len(uncommon); row++ {
				if uncommon[row][col] != '1' {
					uncommon = remove(uncommon, row)
					row--
				}
			}
		}
	}
	var intUncommon []int
	for col := 0; col < len(uncommon[0]); col++ {
		if uncommon[0][col] == '1' { intUncommon = append(intUncommon, 1) }
		if uncommon[0][col] == '0' { intUncommon = append(intUncommon, 0) }
	}

	return intUncommon
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
	lines := getFileLines("example")
	lines2 := getFileLines("example")
	common := common(lines)
	inverse := invertBin(common)
	fmt.Println(lines)
	fmt.Println(binToInt(common) * binToInt(inverse))
	fmt.Println(binToInt(commonWithReduction(lines)) * binToInt(uncommonWithReduction(lines2)))
	fmt.Println(lines)
}
