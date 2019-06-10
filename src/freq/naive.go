package freq

import (
	"bufio"
	"os"
	"sort"
)

func Naive(file string) []WordFreq {
	fd, err := os.Open(file)
	if err != nil {
		panic(err)
	}

	counters := make(map[string]int)

	sc := bufio.NewScanner(fd)
	sc.Split(scanWords)

	for sc.Scan() {
		word := sc.Text()
		counters[word]++
	}

	if err := sc.Err(); err != nil {
		panic(err)
	}


	ret := make([]WordFreq, 0, len(counters))
	for word, count := range counters {
		ret = append(ret, WordFreq{word, count})
	}
	sort.Slice(ret, func(i, j int) bool {
		return ret[i].Frequency > ret[j].Frequency
	})

	return ret
}

func scanWords(data []byte, atEOF bool) (advance int, token []byte, err error) {
	var breakOnLetter bool
	for _, b := range data {
		if 'a' <= b && b <= 'z' {
			if breakOnLetter {
				break
			}

			token = append(token, b)
		} else if 'A' <= b && b <= 'Z' {
			if breakOnLetter {
				break
			}

			// make lowercase
			token = append(token, b + 32)
		} else {
			// skip one char and break
			advance++
			breakOnLetter = true
		}
	}

	// request more data
	if !atEOF && len(data) == len(token) {
		return advance, nil, nil
	}


	advance += len(token)
	return
}