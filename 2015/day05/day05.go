package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

const FILENAME = "2015/day05/day-5-input.txt"

type TestCase struct {
	word   string
	isNice bool
}

func isNicePt1(word string) bool {
	numVowels := 0
	hasDoubledLetter := false
	prevLetter := ""

	for _, letter := range strings.Split(word, "") {
		if strings.Contains("aeiou", letter) {
			numVowels++
		}
		if letter == prevLetter {
			hasDoubledLetter = true
		}
		prevLetter = letter
	}

	// must contain at least 3 vowels
	if numVowels < 3 {
		return false
	}

	// must contain one letter that appears twice in a row
	if !hasDoubledLetter {
		return false
	}

	// must not contain the strings "ab", "cd", "pq", or "xy"
	naughtySubstrings := []string{"ab", "cd", "pq", "xy"}
	for _, substring := range naughtySubstrings {
		if strings.Contains(word, substring) {
			return false
		}
	}

	return true
}

func isNicePt1Tests() {
	tcs := []TestCase{
		{"ugknbfddgicrmopn", true},
		{"aaa", true},
		{"jchzalrnumimnmhp", false},
		{"haegwjzuvuyypxyu", false},
		{"dvszwmarrgswjxmb", false},
	}

	for _, tc := range tcs {
		if isNicePt1(tc.word) != tc.isNice {
			fmt.Printf(
				"Test case failed! Expected: isNicePt1(%v) to be %v", tc.word, tc.isNice)
		}
	}
}

func isNicePt2(word string) bool {
	pairsToIndices := make(map[string][]int)
	triplets := make([]string, 0, len(word)-2)

	for i := 0; i < len(word)-1; i++ {
		pair := word[i : i+2]
		if _, ok := pairsToIndices[pair]; !ok {
			pairsToIndices[pair] = make([]int, 0)
		}
		pairsToIndices[pair] = append(pairsToIndices[pair], i)

		if i < len(word)-2 {
			triplets = append(triplets, word[i:i+3])
		}
	}

	// must contain pair of two letters that appear at least twice without overlapping
	var foundDoublePair bool
	for _, indices := range pairsToIndices {
		if len(indices) > 2 {
			foundDoublePair = true
			break
		} else if len(indices) == 2 {
			if indices[1]-indices[0] > 1 {
				foundDoublePair = true
			}
			break
		}
	}
	if !foundDoublePair {
		return false
	}

	// must contain at least one letter which repeats with exactly one letter in between
	for _, triplet := range triplets {
		if triplet[0] == triplet[2] {
			return true
		}
	}
	return false
}

func isNicePt2Tests() {
	tcs := []TestCase{
		{"qjhvhtzxzqqjkmpb", true},
		{"xxyxx", true},
		{"aaa", false},
		{"uurcxstgmygtbstg", false},
		{"ieodomkazucvgmuy", false},
	}

	for _, tc := range tcs {
		if isNicePt2(tc.word) != tc.isNice {
			fmt.Printf(
				"Test case failed! Expected: isNicePt2(%v) to be %v", tc.word, tc.isNice)
		}
	}
}

func main() {
	isNicePt1Tests()
	isNicePt2Tests()

	data, err := ioutil.ReadFile(FILENAME)
	if err != nil {
		panic(err)
	}
	words := strings.Split(strings.TrimSpace(string(data)), "\n")

	answer1 := 0
	answer2 := 0
	for _, word := range words {
		if isNicePt1(word) {
			answer1++
		}
		if isNicePt2(word) {
			answer2++
		}
	}
	fmt.Println(answer1, "strings are nice under Part 1 rules.")
	fmt.Println(answer2, "strings are nice under Part 2 rules.")
}
