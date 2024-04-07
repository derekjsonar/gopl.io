// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 6.
//!+

// Echo2 prints its command-line arguments.
package main

import (
	"fmt"
	"os"
)

func main() {
	for idx, arg := range os.Args[1:] {
		_, err := fmt.Fprintf(os.Stdout, "idx: %d, val: %s\n", idx, arg)
		if err != nil {
			return
		}
	}
}

//!-
