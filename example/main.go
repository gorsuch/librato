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
