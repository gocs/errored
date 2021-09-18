package errored_test

import (
	"fmt"

	"github.com/gocs/errored"
)

// ErrBad occurs when the error is random
const ErrBad = errored.New("bad error")

func doSomething() error {
	return ErrBad
}

func Example() {
	if err := doSomething(); err != nil {
		fmt.Println( err)
	}
	// Output: bad error
}
