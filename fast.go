package freq

import (
	"fmt"
	"io/ioutil"
	"sort"
	"strings"
	"time"
)

func Fast(file string) []WordFreq {
	start := time.Now()
	buff, err := ioutil.ReadFile(file)
	if err != nil {
		panic(err)
	}

	fmt.Println("read", time.Since(start))

	counters := make(map[string]int, 29e3)
	var s strings.Builder
	for _, b := range buff {
		if 'a' <= b && b <= 'z' {
			s.WriteByte(b)
		} else if 'A' <= b && b <= 'Z' {
			s.WriteByte(b + 32)
		} else {
			if s.Len() > 0 {
				word := s.String()
				counters[word]++

				s.Reset()
			}
		}
	}

	if s.Len() > 0 {
		counters[s.String()]++
	}

	fmt.Println("center" , time.Since(start))

	ret := make([]WordFreq, 0, len(counters))
	for word, count := range counters {
		ret = append(ret, WordFreq{word, count})
	}
	sort.Slice(ret, func(i, j int) bool {
		return ret[i].Frequency > ret[j].Frequency
	})

	return ret
}