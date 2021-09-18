# errored

declare constant errors

## philosophy

When an error is created, it is usually never get edited, is just string, and is expected not to change throughout the runtime.

## installation

```
go get github.com/gocs/errored
```

## how to use

```go
package main

import (
	"fmt"

	"github.com/gocs/errored"
)

// ErrBad occurs when the error is random
const ErrBad = errored.New("bad error")

func doSomething() error {
	return ErrBad
}

func main() {
	if err := doSomething(); err != nil {
		fmt.Println("error:", err)
	}
}

```

## contributing

Because of the Go2 errors changes, this package is not accepting proposals for new functionality. With that said, we welcome pull requests, bug fixes and issue reports. 

Before sending a PR, please discuss your change by raising an issue.

## license

MIT (c) 2021 gocs