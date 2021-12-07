package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
)

/*

each lanternfish creates a new lanternfish once every 7 days

When lanternfish reach 0 day, create a new lanternfish with timer 8
	- next day the timer 0 reset to 6

*/

func parseFile(fileName string) []int {
	var lanternfishes []int

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
			lanternfishes = append(lanternfishes, number)
		}
	}

	return lanternfishes
}

func main() {
	fileName := flag.String("file", "", "File to use")
	DAYS := flag.Int("days", 80, "File to use")
	flag.Parse()
	if *fileName == "" {
		flag.PrintDefaults()
		os.Exit(1)
	}

	lanternfishes := parseFile(*fileName)

	daysFishes := make([]int, 9)

	// REFERENCE: https://zonito.medium.com/lantern-fish-day-6-advent-of-code-2021-python-solution-4444387a8380

	for _, fish := range lanternfishes {
		daysFishes[fish] += 1
	}

	for i := 0; i < *DAYS; i++ {
		today := i % 9
		daysFishes[(today+7)%9] += daysFishes[today]
	}

	fmt.Println(daysFishes)
	count := 0

	for _, value := range daysFishes {
		count += value
	}

	fmt.Println(count)

}
