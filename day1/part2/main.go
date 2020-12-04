package main

import (
	"fmt"
	"time"

	"github.com/bl1nk/aoc2020/x"
)

func main() {
	defer x.Took(time.Now())
	lines := x.InputFromPwd()

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
