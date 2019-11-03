package main

import (
	"encoding/hex"
	"fmt"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func atob(s string) []byte {
	b, _ := hex.DecodeString(s)
	return b
}

func TestUnmarshalPerson(t *testing.T) {
	tests := []struct {
		name   string
		b      []byte
		expect Person
	}{
		{
			name: "success",
			b:    atob("0a070a05416c69636512020814"),
			expect: Person{
				Name: &Name{Value: "Alice"},
				Age:  &Age{Value: 20},
			},
		},
		{
			// ゼロ値
			name:   "zero",
			b:      atob(""),
			expect: Person{},
		},
		{
			// Ageのみゼロ値
			name: "age_is_zero",
			b:    atob("0a070a05416c696365"),
			expect: Person{
				Name: &Name{Value: "Alice"},
			},
		},
		{
			// Nameのみゼロ値
			name: "name_is_zero",
			b:    atob("12020814"),
			expect: Person{
				Age: &Age{Value: 20},
			},
		},
		{
			// Varintが2バイトになる場合
			name: "varint_2bytes",
			b:    atob("1203088301"),
			expect: Person{
				Age: &Age{Value: 131},
			},
		},
		{
			// Varintが3バイトになる場合
			name: "varint_3bytes",
			b:    atob("120408928002"),
			expect: Person{
				Age: &Age{Value: 32786},
			},
		},
	}

	for i, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := Person{}
			if err := p.Unmarshal(tt.b); err != nil {
				t.Fatalf("test[%d - failed to Unmarshal. got err:%q", i, err)
			}
			fmt.Printf("%#v\n", p)
			if diff := cmp.Diff(p, tt.expect); diff != "" {
				t.Fatalf("test[%v - failed to Unmarshal. expected=%v, got=%v", i, tt.expect, p)
			}
		})
	}
}
