package main

import (
	"fmt"
	"time"

	"github.com/bl1nk/aoc2020/x"
)

func main() {
	defer x.Took(time.Now())

	var entries []int
	for _, line := range x.InputFromPwd() {
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
