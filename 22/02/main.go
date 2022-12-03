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
// X -> Rock
// Y -> Paper
// Z -> Scissors

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

	total := 0
	for scanner.Scan() {
		scanner.Bytes()

		round := strings.Fields(scanner.Text())

		opponentGuess := strings.TrimRight(round[0], "\n")
		myGuess := strings.TrimRight(round[1], "\n")

		if opponentGuess == "A" {
			if myGuess == "X" {
				total += 4
			} else if myGuess == "Y" {
				total += 8
			} else if myGuess == "Z" {
				total += 3
			}
		} else if opponentGuess == "B" {
			if myGuess == "X" {
				total += 1
			} else if myGuess == "Y" {
				total += 5
			} else if myGuess == "Z" {
				total += 9
			}
		} else if opponentGuess == "C" {
			if myGuess == "X" {
				total += 7
			} else if myGuess == "Y" {
				total += 2
			} else if myGuess == "Z" {
				total += 6
			}
		}
	}
	fmt.Println(total)
}

// func part1(reader io.Reader) string[] {
//   scanner := bufio.NewScanner(reader)
//   scanner.Split(bufio.ScanLines)
//
//   for scanner.Scan() {
//     scanner.Bytes()
//     fmt.Println(scanner.Text())
//   }
// }
