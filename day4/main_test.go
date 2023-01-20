package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPairStringToMinMax(t *testing.T) {
	str := "2-4,8-12"
	a, b := pairStringToMinMax(str)
	assert.EqualValues(t, 2, a.min)
	assert.EqualValues(t, 4, a.max)
	assert.EqualValues(t, 8, b.min)
	assert.EqualValues(t, 12, b.max)

	str = "8-3,22-127"
	a, b = pairStringToMinMax(str)
	assert.EqualValues(t, 3, a.min)
	assert.EqualValues(t, 8, a.max)
	assert.EqualValues(t, 22, b.min)
	assert.EqualValues(t, 127, b.max)
}

func TestContainedIn(t *testing.T) {
	type tdata struct {
		str      string
		expected bool
	}

	for _, data := range []tdata{
		{"2-4,6-8", false},
		{"2-3,4-5", false},
		{"5-7,7-9", false},
		{"2-8,3-7", true},
		{"6-6,4-6", true},
		{"2-6,4-8", false},
		{"6-8,2-4", false},
	} {
		t.Run(data.str, func(t *testing.T) {
			assert.EqualValues(t, data.expected, containedIn(pairStringToMinMax(data.str)))
		})
	}
}

func TestOverlaps(t *testing.T) {
	type tdata struct {
		str      string
		expected bool
	}

	for _, data := range []tdata{
		{"5-7,7-9", true},
		{"3-8,3-7", true},
		{"6-6,4-6", true},
		{"2-6,4-8", true},
	} {
		t.Run(data.str, func(t *testing.T) {
			assert.EqualValues(t, data.expected, overlaps(pairStringToMinMax(data.str)))
		})
	}
}
