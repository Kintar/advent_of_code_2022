package main

import "github.com/kintar/aoc2022/util"

func main() {
	loaded := false
	stacks := [9][]byte{}
	util.RangeOverInputFile(func(d []byte) {
		str := string(d)
		if !loaded {
			if str == " 1   2   3   4   5   6   7   8   9 " {
				loaded = true
				return
			}

			return
		}
	})
}
