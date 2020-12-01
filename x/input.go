package x

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

func ReadInput(filepath string) []string {
	b, err := ioutil.ReadFile(filepath)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	return strings.Split(string(b), "\n")
}

func MustAtoi(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return i
}
