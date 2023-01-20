package main

import (
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

func TestFindHalves_Panics(t *testing.T) {
	assert.Panics(t, func() { findHalves([]byte("abc")) })
}

func TestFindHalves(t *testing.T) {
	assert.EqualValues(t, findHalves([]byte("abcd")), [2]string{"ab", "cd"})
}

func TestFindAllCommon(t *testing.T) {
	assert.Equal(t, "c"[0], findAllCommon(findHalves([]byte("abccde"))))
	assert.EqualValues(t, 0, findAllCommon(findHalves([]byte("abcCde"))))
}

func TestPriorityOf(t *testing.T) {
	alphabet := "abcdefghijklmnopqrstuvwxyz"
	for i, c := range alphabet {
		assert.EqualValues(t, i+1, priorityOf(uint8(c)), string(c))
	}
	for i, c := range strings.ToUpper(alphabet) {
		assert.EqualValues(t, i+27, priorityOf(uint8(c)), string(c))
	}

}

func TestFindCommon(t *testing.T) {
	lines := [3][]byte{[]byte("vJrwpWtwJgWrhcsFMMfFFhFp"),
		[]byte("jqHRNqRjqzjGDLGLrsFMfFZSrLrFZsSL"),
		[]byte("PmmdzqPrVvPwwTWBwg"),
	}
	assert.EqualValues(t, "r", string(findCommon(lines)))
}
