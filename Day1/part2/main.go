package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const fileName = "input"

func alternative1() {
	/*
		{
			"A": sum([1, 2, 3, ...]),
			"B": sum([2, 3, 5, ...]),
		}
	*/
	f, err := os.Open(fileName)
	if err != nil {
		fmt.Println("File not found")
		panic(err)
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)

	dictLetters := make(map[string]int)

	for scanner.Scan() {
		lineSplitted := strings.Split(scanner.Text(), " ")

		number, _ := strconv.Atoi(lineSplitted[0])

		// fmt.Println("Number: ", number)
		for _, letter := range lineSplitted[1:] {
			if letter == " " || letter == "" {
				continue
			}
			val, exists := dictLetters[letter]
			if !exists {
				dictLetters[letter] = number
			} else {
				dictLetters[letter] = val + number
			}
		}
	}

	b, err := json.MarshalIndent(dictLetters, "", "  ")
	fmt.Println("Dict: ", string(b))

	var countIncreased int = 0
	var index int = 0
	var previousIndex int = 0

	for key, value := range dictLetters {
		fmt.Println("Checking: ", key, ":", value)
		index = value
		if previousIndex != 0 && index > previousIndex {
			countIncreased += 1
		}

		previousIndex = index
	}

	fmt.Println("Total increased: ", countIncreased)
}

func sumIncreased(tempArray []int, totalMeasurements *[]int, index *int, previousIndex *int, countIncreased *int) {
	var total int = 0
	for _, value := range tempArray {
		total += value
	}

	*totalMeasurements = append(*totalMeasurements, total)

	*index = (*totalMeasurements)[len((*totalMeasurements))-1]

	if (len(*totalMeasurements)-1) > 0 && *previousIndex != 0 && *index > *previousIndex {
		*countIncreased += 1
	}

	*previousIndex = *index
}

func main() {
	f, err := os.Open(fileName)
	if err != nil {
		fmt.Println("File not found")
		panic(err)
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)

	totalMeasurements := []int{}
	var tempArray []int
	var indexControl int = 0

	var countIncreased int = 0
	var index int = 0
	var previousIndex int = 0

	for scanner.Scan() {
		number, _ := strconv.Atoi(scanner.Text())

		if indexControl == 3 {
			sumIncreased(tempArray, &totalMeasurements, &index, &previousIndex, &countIncreased)
			tempArray = tempArray[1:]
			indexControl -= 1
		}
		tempArray = append(tempArray, number)
		indexControl += 1
	}

	// get the last values because the scanner found EOF
	sumIncreased(tempArray, &totalMeasurements, &index, &previousIndex, &countIncreased)

	fmt.Println("Total increased: ", countIncreased)
}
