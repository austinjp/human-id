package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/austinjp/niceid"
)

func main() {
	n := flag.Int("n", 1, "number of ids to generate")
	flag.Parse()

	if *n < 1 {
		fmt.Fprintln(os.Stderr, "niceid: -n must be >= 1")
		os.Exit(2)
	}

	for i := 0; i < *n; i++ {
		fmt.Println(niceid.ID(niceid.Options{}))
	}
}
