package main

import (
	"bufio"
	_ "embed"
	"fmt"
	"io"
	"strings"
)

type choice string

const (
	rock     choice = "rock"
	paper           = "paper"
	scissors        = "scissors"
)

func valueOf(c choice) int {
	switch c {
	case rock:
		return 1
	case paper:
		return 2
	case scissors:
		return 3
	default:
		panic(fmt.Sprintf("wtf is choice %v?", c))
	}
}

func scoreOf(them, me choice) int {
	score := valueOf(me)
	if them == me {
		score += 3
	}
	if beats(me, them) {
		score += 6
	}
	return score
}

func beats(a, b choice) bool {
	if a == rock {
		return b == scissors
	}
	if a == paper {
		return b == rock
	}
	if a == scissors {
		return b == paper
	}
	return false
}

func parseChoice(s string) choice {
	switch s {
	case "A", "X":
		return rock
	case "B", "Y":
		return paper
	case "C", "Z":
		return scissors
	}
	panic(fmt.Sprintf("bad input: %s", s))
}

//go:embed data.txt
var data string

func main() {
	reader := bufio.NewReader(strings.NewReader(data))
	done := false
	score := 0
	for !done {
		l, _, err := reader.ReadLine()
		if err == io.EOF {
			done = true
		}
		if len(l) == 0 {
			continue
		} else {
			str := string(l)
			tokens := strings.Split(str, " ")
			them := parseChoice(tokens[0])
			me := parseChoice(tokens[1])
			score += scoreOf(them, me)
		}
	}
	fmt.Println(score)
}
