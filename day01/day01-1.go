package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func main() {
	b, err := ioutil.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}

	lines := strings.Split(string(b), "\n")
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

			fmt.Printf("%d+%d=2020\n", x, y)
			fmt.Printf("%d*%d=%d\n", x, y, x*y)
		}
	}
}
