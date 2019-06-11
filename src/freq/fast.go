package freq

import (
	"io/ioutil"
	"sort"
	"strings"
	"sync"
	"unsafe"
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
		if isAscii(buff[splitIndex]) {
			splitIndex++
		} else {
			break
		}
	}

	// 2. start helpful thread
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		var s = newStringsBuilder(len(buff) - splitIndex)
		for _, b := range buff[splitIndex:] {
			if 'a' <= b && b <= 'z' {
				s.WriteByte(b)
			} else if 'A' <= b && b <= 'Z' {
				s.WriteByte(b + 32)
			} else {
				if s.Len() > 0 {
					word := s.String()
					counters2[word]++

					s.Reset()
				}
			}
		}

		if s.Len() > 0 {
			counters2[s.String()]++
		}

		wg.Done()
	}()

	/* 3. main computation */
	var s = newStringsBuilder(splitIndex)
	for _, b := range buff[:splitIndex] {
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
	/* /MAIN THREAD */

	ret := make([]WordFreq, 0, len(counters))

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

type stringsBuilder struct {
	offset int
	len int
	buf []byte
}

func newStringsBuilder(max int) *stringsBuilder {
	return &stringsBuilder{buf: make([]byte, max)}
}

func (sb *stringsBuilder) WriteByte(b byte) {
	sb.buf[sb.offset + sb.len] = b
	sb.len++
}

func (sb *stringsBuilder) Len() int {
	return sb.len
}

func (sb *stringsBuilder) String() string {
	buf := sb.buf[sb.offset:sb.offset + sb.len]
	return *(*string)(unsafe.Pointer(&buf))
}

func (sb *stringsBuilder) Reset() {
	sb.offset += sb.len
	sb.len = 0
}

func isAscii(b byte) bool {
	return ('a' <= b && b <= 'z') || ('A' <= b && b <= 'Z')
}