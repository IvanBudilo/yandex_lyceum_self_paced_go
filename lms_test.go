package main

import (
	"testing"
)

type Test struct {
	in       []byte
	outLen   int
	outError error
}

var tests = []Test{
	{[]byte("Hello"), 5, nil},
	{[]byte{}, 0, nil},
	{[]byte{0x80}, 0, ErrInvalidUTF8},
	{[]byte("你"), 1, nil},
}

func TestGetUTFLength(t *testing.T) {
	for i, test := range tests {
		length := test.outLen
		err := test.outError
		getLen, getErr := GetUTFLength(test.in)
		if getLen != length && getErr != err {
			t.Fatalf("#%d GetUTFLength(%q)=%d, %e, want: %d, %e", i, test.in, getLen, getErr, test.outLen, test.outError)
		}
	}
}
