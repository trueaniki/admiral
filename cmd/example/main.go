package main

import (
	"fmt"

	"github.com/trueaniki/admiral"
)

func main() {
	a := admiral.New("example", "example app")
	a.AddCommand("hello", "say hello").AddCommand("world", "say world").AddFlag("name", "n", "your name")
	a.Command("hello").Command("world").AddFlag("namename", "n", "your name")
	a.Command("hello").Command("world").AddFlag("namenamename", "n", "your name")
	a.Command("hello").Command("world").AddFlag("namenamenamename", "n", "your name")
	a.Command("hello").Command("world").AddFlag("namenamenamenamename", "n", "your name")
	fmt.Println(a.Command("hello").Command("world").Help())
}
