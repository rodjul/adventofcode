package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func parseFile(fileName string) []int {
	var numbers []int

	f, err := os.Open(fileName)
	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		text := scanner.Text()

		splitted := strings.Split(text, ",")

		for _, value := range splitted {
			number, _ := strconv.Atoi(value)
			numbers = append(numbers, number)
		}
	}

	return numbers
}

func sum(values []int) int {
	sum := 0

	for _, value := range values {
		if value < 0 {
			sum += (value * -1) // converting negative to positive
			continue
		}

		sum += value
	}

	return sum
}

func sumGauss(n int) int {
	if n < 0 {
		n = n * -1
	}
	return (n * (n + 1)) / 2
}

func showCheapest(values map[int]string) {
	keys := make([]int, 0)
	for k, _ := range values {
		keys = append(keys, k)
	}
	sort.Ints(keys)
	for _, key := range keys {
		fmt.Println("Cheapest: ", key, values[key])
		break
	}
}

func part1(numbers []int, positions int) {
	cheapestPosition := make(map[int]string)

	for position := 1; position <= positions; position++ {
		fuel := 0
		tempNumbers := make([]int, len(numbers))
		copy(tempNumbers, numbers)

		for i, value := range tempNumbers {
			tempNumbers[i] = (value + (-position))
		}

		fuel = sum(tempNumbers)
		if cheapestPosition[fuel] == "" {
			cheapestPosition[fuel] = strconv.Itoa(position)
		} else {
			cheapestPosition[fuel] = cheapestPosition[fuel] + "," + strconv.Itoa(position)
		}
	}

	fmt.Println("Part 1:")
	showCheapest(cheapestPosition)
}

func part2(numbers []int, positions int) {
	cheapestPosition := make(map[int]string)
	for position := 1; position <= positions; position++ {
		fuel := 0
		tempNumbers := make([]int, len(numbers))
		copy(tempNumbers, numbers)

		for i, value := range tempNumbers {
			tempNumbers[i] = sumGauss(value + (-position))
		}

		fuel = sum(tempNumbers)
		if cheapestPosition[fuel] == "" {
			cheapestPosition[fuel] = strconv.Itoa(position)
		} else {
			cheapestPosition[fuel] = cheapestPosition[fuel] + "," + strconv.Itoa(position)
		}

	}

	fmt.Println("Part 2:")
	showCheapest(cheapestPosition)
}

func main() {
	fileName := flag.String("file", "", "File to use")
	MAX_POSITIONS := flag.Int("pos", 20, "Max positions")
	flag.Parse()
	if *fileName == "" {
		flag.PrintDefaults()
		os.Exit(1)
	}

	numbers := parseFile(*fileName)

	part1(numbers, *MAX_POSITIONS)

	part2(numbers, *MAX_POSITIONS)

}
