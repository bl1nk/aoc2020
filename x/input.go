package x

import (
	"bufio"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

func ReadString() string {
	s, err := ioutil.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}

	return strings.TrimSpace(string(s))
}
func ReadLines() []string {
	f, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}

	var lines []string
	s := bufio.NewScanner(f)
	for s.Scan() {
		lines = append(lines, s.Text())
	}

	return lines
}

func MustAtoi(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return i
}
