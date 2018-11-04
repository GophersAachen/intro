package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	// open OMIT
	f, err := os.Open("Göthe-Faust.txt")
	if err != nil {
		panic(err)
	}

	defer f.Close()
	// open OMIT

	// scan OMIT
	sc := bufio.NewScanner(f)
	sc.Split(bufio.ScanWords)

	words := make(map[string]int)
	for sc.Scan() {
		word := sc.Text()
		word = strings.TrimRight(word, ":")
		word = strings.ToLower(word)

		if len(word) < 6 {
			continue
		}
		words[word]++
	}
	// scan OMIT

	// print OMIT
	word, count := "", 0
	for w, c := range words {
		if c > count {
			word = w
			count = c
		}
	}

	fmt.Printf("most used word: %q (%d times)\n", word, words[word])
	// print OMIT
}
