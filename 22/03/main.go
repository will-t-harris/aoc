package main

import (
	"bufio"
	"bytes"
	"io"
	"log"
	"os"
)

// split each row in half
// find the one character that is repeated in both halves
// a->z have priority 1->26
// A-Z have priorities 27 -> 52
// return sum of priorities

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

	for scanner.Scan() {
		lineLength := len(scanner.Text())

		firstHalf := scanner.Text()[lineLength/2:]
		secondHalf := scanner.Text()[:lineLength/2]

		for _, firstHalfChar := range firstHalf {
			for _, secondHalfChar := range secondHalf {

			}

		}
	}
}

type Priority int

const (
	_ Priority = iota
	a
	b
	c
	d
	e
	f
	g
	h
	i
	j
	k
	l
	m
	n
	o
	p
	q
	r
	s
	t
	u
	v
	w
	x
	y
	z
	A
	B
	C
	D
	E
	F
	G
	H
	I
	J
	K
	L
	M
	N
	O
	P
	Q
	R
	S
	T
	U
	V
	W
	X
	Y
	Z
)
