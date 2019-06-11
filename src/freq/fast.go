package freq

import (
	"io/ioutil"
	"sort"
	"unsafe"
)

func Fast(file string) []WordFreq {
	buff, err := ioutil.ReadFile(file)
	if err != nil {
		panic(err)
	}

	counters := make(map[string]int, 29e3)
	var s = newStringsBuilder(len(buff) + 10)
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

	//fmt.Println("center" , time.Since(start))

	ret := make([]WordFreq, 0, len(counters))
	for word, count := range counters {
		ret = append(ret, WordFreq{word, count})
	}
	sort.Slice(ret, func(i, j int) bool {
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