package main

import (
	"bufio"
	"fmt"
	"os"
	"time"
)

func main() {
	f, err := os.Open("input")
	if err != nil {
		panic(err)
	}
	defer func() { _ = f.Close() }()
	reader := bufio.NewReader(f)
	start := time.Now()

	var line []byte
	sum := 0
	badgeSum := 0
	badgeLine := 0
	badgeData := [3][]byte{}
	for line, _, err = reader.ReadLine(); err == nil; line, _, err = reader.ReadLine() {
		common := priorityOf(findAllCommon(findHalves(line)))
		badgeData[badgeLine] = line
		badgeLine++
		if badgeLine == 3 {
			badgeSum += int(priorityOf(findCommon(badgeData)))
			badgeLine = 0
		}
		sum += int(common)
	}

	end := time.Now()
	fmt.Printf("\nDone in %.3fus\n", float64(end.Sub(start).Nanoseconds())/1000.0)
	fmt.Println("\nSum of priorities is", sum)
	fmt.Println("Sum of badges is", badgeSum)
}

func findHalves(line []byte) [2][]byte {
	result := [2][]byte{}
	l := len(line)
	if l%2 != 0 {
		panic(fmt.Sprintf("length %d is not even", l))
	}
	result[0] = line[0 : l/2]
	result[1] = line[l/2:]
	return result
}

func findAllCommon(compartments [2][]byte) uint8 {
	for _, c := range compartments[0] {
		for _, cc := range compartments[1] {
			if c == cc {
				return c
			}
		}
	}

	return 0
}

const ucasePriority uint8 = 65
const lcasePriority uint8 = 96

func priorityOf(char uint8) uint8 {
	if char >= lcasePriority {
		return char - lcasePriority
	}
	return 27 + char - ucasePriority
}

func findCommon(lines [3][]byte) byte {
	for _, c := range lines[0] {
		for _, d := range lines[1] {
			if c == d {
				for _, e := range lines[2] {
					if d == e {
						return e
					}
				}
			}
		}
	}
	return 0
}
