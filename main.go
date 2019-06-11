package main

import (
	"bufio"
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
	w := bufio.NewWriter(fd)
	for _, word := range words {
		fmt.Fprintln(w, word.Frequency, " ", word.Word)
	}
	w.Flush()

	fmt.Println(time.Since(start))
}
