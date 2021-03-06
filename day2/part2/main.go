package main

import (
	"fmt"
	"strings"
	"time"

	"github.com/bl1nk/aoc2020/x"
)

func main() {
	defer x.Took(time.Now())
	fmt.Println(solve(x.ReadLines()))
}

type line struct {
	policy   policy
	password string
}

type policy struct {
	first  int
	second int
	char   string
}

func (l *line) valid() bool {
	firstMatch := string(l.password[l.policy.first-1]) == l.policy.char
	secondMatch := string(l.password[l.policy.second-1]) == l.policy.char
	return firstMatch != secondMatch
}

func readLine(s string) line {
	a := strings.Split(s, ": ")
	b := strings.Split(a[0], " ")
	c := strings.Split(b[0], "-")

	return line{
		policy: policy{
			first:  x.MustAtoi(c[0]),
			second: x.MustAtoi(c[1]),
			char:   b[1],
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
