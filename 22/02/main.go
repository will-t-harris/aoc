package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

//// Column 1
// A -> Rock
// B -> Paper
// C -> Scissors

//// Column 2
// X -> Rock --> I need to lose
// Y -> Paper --> I need to draw
// Z -> Scissors --> I need to win

//// Shape points
// Rock -> 1 point
// Paper -> 2 points
// Scissors -> 3 points

//// Outcome points
// Loss -> 0 points
// Draw -> 3 points
// Win -> 6 points

func main() {
	file, err := os.Open("input.txt")

	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	var buffer bytes.Buffer
	teeReader := io.TeeReader(file, &buffer)

	scanner := bufio.NewScanner(teeReader)
	scanner.Split(bufio.ScanLines)

	part1Total := 0
	part2Total := 0

	for scanner.Scan() {
		round := strings.Fields(scanner.Text())

		opponentGuess := strings.TrimRight(round[0], "\n")
		myGuess := strings.TrimRight(round[1], "\n")

		part1Total += calculateScore(myGuess, opponentGuess)

		if opponentGuess == "A" {
			if myGuess == "X" {
				myGuess = "Z"
			} else if myGuess == "Y" {
				myGuess = "X"
			} else {
				myGuess = "Y"
			}
		} else if opponentGuess == "B" {
			if myGuess == "X" {
				myGuess = "X"
			} else if myGuess == "Y" {
				myGuess = "Y"
			} else {
				myGuess = "Z"
			}
		} else {
			if myGuess == "X" {
				myGuess = "Y"
			} else if myGuess == "Y" {
				myGuess = "Z"
			} else {
				myGuess = "X"
			}
		}

		part2Total += calculateScore(myGuess, opponentGuess)
	}
	fmt.Println("Part 1: ", part1Total)
	fmt.Println("Part 2: ", part2Total)
}

func calculateScore(myGuess string, opponentGuess string) int {
	if opponentGuess == "A" {
		if myGuess == "X" {
			return 4
		} else if myGuess == "Y" {
			return 8
		} else {
			return 3
		}
	} else if opponentGuess == "B" {
		if myGuess == "X" {
			return 1
		} else if myGuess == "Y" {
			return 5
		} else {
			return 9
		}
	} else {
		if myGuess == "X" {
			return 7
		} else if myGuess == "Y" {
			return 2
		} else {
			return 6
		}
	}
}
