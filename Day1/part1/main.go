package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

const fileName = "input"

func main() {

	f, err := os.Open(fileName)
	if err != nil {
		fmt.Println("File not found")
		panic(err)
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)

	var countIncreased int = 0
	var index int = 0
	var previousIndex int = 0

	for scanner.Scan() {
		text := scanner.Text()
		index, _ = strconv.Atoi(text)

		if previousIndex != 0 && index > previousIndex {
			countIncreased += 1
		}

		previousIndex = index
	}

	fmt.Println("Total increased: ", countIncreased)

}
