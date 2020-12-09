package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"time"
	"unicode"

	"github.com/bl1nk/aoc2020/x"
)

func main() {
	passportProcessing()
	partTwo()
}

// passportProcessing solves Advent of Code 2020 day 4 part 1
func passportProcessing() {
	defer x.Took(time.Now())
	input := x.ReadString()
	fmt.Println(countValid(input, false))
}

func partTwo() {
	defer x.Took(time.Now())
	input := x.ReadString()
	fmt.Println(countValid(input, true))
}

func countValid(in string, validate bool) int {
	var count int

	passports := strings.Split(in, "\n\n")
	for _, passport := range passports {
		if validatePassport(passport, validate) {
			count++
		}
	}

	return count
}

func validatePassport(input string, validate bool) bool {
	currentFields := strings.Split(strings.Replace(input, " ", "\n", -1), "\n")
	var valid int
	for field, ok := range requiredFields {
		for _, f := range currentFields {
			kv := strings.Split(f, ":")
			if field == kv[0] && (!validate || ok(kv[1])) {
				valid++
			}
		}
	}
	return valid == len(requiredFields)
}

type validator func(string) bool

var requiredFields = map[string]validator{
	"byr": func(val string) bool {
		return isInRange(val, 1920, 2002)
	},
	"iyr": func(val string) bool {
		return isInRange(val, 2010, 2020)
	},
	"eyr": func(val string) bool {
		return isInRange(val, 2020, 2030)
	},
	"hgt": isValidHeight,
	"hcl": isValidHairColor,
	"ecl": isValidEyeColor,
	"pid": isValidPassportID,
}

func isInRange(val string, atLeast, atMost int) bool {
	iv, err := strconv.Atoi(val)
	if err != nil {
		return false
	}
	return atLeast <= iv && iv <= atMost
}

func isValidHeight(val string) bool {
	units := val[len(val)-2:]
	n, err := strconv.Atoi(val[:len(val)-2])
	if err != nil {
		return false
	}
	return (units == "cm" && 150 <= n && n <= 193) || (units == "in" && 59 <= n && n <= 76)
}

var validHairColor = regexp.MustCompile(`^#[0-9a-f]{6}$`)

func isValidHairColor(val string) bool {
	return validHairColor.Match([]byte(val))
}

func isValidEyeColor(val string) bool {
	switch val {
	case "amb", "blu", "brn", "gry", "grn", "hzl", "oth":
		return true
	}

	return false
}

func isValidPassportID(val string) bool {
	for _, c := range val {
		if !unicode.IsDigit(c) {
			return false
		}
	}
	return len(val) == 9
}
