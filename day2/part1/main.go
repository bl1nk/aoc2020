package main

import (
	"fmt"
	"strings"
	"time"

	"github.com/bl1nk/aoc2020/x"
)

func main() {
	defer x.Took(time.Now())
	fmt.Println(solve(x.InputFromPwd()))
}

type line struct {
	policy   policy
	password string
}

type policy struct {
	min  int
	max  int
	char string
}

func (l *line) valid() bool {
	return l.policy.min <= strings.Count(l.password, l.policy.char) && strings.Count(l.password, l.policy.char) <= l.policy.max
}

func readLine(s string) line {
	a := strings.Split(s, ": ")
	b := strings.Split(a[0], " ")
	c := strings.Split(b[0], "-")

	return line{
		policy: policy{
			min:  x.MustAtoi(c[0]),
			max:  x.MustAtoi(c[1]),
			char: b[1],
		},
		password: a[1],
	}
}

func solve(input []string) int {
	var lines []line
	for _, l := range input {
		lines = append(lines, readLine(l))
	}

	var validLines int
	for _, l := range lines {
		if l.valid() {
			validLines++
		}
	}

	return validLines
}
