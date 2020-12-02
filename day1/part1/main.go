package main

import (
	"flag"
	"fmt"
	"time"

	"github.com/bl1nk/aoc2020/x"
)

var input = flag.String("input", "", "Puzzle input")

func main() {
	defer x.Took(time.Now())
	flag.Parse()
	lines := x.ReadInput(*input)

	var entries []int
	for _, line := range lines {
		entries = append(entries, x.MustAtoi(line))
	}

	fmt.Println(solve(entries))
}

func solve(entries []int) int {
	for i := range entries {
		for j := i + 1; j < len(entries); j++ {
			if 2020-entries[i] == entries[j] {
				return entries[i] * entries[j]
			}
		}
	}
	return -1
}
