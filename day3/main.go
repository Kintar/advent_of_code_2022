package main

import (
	"bufio"
	_ "embed"
	"fmt"
	"io"
	"strings"
)

//go:embed data.txt
var data string

func priority(b byte) int {
	if b > 96 {
		return int(b - 96)
	}
	return int(b - 64)
}

type empty struct{}

func main() {
	reader := bufio.NewReader(strings.NewReader(data))
	done := false
	dupes := make(map[byte]empty)
	for !done {
		l, _, err := reader.ReadLine()
		if err == io.EOF {
			done = true
		}
		count := len(l)
		if count == 0 {
			continue
		}
		p1, p2 := l[0:count/2], l[count/2:]
		for _, v := range p1 {
			for _, vv := range p2 {
				if v == vv {
					dupes[v] = empty{}
				}
			}
		}
	}
	score := 0
	for k := range dupes {
		score += priority(k)
	}
	fmt.Println(score)
}
