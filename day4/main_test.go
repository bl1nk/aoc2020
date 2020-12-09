package main

import (
	"fmt"
	"testing"
)

func TestIsInRange(t *testing.T) {
	testcases := []struct {
		value   string
		atLeast int
		atMost  int
		want    bool
	}{
		{
			value:   "2002",
			atLeast: 1920,
			atMost:  2002,
			want:    true,
		},
		{
			value:   "2003",
			atLeast: 1920,
			atMost:  2002,
			want:    false,
		},
	}

	for _, tc := range testcases {
		t.Run(tc.value, func(t *testing.T) {
			if got := isInRange(tc.value, tc.atLeast, tc.atMost); got != tc.want {
				t.Errorf("%d <= %s <= %d, got: %t, want: %t", tc.atLeast, tc.value, tc.atMost, got, tc.want)
			}
		})
	}
}

func TestIsValidHeight(t *testing.T) {
	testcases := []struct {
		value string
		want  bool
	}{
		{
			value: "60in",
			want:  true,
		},
		{
			value: "190cm",
			want:  true,
		},
		{
			value: "190in",
			want:  false,
		},
		{
			value: "190",
			want:  false,
		},
	}

	for _, tc := range testcases {
		t.Run(tc.value, func(t *testing.T) {
			if got := isValidHeight(tc.value); got != tc.want {
				t.Errorf("value: %q, got: %t, want: %t", tc.value, got, tc.want)
			}
		})
	}
}

func TestIsValidHairColor(t *testing.T) {
	testcases := []struct {
		value string
		want  bool
	}{
		{
			value: "#123abc",
			want:  true,
		},
		{
			value: "#123abz",
			want:  false,
		},
		{
			value: "123abc",
			want:  false,
		},
	}

	for _, tc := range testcases {
		t.Run(tc.value, func(t *testing.T) {
			if got := isValidHairColor(tc.value); got != tc.want {
				t.Errorf("value: %q, got: %t, want: %t", tc.value, got, tc.want)
			}
		})
	}
}

func TestIsValidEyeColor(t *testing.T) {
	testcases := []struct {
		value string
		want  bool
	}{
		{
			value: "brn",
			want:  true,
		},
		{
			value: "wat",
			want:  false,
		},
	}

	for _, tc := range testcases {
		t.Run(tc.value, func(t *testing.T) {
			if got := isValidEyeColor(tc.value); got != tc.want {
				t.Errorf("value: %q, got: %t, want: %t", tc.value, got, tc.want)
			}
		})
	}
}

func TestIsValidPassportID(t *testing.T) {
	testcases := []struct {
		value string
		want  bool
	}{
		{
			value: "000000001",
			want:  true,
		},
		{
			value: "0123456789",
			want:  false,
		},
	}

	for _, tc := range testcases {
		t.Run(tc.value, func(t *testing.T) {
			if got := isValidPassportID(tc.value); got != tc.want {
				t.Errorf("value: %q, got: %t, want: %t", tc.value, got, tc.want)
			}
		})
	}
}

func TestValidatePassport_Part1(t *testing.T) {
	testcases := []struct {
		passport string
		want     bool
	}{
		{
			passport: `ecl:gry pid:860033327 eyr:2020 hcl:#fffffd
byr:1937 iyr:2017 cid:147 hgt:183cm`,
			want: true,
		},
		{
			passport: `iyr:2013 ecl:amb cid:350 eyr:2023 pid:028048884
hcl:#cfa07d byr:1929
`,
			want: false,
		},
		{
			passport: `hcl:#ae17e1 iyr:2013
eyr:2024
ecl:brn pid:760753108 byr:1931
hgt:179cm`,
			want: true,
		},
		{
			passport: `hcl:#cfa07d eyr:2025 pid:166559648
iyr:2011 ecl:brn hgt:59in`,
			want: false,
		},
	}

	for idx, tc := range testcases {
		t.Run(fmt.Sprintf("Day 4 part 1 (%d)", idx), func(t *testing.T) {
			if got := validatePassport(tc.passport, false); got != tc.want {
				t.Errorf("passport: %q, got: %t, want: %t\n", tc.passport, got, tc.want)
			}
		})
	}
}

func TestValidatePassport_Part2(t *testing.T) {
	testcases := []struct {
		passport string
		want     bool
	}{
		{
			passport: `eyr:1972
cid:100
hcl:#18171d ecl:amb hgt:170 pid:186cm iyr:2018 byr:1926`,
			want: false,
		},
		{
			passport: `iyr:2019
hcl:#602927 eyr:1967 hgt:170cm
ecl:grn pid:012533040 byr:1946`,
			want: false,
		},
		{
			passport: `hcl:dab227 iyr:2012
ecl:brn hgt:182cm pid:021572410 eyr:2020 byr:1992 cid:277`,
			want: false,
		},
		{
			passport: `hgt:59cm ecl:zzz
eyr:2038 hcl:74454a iyr:2023
pid:3556412378 byr:2007`,
			want: false,
		},
		{
			passport: `pid:087499704 hgt:74in ecl:grn iyr:2012 eyr:2030 byr:1980
hcl:#623a2f`,
			want: true,
		},
		{
			passport: `eyr:2029 ecl:blu cid:129 byr:1989
iyr:2014 pid:896056539 hcl:#a97842 hgt:165cm`,
			want: true,
		},
		{
			passport: `hcl:#888785
hgt:164cm byr:2001 iyr:2015 cid:88
pid:545766238 ecl:hzl
eyr:2022`,
			want: true,
		},
		{
			passport: `iyr:2010 hgt:158cm hcl:#b6652a ecl:blu byr:1944 eyr:2021 pid:093154719`,
			want:     true,
		},
	}

	for idx, tc := range testcases {
		t.Run(fmt.Sprintf("Day 4 part 2 (%d)", idx), func(t *testing.T) {
			if got := validatePassport(tc.passport, true); got != tc.want {
				t.Errorf("passport: %q, got: %t, want: %t\n", tc.passport, got, tc.want)
			}
		})
	}
}
