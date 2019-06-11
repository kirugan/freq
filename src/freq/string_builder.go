package freq

import "unsafe"

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
