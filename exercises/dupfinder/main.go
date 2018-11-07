package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, "usage: dir\n")
		os.Exit(1)
	}

	known := make(map[string]string)

	err := filepath.Walk(os.Args[1], func(name string, fi os.FileInfo, err error) error {
		fmt.Printf("checking %v\n", name)

		if err != nil {
			return err
		}

		if !fi.Mode().IsRegular() {
			return nil
		}

		fmt.Printf("  file %v\n", name)

		// md5.Sum([]byte)
		filename, ok := known[hash]
		if ok {

		}

		return nil
	})

	if err != nil {
		panic(err)
	}
}
