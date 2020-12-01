package main

import (
	"flag"
	"fmt"
	"strconv"

	"github.com/bl1nk/aoc2020/x"
)

var input = flag.String("input", "", "Puzzle input")

func main() {
	flag.Parse()
	lines := x.ReadInput(*input)

	for i, xs := range lines {
		x, err := strconv.Atoi(xs)
		if err != nil {
			panic(err)
		}

		for j := i + 1; j < len(lines); j++ {
			y, err := strconv.Atoi(lines[j])
			if err != nil {
				panic(err)
			}

			if x+y != 2020 {
				continue
			}

			fmt.Println(x*y)
		}
	}
}
