package freq

import (
	"io/ioutil"
	"sort"
	"strings"
	"sync"
)

func Fast(file string) []WordFreq {
	buff, err := ioutil.ReadFile(file)
	if err != nil {
		panic(err)
	}

	counters := make(map[string]int, 15e3)
	counters2 := make(map[string]int, 15e3)

	// 1. compute split index
	splitIndex := len(buff) / 2
	for {
		if isAsciiLetter(buff[splitIndex]) {
			splitIndex++
		} else {
			break
		}
	}

	// 2. start helpful thread
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		countWords(buff[splitIndex:], counters2)
		wg.Done()
	}()

	/* 3. main computation */
	countWords(buff[:splitIndex], counters)
	/* /MAIN THREAD */

	ret := make([]WordFreq, 0, len(counters) * 2)
	wg.Wait()
	for word, count := range counters {
		ret = append(ret, WordFreq{word, count + counters2[word] })
	}
	for word, count := range counters2 {
		if _, found := counters[word]; !found {
			ret = append(ret, WordFreq{word, count})
		}
	}

	sort.Slice(ret, func(i, j int) bool {
		if ret[i].Frequency == ret[j].Frequency {
			return strings.Compare(ret[i].Word, ret[j].Word) == -1
		}

		return ret[i].Frequency > ret[j].Frequency
	})

	return ret
}

func countWords(buff []byte, counters map[string]int) {
	var s = newStringsBuilder(len(buff))
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
}

func isAsciiLetter(b byte) bool {
	return ('a' <= b && b <= 'z') || ('A' <= b && b <= 'Z')
}