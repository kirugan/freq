package main

import (
	"fmt"
	"freq"
	"os"
	"time"
)

func main() {
	start := time.Now()
	if len(os.Args) != 3 {
		panic("wrong usage: program <input file> <output file>")
	}

	fd, err := os.OpenFile(os.Args[2], os.O_CREATE | os.O_TRUNC | os.O_WRONLY, 0666)
	if err != nil {
		panic(err)
	}
	defer fd.Close()

	words := freq.Fast(os.Args[1])
	for _, word := range words {
		fmt.Fprintln(fd, word.Frequency, " ", word.Word)
	}

	fmt.Println(time.Since(start))
}
