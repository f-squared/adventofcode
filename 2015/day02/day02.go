package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

const FILENAME = "2015/day02/day-2-input.txt"

type Present struct {
	l int
	w int
	h int
}

func (p Present) wrappingPaper() int {
	areas := make([]int, 3)
	areas[0] = p.l * p.w
	areas[1] = p.w * p.h
	areas[2] = p.h * p.l

	minArea := areas[0]
	total := 0
	for _, area := range areas {
		if area < minArea {
			minArea = area
		}
		total += 2 * area
	}
	total += minArea

	return total
}

func (p Present) ribbon() int {
	perimeters := make([]int, 3)
	perimeters[0] = 2 * (p.l + p.w)
	perimeters[1] = 2 * (p.w + p.h)
	perimeters[2] = 2 * (p.h + p.l)

	minPerimeter := perimeters[0]
	for _, perimeter := range perimeters {
		if perimeter < minPerimeter {
			minPerimeter = perimeter
		}
	}

	return minPerimeter + p.l*p.w*p.h
}

func atoi(strNum string) int {
	num, err := strconv.Atoi(strNum)
	if err != nil {
		panic(err)
	}
	return num
}

func main() {
	data, err := ioutil.ReadFile(FILENAME)
	if err != nil {
		panic(err)
	}

	rawPresents := strings.Split(strings.TrimSpace(string(data)), "\n")
	presents := make([]Present, len(rawPresents))
	for i, rawPresent := range rawPresents {
		ds := strings.Split(rawPresent, "x")
		p := Present{
			atoi(ds[0]),
			atoi(ds[1]),
			atoi(ds[2]),
		}
		presents[i] = p
	}

	totalWrappingPaper := 0
	totalRibbon := 0
	for _, p := range presents {
		totalWrappingPaper += p.wrappingPaper()
		totalRibbon += p.ribbon()
	}
	fmt.Printf("The elves should order %d square feet of wrapping paper.\n", totalWrappingPaper)
	fmt.Printf("The elves should order %d square feet of ribbon.\n", totalRibbon)
}
