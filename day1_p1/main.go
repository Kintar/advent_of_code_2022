package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
)

func main() {
	f, err := os.Open("data.txt")
	if err != nil {
		panic(err)
	}

	max := 0
	current := 0
	reader := bufio.NewReader(f)
	done := false
	for !done {
		l, _, err := reader.ReadLine()

		if err == io.EOF {
			done = true
		}

		if len(l) == 0 {
			if current > max {
				max = current
			}
			current = 0
			continue
		}

		val, err := strconv.ParseInt(string(l), 10, 64)
		if err != nil {
			panic(err)
		}
		current += int(val)
	}

	fmt.Println(max)
}
