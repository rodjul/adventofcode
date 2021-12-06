package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var positionX int = 0
var positionY int = 0

func part1(fileName string) {
	f, err := os.Open(fileName)
	if err != nil {
		fmt.Println("File not found")
		panic(err)
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		line := scanner.Text()

		lineSplited := strings.Split(line, " ")
		action := lineSplited[0]
		total, _ := strconv.Atoi(lineSplited[1])

		switch action {
		case "forward":
			positionX = positionX + total
			break
		case "down":
			positionY = positionY + total
			break
		case "up":
			positionY = positionY - total
			break
		default:
			break
		}
	}

	fmt.Println("Final X: ", positionX)
	fmt.Println("Final Y: ", positionY)
	fmt.Println("Final: ", positionY*positionX)
}

func part2(fileName string) {

	var positionAIM int = 0

	f, err := os.Open(fileName)
	if err != nil {
		fmt.Println("File not found")
		panic(err)
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		line := scanner.Text()

		lineSplited := strings.Split(line, " ")
		action := lineSplited[0]
		total, _ := strconv.Atoi(lineSplited[1])

		switch action {
		case "forward":
			positionX = positionX + total
			positionY = positionY + (positionAIM * total)
			break
		case "down":
			positionAIM = positionAIM + total
			break
		case "up":
			positionAIM = positionAIM - total
			break
		default:
			break
		}
	}

	fmt.Println("Final X: ", positionX)
	fmt.Println("Final Y: ", positionY)
	fmt.Println("Final depth: ", positionAIM)
	fmt.Println("Final: ", positionX*positionY)
}

func main() {
	fileName := flag.String("file", "", "File to use")
	flag.Parse()
	if *fileName == "" {
		flag.PrintDefaults()
		os.Exit(1)
	}
	// part1(fileName)
	part2(*fileName)
}
