package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"sort"
	"strconv"
)

func main() {
	f, err := os.Open("data.txt")
	if err != nil {
		panic(err)
	}

	elves := make([]int, 0)

	reader := bufio.NewReader(f)
	done := false
	current := 0
	for !done {
		l, _, err := reader.ReadLine()

		if err == io.EOF {
			done = true
		}

		if len(l) == 0 {
			elves = append(elves, current)
			current = 0
		} else {
			val, err := strconv.ParseInt(string(l), 10, 64)
			if err != nil {
				panic(err)
			}
			current += int(val)
		}

	}

	sort.Ints(elves)
	elves = elves[len(elves)-3:]
	if len(elves) != 3 {
		panic("wrong number of results!")
	}
	total := 0
	for _, e := range elves {
		total += e
	}
	fmt.Println(total)
}
