package main

import (
	"fmt"
	"io/ioutil"
)

const FILENAME = "2015/day01/day-1-input.txt"

func main() {
	data, err := ioutil.ReadFile(FILENAME)
	if err != nil {
		panic(err)
	}

	directions := string(data)

	floor := 0
	basementPos := -1
	for i := 0; i < len(directions); i++ {
		direction := string(directions[i])

		if direction == "(" {
			floor += 1
		} else if direction == ")" {
			floor -= 1
		} else {
			panic("Unexpected direction!")
		}

		if basementPos < 0 {
			if floor < 0 {
				basementPos = i + 1
			}
		}
	}

	fmt.Printf("The directions take Santa to floor %d.\n", floor)
	fmt.Printf("Santa entered the basement with direction at position %d.\n", basementPos)
}
