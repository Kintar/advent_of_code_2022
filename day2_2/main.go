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
	case "A":
		return rock
	case "B":
		return paper
	case "C":
		return scissors
	}
	panic(fmt.Sprintf("bad input: %s", s))
}

type outcome int

const (
	lose outcome = iota + 1
	draw
	win
)

func parseOutcome(s string) outcome {
	switch s {
	case "X":
		return lose
	case "Y":
		return draw
	case "Z":
		return win
	default:
		panic(fmt.Sprintf("what outcome is %s?", s))
	}
}

func chooseOutcome(them choice, result outcome) choice {
	if result == draw {
		return them
	}
	if result == win {
		switch them {
		case rock:
			return paper
		case paper:
			return scissors
		case scissors:
			return rock
		}
	}
	if result == lose {
		switch them {
		case rock:
			return scissors
		case paper:
			return rock
		case scissors:
			return paper
		}
	}
	panic("can't pick an outcome")
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
			result := parseOutcome(tokens[1])
			me := chooseOutcome(them, result)
			score += scoreOf(them, me)
		}
	}
	fmt.Println(score)
}
