package main

import (
	"fmt"
	"github.com/kintar/aoc2022/util"
	"strconv"
	"strings"
)

func main() {
	loaded := false
	stacks := [9]string{}
	util.RangeOverInputFile(func(d []byte) {
		str := string(d)
		if !loaded {
			if str == " 1   2   3   4   5   6   7   8   9 " {
				loaded = true
			} else {
				crates := parseCrates(str)
				for i, v := range crates {
					stacks[i] = stacks[i] + v
				}
			}
		} else {
			if cmd, ok := parseMoveCmd(str); ok {
				stacks = doCommand(cmd, stacks)
			}
		}
	})
	fmt.Println()
	stackHeight := 0
	for _, v := range stacks {
		if len(v) > stackHeight {
			stackHeight = len(v)
		}
	}

	for ; stackHeight > 0; stackHeight-- {
		for _, s := range stacks {
			if len(s) >= stackHeight {
				fmt.Print(string(s[stackHeight-1]))
			} else {
				fmt.Print(" ")
			}
		}
		fmt.Println()
	}
}

func doCommand(cmd moveCmd, stacks [9]string) [9]string {
	origin := stacks[cmd.from]
	dest := stacks[cmd.to]
	ol := len(origin)
	stack := origin[ol-cmd.count:]
	origin = origin[:ol-cmd.count]
	for i := len(stack) - 1; i >= 0; i-- {
		dest += string(stack[i])
	}
	stacks[cmd.from] = origin
	stacks[cmd.to] = dest
	return stacks
}

func readCrateLine(stacks [9]string, line string) [9]string {
	crates := parseCrates(line)
	for i, v := range crates {
		stacks[i] = v + stacks[i]
	}
	return stacks
}

func parseCrates(data string) [9]string {
	result := [9]string{}
	for i := 0; i < 10; i++ {
		result[i] = data[1:2]
		if len(data) < 4 {
			break
		}
		data = data[4:]
	}
	for i, v := range result {
		result[i] = strings.TrimSpace(v)
	}
	return result
}

type moveCmd struct {
	count int
	from  int
	to    int
}

func toInt(s string) int {
	v, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		panic(err)
	}
	return int(v)
}

func parseMoveCmd(line string) (moveCmd, bool) {
	cmd := moveCmd{}
	if strings.TrimSpace(line) == "" {
		return cmd, false
	}
	tokens := strings.Split(line, " ")
	cmd.count = toInt(tokens[1])
	cmd.from = toInt(tokens[3]) - 1
	cmd.to = toInt(tokens[5]) - 1
	return cmd, true
}
