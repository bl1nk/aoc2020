package main

import (
	"flag"
	"fmt"

	"github.com/bl1nk/aoc2020/x"
)

var input = flag.String("input", "", "Puzzle input")

func main() {
	flag.Parse()
	lines := x.ReadInput(*input)

	for i, xs := range lines {
		a := x.MustAtoi(xs)
		for j := i + 1; j < len(lines); j++ {
			b := x.MustAtoi(lines[j])
			for k := j + 1; k < len(lines); k++ {
				c := x.MustAtoi(lines[k])
				if a+b+c != 2020 {
					continue
				}
				fmt.Println(a * b * c)
			}
		}
	}
}
