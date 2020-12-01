package x

import (
	"fmt"
	"time"
)

func Took(start time.Time) {
	fmt.Printf("Took %s\n", time.Since(start))
}
