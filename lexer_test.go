package main

import (
	"reflect"
	"testing"
)

func TestLexer_decodeVarint(t *testing.T) {
	type fields struct {
		data         []byte
		position     int
		readPosition int
	}
	tests := []struct {
		name    string
		fields  fields
		want    uint64
		wantErr bool
	}{
		{
			name: "2byte",
			fields: fields{
				data: []byte{
					0x81,
					0x01,
				}, position: 0, readPosition: 1},
			want:    129,
			wantErr: false,
		},
		{
			name:    "1byte",
			fields:  fields{data: []byte{0x10}, position: 0, readPosition: 1},
			want:    16,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := &Lexer{
				data:         tt.fields.data,
				position:     tt.fields.position,
				readPosition: tt.fields.readPosition,
			}
			got, err := l.decodeVarint()
			if (err != nil) != tt.wantErr {
				t.Errorf("decodeVarint() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("decodeVarint() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLexer_hasNext(t *testing.T) {
	type fields struct {
		data         []byte
		position     int
		readPosition int
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		{
			name: "success",
			fields: fields{
				data:         []byte(`test`),
				position:     0,
				readPosition: 1,
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := &Lexer{
				data:         tt.fields.data,
				position:     tt.fields.position,
				readPosition: tt.fields.readPosition,
			}
			if got := l.hasNext(); got != tt.want {
				t.Errorf("hasNext() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLexer_next(t *testing.T) {
	type fields struct {
		data         []byte
		position     int
		readPosition int
	}
	tests := []struct {
		name   string
		fields fields
	}{
		{
			name: "",
			fields: fields{
				data:         []byte(`test`),
				position:     0,
				readPosition: 1,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := &Lexer{
				data:         tt.fields.data,
				position:     tt.fields.position,
				readPosition: tt.fields.readPosition,
			}
			l.next()
			if l.position != tt.fields.position+1 {
				t.Errorf("not next step")
			}
			if l.readPosition != l.position+1 {
				t.Errorf("not next step")
			}
		})
	}
}

func TestLexer_readByte(t *testing.T) {
	type fields struct {
		data         []byte
		position     int
		readPosition int
	}
	tests := []struct {
		name   string
		fields fields
		want   byte
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := &Lexer{
				data:         tt.fields.data,
				position:     tt.fields.position,
				readPosition: tt.fields.readPosition,
			}
			if got := l.readByte(); got != tt.want {
				t.Errorf("readByte() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLexer_readBytes(t *testing.T) {
	type fields struct {
		data         []byte
		position     int
		readPosition int
	}
	type args struct {
		n int
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []byte
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := &Lexer{
				data:         tt.fields.data,
				position:     tt.fields.position,
				readPosition: tt.fields.readPosition,
			}
			got, err := l.readBytes(tt.args.n)
			if (err != nil) != tt.wantErr {
				t.Errorf("readBytes() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("readBytes() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLexer_skip(t *testing.T) {
	type fields struct {
		data         []byte
		position     int
		readPosition int
	}
	type args struct {
		n int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{
			name: "",
			fields: fields{
				data:         []byte(`test`),
				position:     0,
				readPosition: 1,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := &Lexer{
				data:         tt.fields.data,
				position:     tt.fields.position,
				readPosition: tt.fields.readPosition,
			}
			l.skip(2)
			if l.position != tt.fields.position+2 {
				t.Errorf("not next step")
			}
			if l.readPosition != l.position+1 {
				t.Errorf("not next step")
			}
		})
	}
}
