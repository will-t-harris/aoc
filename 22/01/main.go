package main

import (
	"bufio"
	"bytes"
	"io"
	"log"
	"os"
	"strconv"

	"github.com/samber/lo"
	"golang.org/x/exp/slices"
)

func main() {
  file, err := os.Open("input.txt")

  if err != nil {
    log.Fatal(err)
  }

  defer file.Close()

  var buffer bytes.Buffer
  tee := io.TeeReader(file, &buffer)

  part1Answer, err := part1(tee)
  if err != nil {
    log.Fatal(err)
  }

  log.Println(part1Answer)

  part2Answer, err := part2(&buffer)
  if err != nil {
    log.Fatal(err)
  }

  log.Println(part2Answer)
}

func part1(reader io.Reader) (int, error) {
  calories := getCalories(reader)
  return lo.Max(calories), nil
}

func part2(reader io.Reader) (int, error) {
  calories := getCalories(reader)

  slices.Sort(calories)

  total := 0

  for _, i := range calories[len(calories)-3:] {
    total += i
  }

  return total, nil
}

func getCalories(reader io.Reader) []int {
  scanner := bufio.NewScanner(reader)
  scanner.Split(bufio.ScanLines)

  var (
    total int
    calories []int
  )

  for scanner.Scan() {
    scanner.Bytes()
    text := scanner.Text()

    if text == "" {
      calories = append(calories, total)
      total = 0
      continue
    }

    calories, error := strconv.Atoi(text)
    if error != nil {
      continue
    }

    total += calories
  }

  calories = append(calories, total)
  
  return calories
}
