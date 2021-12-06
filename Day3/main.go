package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strconv"
)

// gamma rate   = most common number in each column from binary number
// epsilon rate = least common number in each column from binary number

/*

binary number length = 5
e.g. 				 = 00100
e.g. 				 = 11110
e.g. 				 = 10110
e.g. 				 = 10111

{
	0: {
		0: count(0),
		1: count(1),
	},
	1: {
		0: count(0),
		1: count(1),
	},
}

{
	0: {
		0: 1,
		1: 3,
	},
	1: {
		0: 3,
		1: 1,
	},
	2: {
		0: 0,
		1: 4
	}
}

*/

// const fileName = "input"

func findMostCommon(dict map[int]int) int {
	if dict[48] > dict[49] {
		return 0
	}
	return 1
}

func findLeastCommon(dict map[int]int) int {
	if dict[48] < dict[49] {
		return 0
	}
	return 1
}

func arrayIntToString(array []int) string {
	var number string
	for _, value := range array {
		number += strconv.Itoa(value)
	}
	return number
}

func part1(fileName string) {
	f, err := os.Open(fileName)
	if err != nil {
		fmt.Println("File not found")
		panic(err)
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)

	var mapBinary = make(map[int]map[int]int)

	var countLength = 0
	var foundLength = true

	for scanner.Scan() {
		text := scanner.Text()

		for i, _ := range text {
			number := int(text[i])
			if _, exists := mapBinary[i][number]; !exists {
				mapBinary[i] = make(map[int]int)
				mapBinary[i][48] = 0
				mapBinary[i][49] = 0
			}
			mapBinary[i][number] = mapBinary[i][number] + 1

			if foundLength {
				countLength += 1
			}
		}
		foundLength = false

	}

	fmt.Println("Length: ", countLength)
	fmt.Println(mapBinary)

	var common []int
	var least []int

	var i int = 0
	for i < countLength {
		common = append(common, findMostCommon(mapBinary[i]))
		i += 1
	}
	i = 0
	for i < countLength {
		least = append(least, findLeastCommon(mapBinary[i]))
		i += 1
	}

	fmt.Println("Common: ", common)
	fmt.Println("Least: ", least)

	gammaRate, _ := strconv.ParseInt(arrayIntToString(common), 2, 64)
	epsilonRate, _ := strconv.ParseInt(arrayIntToString(least), 2, 64)

	powerConsumption := gammaRate * epsilonRate
	fmt.Println("Power consumption: ", powerConsumption)
}

func getDomaintBit(input []string, index int, isBitDominant int) int {
	var countBits = make(map[string]int)
	countBits["0"] = 0
	countBits["1"] = 0

	for _, key := range input {
		countBits[string(key[index])] = countBits[string(key[index])] + 1
	}

	if countBits["0"] == countBits["1"] {
		return -1
	}

	if isBitDominant > 0 {
		if countBits["0"] > countBits["1"] {
			return 0
		}
		return 1
	}

	if countBits["0"] < countBits["1"] {
		return 0
	}
	return 1
}

func filterInput(input []string, bitDominant int) []string {

	var lengthText = len(input[0])

	var filter []string
	var tempFilter []string

	filter = input

	for len(filter) > 1 {
		var i = 0
		for i < lengthText {
			if i == 0 {
				i += 1
			}

			var numberDominant = getDomaintBit(filter, i, bitDominant)

			if numberDominant == -1 {
				numberDominant = bitDominant
			}

			for _, key := range filter {
				if string(key[i]) == strconv.Itoa(numberDominant) {
					tempFilter = append(tempFilter, key)
				}
			}

			if len(tempFilter) == 1 {
				return tempFilter
			}
			// fmt.Println("filter:", filter)
			// fmt.Println("tempFilter:", tempFilter)

			filter = tempFilter
			tempFilter = make([]string, 0)

			i += 1
		}

	}

	return tempFilter
}

func part2New(fileName string) {
	f, err := os.Open(fileName)
	if err != nil {
		fmt.Println("File not found")
		panic(err)
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)

	var mapBinary = make(map[string][]string)
	mapBinary["0"] = make([]string, 0)
	mapBinary["1"] = make([]string, 0)

	for scanner.Scan() {
		text := scanner.Text()

		if string(text[0]) == "0" {
			mapBinary["0"] = append(mapBinary["0"], text)
		} else if string(text[0]) == "1" {
			mapBinary["1"] = append(mapBinary["1"], text)
		}
	}

	var filter1 = filterInput(mapBinary["1"], 1)[0]
	var oxygenRating, _ = strconv.ParseInt(filter1, 2, 64)
	var filter2 = filterInput(mapBinary["0"], 0)[0]
	var co2Rating, _ = strconv.ParseInt(filter2, 2, 64)

	fmt.Println("Oxygen Rating value: ", oxygenRating)
	fmt.Println("CO2 Rating value: ", co2Rating)

	lifeSupportRating := co2Rating * oxygenRating

	fmt.Println("Life supporting rating: ", lifeSupportRating)

}

func main() {

	fileName := flag.String("file", "", "File to use")
	flag.Parse()
	if *fileName == "" {
		flag.PrintDefaults()
		os.Exit(1)
	}

	// part1(*fileName)
	part2New(*fileName)

}
