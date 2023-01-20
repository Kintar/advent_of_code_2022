package util

import (
	"bufio"
	"fmt"
	"os"
	"time"
)

func MustOpen(file string) *os.File {
	f, err := os.Open(file)
	if err != nil {
		panic(err)
	}
	return f
}

func Close(f *os.File) {
	_ = f.Close()
}

func RangeOverInputFile(f func([]byte)) {
	start := time.Now()
	input := MustOpen("input")
	defer Close(input)
	reader := bufio.NewReader(input)
	var line []byte
	var err error
	for line, _, err = reader.ReadLine(); err == nil; line, _, err = reader.ReadLine() {
		f(line)
	}
	fmt.Printf("Iteration complete in %.3fµs\n", float64(time.Now().Sub(start))/1000.0)
}
