package main

import "fmt"

type Person struct {
	Name *Name
	Age  *Age
}

func (p *Person) Unmarshal(b []byte) error {

	l := NewLexer(b)

	for l.hasNext() {
		key := uint(l.readByte())
		tag := key >> 3
		wire := int(key) & 7

		switch wire {
		case 2: // wire type: Length-delimited
			length := int(l.readByte())
			v, err := l.readBytes(length)
			if err != nil {
				return fmt.Errorf("unexpedted error[length: %d]: %w", length, err)
			}
			switch tag {
			case 0:
				return ERROR_ILLEGAL_TAG_0
			case 1:
				p.Name = &Name{}
				if err := p.Name.Unmarshal(v); err != nil {
					return fmt.Errorf("Name.Unmarshal error: %w", err)
				}
			case 2:
				p.Age = &Age{}
				if err := p.Age.Unmarshal(v); err != nil {
					return fmt.Errorf("Age.Unmarshal error: %w", err)
				}
			default:
				return fmt.Errorf("unexpected error")
			}
		default:
			return fmt.Errorf("unexpected wire type: %d", wire)
		}
	}

	return nil
}
