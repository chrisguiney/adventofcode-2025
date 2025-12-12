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
	}

	for i, test := range tests {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			if have := isInvalidID(test.id); have != test.want {
				t.Errorf("isInvalidID(%d) = %t, want %t", test.id, have, test.want)
			}
		})
	}
}
