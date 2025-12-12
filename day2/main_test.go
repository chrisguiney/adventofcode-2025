package main

import (
	"fmt"
	"testing"
)

type invalidIDTest struct {
	id   int
	want bool
}

func TestIsInvalidID(t *testing.T) {
	tests := []invalidIDTest{
		{id: 1, want: false},
		{id: 1212, want: true},
		{id: 55, want: true},
		{id: 6464, want: true},
		{id: 123123, want: true},
		{id: 101, want: false},
		{id: 11, want: true},
		{id: 22, want: true},
		{id: 99, want: true},
		{id: 111, want: true},
		{id: 999, want: true},
		{id: 1010, want: true},
		{id: 222222, want: true},
		{id: 38593859, want: true},
		{id: 565656, want: true},
		{id: 824824824, want: true},
		{id: 2121212121, want: true},
	}

	for i, test := range tests {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			if have := isInvalidID(test.id); have != test.want {
				t.Errorf("isInvalidID(%d) = %t, want %t", test.id, have, test.want)
			} else {
				t.Logf("isInvalidID(%d) = %t", test.id, have)
			}
		})
	}
}
