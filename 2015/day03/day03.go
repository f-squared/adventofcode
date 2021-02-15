package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

const FILENAME = "2015/day03/day-3-input.txt"

type Position struct {
	x int
	y int
}

func main() {
	data, err := ioutil.ReadFile(FILENAME)
	if err != nil {
		panic(err)
	}
	directions := strings.TrimSpace(string(data))

	// Part 1
	locations := make(map[Position]int)
	x := 0
	y := 0
	locations[Position{x, y}] = 1

	for i := 0; i < len(directions); i++ {
		direction := string(directions[i])

		if direction == "^" {
			x += 1
		} else if direction == "v" {
			x -= 1
		} else if direction == ">" {
			y += 1
		} else if direction == "<" {
			y -= 1
		} else {
			panic("Unexpected direction!")
		}

		pos := Position{x, y}
		if _, ok := locations[pos]; !ok {
			locations[pos] = 0
		}
		locations[pos] += 1
	}

	fmt.Printf("Santa delivers at least one present to %v houses.\n", len(locations))

	// Part 2
	locations = make(map[Position]int)
	santaX := 0
	santaY := 0
	roboX := 0
	roboY := 0
	locations[Position{0, 0}] = 2

	for i := 0; i < len(directions); i++ {
		direction := string(directions[i])

		if i%2 == 0 {
			x = santaX
			y = santaY
		} else {
			x = roboX
			y = roboY
		}

		if direction == "^" {
			x += 1
		} else if direction == "v" {
			x -= 1
		} else if direction == ">" {
			y += 1
		} else if direction == "<" {
			y -= 1
		} else {
			panic("Unexpected direction!")
		}

		pos := Position{x, y}
		if _, ok := locations[pos]; !ok {
			locations[pos] = 0
		}
		locations[pos] += 1

		if i%2 == 0 {
			santaX = x
			santaY = y
		} else {
			roboX = x
			roboY = y
		}
	}
	fmt.Printf("Santa and Robo Santa deliver at least one present to %v houses.\n", len(locations))
}
