package main

import (
	"fmt"
	"freq"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		panic("wrong usage: program <input file>")
	}

	words := freq.Fast(os.Args[1])
	for _, word := range words {
		fmt.Fprintln(os.Stdout, word.Frequency, " ", word.Word)
	}
}
