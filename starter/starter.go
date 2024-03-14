package starter

import (
	"fmt"
	"math"
)

// https://dmoj.ca/problem/lkp18c2p1
// https://dmoj.ca/problem/lkp18c2p1/submissions/nilptr/

type FoodLines struct {
	lineQty  int
	arrivals int
	lines    map[int]int
}

func NewFoodLines(lineQty, arrivals int, peopleInLines ...int) *FoodLines {
	lines := make(map[int]int)
	for k, v := range peopleInLines {
		lines[k] = v
	}

	return &FoodLines{
		lineQty:  lineQty,
		arrivals: arrivals,
		lines:    lines,
	}
}

func (f *FoodLines) getSmallestLine() int {
	smallest := math.MaxInt
	smallestIdx := 0
	for k, v := range f.lines {
		if v < smallest {
			smallest = v
			smallestIdx = k
		}
	}
	return smallestIdx
}

func (f *FoodLines) Solve() {
	for f.arrivals > 0 {
		idx := f.getSmallestLine()
		fmt.Printf("%d\n", f.lines[idx])
		f.lines[idx]++
		f.arrivals--
	}
}
