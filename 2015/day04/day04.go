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
func nonConcurrentSolution() {
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
	suffix       int
}

const numRoutines = 12

type AdventCoinMiner struct {
	answerChan  chan Answer
	mx          sync.RWMutex
	answer1     int
	answer2     int
	numResponses1 int
	numResponses2 int
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
			m.answerChan <- Answer{1, suffix} // this will be ignored in calculation of final answer
			m.mx.RLock()
			foundPart1 = true
		}
		if m.answer2 != 0 && m.answer2 < suffix {
			m.mx.RUnlock()
			m.answerChan <- Answer{2, suffix} // this will be ignored in calculation of final answer
			return
		}
		m.mx.RUnlock()

		h := calcMD5(suffix)
		if h[:5] == "00000" && !foundPart1 {
			m.answerChan <- Answer{1, suffix}
			foundPart1 = true // don't send more part 1 answers to channel
		}
		if h[:6] == "000000" {
			m.answerChan <- Answer{2, suffix}
			return
		}
		suffix += numRoutines
	}
}

func minButNotZero(x, y int) int {
	if x == 0 {
		return y
	}
	if y == 0 {
		return x
	}
	if x < y {
		return x
	}
	return y
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
				m.mx.Lock()
				m.answer1 = minButNotZero(m.answer1, answer.suffix)
				m.mx.Unlock()
				m.numResponses1++
			}
			if answer.part == 2 {
				m.mx.Lock()
				m.answer2 = minButNotZero(m.answer2, answer.suffix)
				m.mx.Unlock()
				m.numResponses2++
			}
		}

		if m.numResponses1 == numRoutines && m.numResponses2 == numRoutines {
			fmt.Println("Part 1 answer:", m.answer1)
			fmt.Println("Part 2 answer:", m.answer2)
			return
		}
	}
}

func main() {
	start := time.Now()
	nonConcurrentSolution()
	elapsed := time.Now().Sub(start)
	fmt.Printf("Non-concurrent solution took %v!\n\n", elapsed)

	start2 := time.Now()
	concurrentSolution()
	elapsed2 := time.Now().Sub(start2)
	fmt.Printf("Concurrent solution took %v!\n", elapsed2)
}
