package main

import (
	"fmt"
	"strings"
	"time"

	"github.com/bl1nk/aoc2020/x"
)

func main() {
	passportProcessing()
}

// passportProcessing solves Advent of Code 2020 day 4 part 1
func passportProcessing() {
	defer x.Took(time.Now())
	input := x.ReadString()
	fmt.Println(countValid(input))
}

func countValid(in string) int {
	var count int

	passports := strings.Split(strings.Replace(in, " ", "\n", -1), "\n\n")
	for _, passport := range passports {
		if validatePassport(passport) {
			count++
		}
	}

	return count
}

func validatePassport(input string) bool {
	requiredFields := []string{"byr", "iyr", "eyr", "hgt", "hcl", "ecl", "pid"} // cid (Country ID) not required
	currentFields := strings.Split(input, "\n")
	var valid int
	for _, rf := range requiredFields {
		for _, f := range currentFields {
			kv := strings.Split(f, ":")
			if rf == kv[0] {
				valid++
			}
		}
	}
	return valid == len(requiredFields)
}
