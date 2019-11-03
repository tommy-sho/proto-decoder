package main

import (
	"errors"
	"math/bits"
)

type Lexer struct {
	data         []byte
	position     int
	readPosition int
}

func NewLexer(in []byte) *Lexer {
	l := &Lexer{data: in, position: 0}
	l.readPosition = l.position + 1
	return l
}

func (l *Lexer) hasNext() bool {
	return l.readPosition < len(l.data)
}

func (l *Lexer) next() {
	l.position++
	l.readPosition = l.position + 1
}

func (l *Lexer) skip(n int) {
	l.position = l.position + n
	l.readPosition = l.position + 1
}

func (l *Lexer) readByte() byte {
	b := l.data[l.position]
	l.next()
	return b
}

func (l *Lexer) readBytes(n int) ([]byte, error) {
	if l.position+n > len(l.data) {
		return nil, errors.New("out of range")
	}

	res := l.data[l.position : l.position+n]
	l.skip(n)

	return res, nil
}

/*decoder for Varint*/
func (l *Lexer) decodeVarint() (uint64, error) {
	if len(l.data) == l.position {
		return 0, errors.New("got EOF")
	}

	var bs []byte
	// read a bytes from data
	b := l.readByte()
	// if first bit of bytes is 1, append bytes until first byte is 0 = last bytes.
	for bits.LeadingZeros8(b) == 0 {
		bs = append(bs, b&0x7f) // logical AND with 0x7f
		b = l.readByte()
	}

	x := uint64(b)
	for i := 0; i < len(bs); i++ {
		x = x<<7 + uint64(bs[len(bs)-1-i])
	}

	return x, nil
}
