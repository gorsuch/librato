librato
=======

wip - lightweight librato reporting client in go

## Goals

* easy to reason about
* use mutexes, don't abuse channels
* lean on `net/http` to do the right thing

## Usage

```go
package main

import (
	"log"

	"github.com/gorsuch/librato"
)

func main() {
	c := librato.Client{
		User:  "michael.gorsuch@gmail.com",
		Token: "REDACTED",
	}

	g := librato.Gauge{
		Name:   "test",
		Source: "home",
		Count:  1,
		Sum:    200,
	}

	c.AddGauge(g)
	err := c.Flush()
	if err != nil {
		log.Fatal(err)
	}
}
```

## TODO

* [ ] documentation of funcs
* [x] support for measure_time, min, max and sum_squares
