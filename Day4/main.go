package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// alterar pra BoardPosition ou NumberPosition
type Board struct {
	Value  int
	Marked bool
}

type Bingo struct {
	// identifier
	id int
	// the boards in game
	Boards [][]Board
}

func getRandomNumbers(text string) []int {
	splitted := strings.Split(text, ",")

	var numbers []int
	for _, value := range splitted {
		getNumber, _ := strconv.Atoi(value)
		numbers = append(numbers, getNumber)
	}

	return numbers
}

func getLineValuesBoard(text string) []int {
	format := strings.ReplaceAll(text, " ", ",")
	fmt.Println("Format line number: ", format)
	splitted := strings.Split(format, ",")

	var numbers []int
	for _, value := range splitted {
		// skipping null values ",2,,0"
		if value == "" || value == " " {
			continue
		}
		getNumber, _ := strconv.Atoi(value)
		numbers = append(numbers, getNumber)
	}

	return numbers
}

func appendBoardValues(BingoSystem *[]Bingo, boards *[][]Board, boardId int) {
	bingoBoard := &Bingo{
		id:     boardId,
		Boards: *boards,
	}
	*BingoSystem = append(*BingoSystem, *bingoBoard)
	*boards = make([][]Board, 5)
	(*boards)[0] = make([]Board, 5)
}

func markNumberInBoard(boards *[][]Board, value int) {
	for x := 0; x < 5; x++ {
		for y := 0; y < 5; y++ {
			if (&(*boards)[x][y]).Value == value {
				(&(*boards)[x][y]).Marked = true
			}
		}
	}
}

func didBoardBingo(boards *[][]Board) bool {
	countPositionMarked := 0
	currentX := 0
	for x := 0; x < 5; x++ {
		for y := 0; y < 5; y++ {
			board := (*boards)[x][y]
			if currentX != x {
				countPositionMarked = 0
				currentX = x
			}
			if board.Marked {
				countPositionMarked += 1
			}

			if countPositionMarked == 5 {
				return true
			}
		}
	}

	currentY := 0
	countPositionMarked = 0
	for y := 0; y < 5; y++ {
		for x := 0; x < 5; x++ {
			board := (*boards)[x][y]
			if currentY != y {
				countPositionMarked = 0
				currentY = y
			}
			if board.Marked {
				countPositionMarked += 1
			}

			if countPositionMarked == 5 {
				return true
			}
		}
	}

	return false
}

func parseFile(fileName string) ([]Bingo, []int) {

	f, err := os.Open(fileName)
	if err != nil {
		fmt.Println("File not found")
		panic(err)
	}

	scanner := bufio.NewScanner(f)

	var randomNumbers []int
	var BingoSystem []Bingo

	var boards [][]Board

	boards = make([][]Board, 5)
	boards[0] = make([]Board, 5)

	isFirstLine := true

	var positionX int = 0
	var positionY int = 0
	var boardId int = 1

	for {
		if ok := scanner.Scan(); !ok {
			appendBoardValues(&BingoSystem, &boards, boardId)
			break
		}

		text := scanner.Text()
		fmt.Println("Text: ", text)

		if text == " " || text == "" || text == "\n" {
			if !isFirstLine {
				positionX = 0
				positionY = 0

				appendBoardValues(&BingoSystem, &boards, boardId)
				boards[positionX] = make([]Board, 5)
				boardId += 1

			} else {
				isFirstLine = false
			}

			continue
		}

		if isFirstLine {
			randomNumbers = getRandomNumbers(text)
			continue
		}

		lineNumbers := getLineValuesBoard(text)

		for _, value := range lineNumbers {
			board := &Board{
				Value:  value,
				Marked: false,
			}

			boards[positionX][positionY] = *board
			positionY += 1
		}

		positionX += 1

		if positionY == 5 && positionX < 5 {
			positionY = 0
			boards[positionX] = make([]Board, 5)
		}

	}

	fmt.Println("Bingo system: ", BingoSystem)

	return BingoSystem, randomNumbers

}

func initBingo(bingoSystem []Bingo, randomNumbers []int) (Bingo, int) {
	for _, number := range randomNumbers {
		fmt.Println("New number: ", number)
		for i := 0; i < len(bingoSystem); i++ {
			markNumberInBoard(&bingoSystem[i].Boards, number)
			winner := didBoardBingo(&bingoSystem[i].Boards)
			if winner {
				fmt.Println("BINGO!!")
				return bingoSystem[i], number
			}
		}
	}

	return Bingo{}, 0
}

func isBoardInWinners(id int, winners []int) bool {
	for _, value := range winners {
		if value == id {
			return true
		}
	}
	return false
}

func initBingoGetLastBoard(bingoSystem []Bingo, randomNumbers []int) (Bingo, int) {
	var boardsWon []int

	for _, number := range randomNumbers {
		fmt.Println("New number: ", number)
		for i := 0; i < len(bingoSystem); i++ {

			markNumberInBoard(&bingoSystem[i].Boards, number)
			winner := didBoardBingo(&bingoSystem[i].Boards)
			if winner {
				// fmt.Println("bingoSystem: ", bingoSystem[i].id, len(bingoSystem))

				if !isBoardInWinners(bingoSystem[i].id, boardsWon) {
					boardsWon = append(boardsWon, bingoSystem[i].id)
				}

				if len(bingoSystem) == len(boardsWon) {
					lastIndex := boardsWon[len(boardsWon)-1] - 1
					return bingoSystem[lastIndex], number
				}

			}
		}
	}

	return Bingo{}, 0
}

func makeScoreBoard(bingo Bingo, numberWinner int) {
	count := 0

	for x := 0; x < 5; x++ {
		for y := 0; y < 5; y++ {
			if bingo.Boards[x][y].Marked == false {
				count += bingo.Boards[x][y].Value
			}
		}
	}

	fmt.Println("Score: ", count*numberWinner)
}

func part1(fileName string) {
	bingoSystem, randomNumbers := parseFile(fileName)

	boardFound, numberWinner := initBingo(bingoSystem, randomNumbers)
	if len(boardFound.Boards) == 0 {
		fmt.Println("ERRO!")
		fmt.Println("Board found: ", boardFound, numberWinner)
		panic("Invalid input")
	}

	fmt.Println("boardFound: ", boardFound)

	makeScoreBoard(boardFound, numberWinner)
}

func part2(fileName string) {
	bingoSystem, randomNumbers := parseFile(fileName)

	boardFound, numberWinner := initBingoGetLastBoard(bingoSystem, randomNumbers)
	if len(boardFound.Boards) == 0 {
		fmt.Println("ERRO!")
		fmt.Println("Board found: ", boardFound, numberWinner)
		panic("Invalid input")
	}

	fmt.Println("boardFound: ", boardFound)

	makeScoreBoard(boardFound, numberWinner)
}

func main() {
	fileName := flag.String("file", "", "File to use")
	flag.Parse()
	if *fileName == "" {
		flag.PrintDefaults()
		os.Exit(1)
	}

	part1(*fileName)

	part2(*fileName)

}
