package main

import (
	"fmt"
	"time"

	"github.com/bl1nk/aoc2020/x"
)

const tree = 35 // #

func walk(right, down int, input []string) int {
	var trees, xPos int
	for yPos := down; yPos < len(input); yPos += down {
		xPos = (xPos + right) % len(input[yPos])
		if input[yPos][xPos] == tree {
			trees++
		}
	}
	return trees
}

func part1() {
	defer x.Took(time.Now())
	fmt.Println(walk(3, 1, x.InputFromPwd()))
}

func part2() {
	defer x.Took(time.Now())
	input := x.InputFromPwd()
	fmt.Println(walk(1, 1, input) * walk(3, 1, input) * walk(5, 1, input) * walk(7, 1, input) * walk(1, 2, input))
}

func main() {
	part1()
	part2()
}
