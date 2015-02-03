package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"

	"github.com/kaneshin/go-hyde/hyde"
)

func main() {
	flag.Parse()
	if flag.NArg() == 0 {
		os.Exit(0)
	}

	var vals []string
	var h hyde.HydeFunc

	switch flag.Arg(0) {
	case "to", "t":
		vals = flag.Args()[1:]
		h = hyde.ToHyde
	case "from", "f":
		vals = flag.Args()[1:]
		h = hyde.FromHyde
	default:
		vals = flag.Args()
		h = hyde.ToHyde
	}

	for _, val := range vals {
		if f, err := strconv.ParseFloat(val, 0); err == nil {
			fmt.Fprintln(os.Stdout, h(f))
		}
	}
}
