package main

import "fmt"

type Age struct {
	Value int32 // tag = 1
}

func (a *Age) Unmarshal(b []byte) error {

	lex := NewLexer(b)

	for lex.hasNext() {
		// read first byte and determine wire and tag number
		key := uint(lex.readByte())
		tag := key >> 3      // tag number
		wire := int(key) & 7 // wire type

		switch wire {
		case 0: // Variant type
			switch tag {
			case 0:
				return ERROR_ILLEGAL_TAG_0
			case 1:
				i, err := lex.decodeVarint()
				if err != nil {
					return fmt.Errorf("decodeVarint error: %w", err)
				}
				a.Value = int32(i)
			default:
				return fmt.Errorf("unexpected tag number: %d", tag)
			}
		default:
			return fmt.Errorf("unexpected wire type: %d", wire)
		}
	}

	return nil
}
