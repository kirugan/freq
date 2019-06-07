package freq

import (
	"bufio"
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

func BenchmarkNaive(b *testing.B) {
	for i := 0; i < b.N; i++ {
		r := strings.NewReader(`------+++++======`)
		sc := bufio.NewScanner(r)
		sc.Buffer(make([]byte, 0, 2), 10)
		sc.Split(scanWords)

		var words []string
		for sc.Scan() {
			words = append(words, sc.Text())
		}

		if sc.Err() != nil {
			b.Fatal()
		}
	}
}

func TestScanWords(t *testing.T) {
	cases := []struct{
		Input  string
		Expect []string
	}{
		{"a+B+c++", []string{"a", "b", "c"}},
		{"------+++++======", nil},
		{"abc+def", []string{"abc", "def"}},
		{"AAA=aaa", []string{"aaa", "aaa"}},
		{"         ", nil},
	}

	for _, cse := range cases {
		r := strings.NewReader(cse.Input)
		sc := bufio.NewScanner(r)
		sc.Buffer(make([]byte, 0, 2), 10)
		sc.Split(scanWords)

		var words []string
		for sc.Scan() {
			words = append(words, sc.Text())
		}

		assert.Equal(t, cse.Expect, words)
		assert.NoError(t, sc.Err())
	}
}
