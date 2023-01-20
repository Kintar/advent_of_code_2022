package main

import (
	"bufio"
	"github.com/stretchr/testify/assert"
	"io"
	"strings"
	"testing"
)

func TestParseCrats(t *testing.T) {
	data := "[F] [W] [B] [L] [P] [D] [L]     [G]"
	expectedCrates := [9]string{
		"F", "W", "B", "L", "P", "D", "L", "", "G",
	}
	crates := parseCrates(data)
	assert.Equal(t, expectedCrates, crates)
}

func TestReadCrateLine(t *testing.T) {
	input := `                        [R] [J] [W]
            [R] [N]     [T] [T] [C]
[R]         [P] [G]     [J] [P] [T]
[Q]     [C] [M] [V]     [F] [F] [H]
[G] [P] [M] [S] [Z]     [Z] [C] [Q]
[P] [C] [P] [Q] [J] [J] [P] [H] [Z]
[C] [T] [H] [T] [H] [P] [G] [L] [V]
[F] [W] [B] [L] [P] [D] [L] [N] [G]
 1   2   3   4   5   6   7   8   9 
`
	expectedStacks := [9]string{
		"FCPGQR",
		"WTCP",
		"BHPMC",
		"LTQSMPR",
		"PHJZVGN",
		"DPJ",
		"LGPZFJTR",
		"NLHCFPTJ",
		"GVZQHTCW",
	}

	reader := bufio.NewReader(strings.NewReader(input))
	stacks := [9]string{}
	for {
		line, _, err := reader.ReadLine()
		if err == io.EOF || string(line) == " 1   2   3   4   5   6   7   8   9 " {
			break
		}
		stacks = readCrateLine(stacks, string(line))
	}

	assert.Equal(t, expectedStacks, stacks)
}

func TestParseMoveCmd(t *testing.T) {
	line := "move 14 from 4 to 1"
	cmd, ok := parseMoveCmd(line)
	assert.True(t, ok)
	expectedCmd := moveCmd{
		count: 14,
		from:  3,
		to:    0,
	}
	assert.Equal(t, expectedCmd, cmd)
}

func TestDoMoveCmd(t *testing.T) {
	stacks := [9]string{
		"abcdefg",
		"hij",
	}
	expectedStacks := [9]string{
		"abcd",
		"hijgfe",
	}
	cmd, _ := parseMoveCmd("move 3 from 1 to 2")
	stacks = doCommand(cmd, stacks)
	assert.Equal(t, expectedStacks, stacks)

	stacks = [9]string{
		"abcdefg",
		"hij",
	}
	expectedStacks = [9]string{
		"abcdefg",
		"",
		"jih",
	}
	cmd, _ = parseMoveCmd("move 3 from 2 to 3")
	stacks = doCommand(cmd, stacks)
	assert.Equal(t, expectedStacks, stacks)

}
