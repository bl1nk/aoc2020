package x

import (
	"fmt"
	"io/ioutil"
	"os"
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
