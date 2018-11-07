package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Fprintf(os.Stderr, "usage: input-dir output-dir\n")
		os.Exit(1)
	}

	// ioutil.ReadDir()
	// for _, filename := range entries {}
	// https://godoc.org/github.com/nfnt/resize

	fmt.Printf("more code here\n")
}
