package main

import "fmt"

type Name struct {
	Value string // tag = 1
}

func (n *Name) Unmarshal(b []byte) error {

	l := NewLexer(b)

	for l.hasNext() {
		// read first byte and determine wire and tag number
		key := uint(l.readByte())
		tag := key >> 3
		wire := int(key) & 7

		switch wire {
		case 2:
			length := int(l.readByte())
			v, err := l.readBytes(length)
			if err != nil {
				return fmt.Errorf("unexpected error[length: %d]: %w", length, err)
			}
			switch tag {
			case 0:
				return ERROR_ILLEGAL_TAG_0
			case 1:
				n.Value = string(v)
			}

		default:
			return fmt.Errorf("unexpedted wire number: %d", wire)
		}
	}

	return nil
}
