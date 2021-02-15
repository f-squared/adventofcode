package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"strconv"
	"strings"
	"sync"
	"time"
)

const secretKey = "iwrupvqb"

func calcMD5(suffix int) string {
	text := strings.Join([]string{secretKey, strconv.Itoa(suffix)}, "")
	h := md5.Sum([]byte(text))
	return hex.EncodeToString(h[:])
}

// Takes about 3-4 seconds on my computer.
func naiveSolution() {
	i := 0
	part1Found := false
	for {
		h := calcMD5(i)
		if h[:5] == "00000" && !part1Found {
			fmt.Println("Part 1 answer:", i)
			part1Found = true
		}
		if h[:6] == "000000" {
			fmt.Println("Part 2 answer:", i)
			return
		}
		i++
	}
}

type Answer struct {
	part         int
	goroutineNum int
	suffix       int
}

const numRoutines = 15

type AdventCoinMiner struct {
	answerChan  chan Answer
	mx          sync.RWMutex
	answer1     int
	answer2     int
	candidates1 [numRoutines]int
	candidates2 [numRoutines]int
}

// Search for parts 1 and 2 answer matches.
// Stop search for each part once that part's answer is found.
func (m *AdventCoinMiner) crunch(num int) {
	suffix := num
	foundPart1 := false
	for {
		m.mx.RLock()
		// If another crunch goroutine has found a more promising answer,
		// then this goroutine's work is done for that part.
		if m.answer1 != 0 && m.answer1 < suffix && !foundPart1 {
			m.mx.RUnlock()
			m.answerChan <- Answer{1, num, suffix} // this will be ignored in calculation of final answer
			m.mx.RLock()
			foundPart1 = true
		}
		if m.answer2 != 0 && m.answer2 < suffix {
			m.mx.RUnlock()
			m.answerChan <- Answer{2, num, suffix} // this will be ignored in calculation of final answer
			return
		}
		m.mx.RUnlock()

		h := calcMD5(suffix)
		if h[:5] == "00000" && !foundPart1 {
			m.answerChan <- Answer{1, num, suffix}
			foundPart1 = true // don't send more part 1 answers to channel
		}
		if h[:6] == "000000" {
			m.answerChan <- Answer{2, num, suffix}
			return
		}
		suffix += numRoutines
	}
}

func minButNotZero(candidates [numRoutines]int) int {
	minAnswer := 0
	for _, v := range candidates {
		if minAnswer == 0 || (v != 0 && v < minAnswer) {
			minAnswer = v
		}
	}
	if minAnswer == 0 {
		panic("Unexpected zero candidate!")
	}
	return minAnswer
}

func numZeroes(candidates [numRoutines]int) int {
	count := 0
	for _, v := range candidates {
		if v == 0 {
			count++
		}
	}
	return count
}

// Well that was gnarly. 
// Not really a big time saver, but a concurrency brain teaser.
// Takes about 2-3 seconds on my computer.
func concurrentSolution() {
	m := AdventCoinMiner{
		answerChan: make(chan Answer),
	}

	for i := 0; i < numRoutines; i++ {
		go m.crunch(i)
	}

	for {
		select {
		case answer := <-m.answerChan:
			if answer.part == 1 {
				m.candidates1[answer.goroutineNum] = answer.suffix
				m.mx.Lock()
				m.answer1 = minButNotZero(m.candidates1)
				m.mx.Unlock()
			}
			if answer.part == 2 {
				m.candidates2[answer.goroutineNum] = answer.suffix
				m.mx.Lock()
				m.answer2 = minButNotZero(m.candidates2)
				m.mx.Unlock()
			}
		}

		if numZeroes(m.candidates1) == 0 && numZeroes(m.candidates2) == 0 {
			fmt.Println("Part 1 answer:", m.answer1)
			fmt.Println("Part 2 answer:", m.answer2)
			return
		}
	}
}

func main() {
	start := time.Now()
	naiveSolution()
	elapsed := time.Now().Sub(start)
	fmt.Printf("Non-concurrent solution took %v!\n\n", elapsed)

	start2 := time.Now()
	concurrentSolution()
	elapsed2 := time.Now().Sub(start2)
	fmt.Printf("Concurrent solution took %v!\n", elapsed2)
}
