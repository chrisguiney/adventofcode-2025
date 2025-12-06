package main

import (
	"fmt"
	"testing"
)

type rotationTest struct {
	start          int
	rotation       rotation
	wantDial       int
	wantDialAtZero int
}

func TestApplyRotationPart1(t *testing.T) {
	tests := []rotationTest{
		{
			start: 11,
			rotation: rotation{
				dir:   'R',
				steps: 8,
			},
			wantDial: 19,
		},
		{
			start:    19,
			rotation: rotation{dir: 'L', steps: 19},
			wantDial: 0,
		},
		{
			start:    5,
			rotation: rotation{dir: 'L', steps: 10},
			wantDial: 95,
		},
		{
			start:    95,
			rotation: rotation{dir: 'R', steps: 5},
			wantDial: 0,
		},
	}

	for i, tt := range tests {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			if have, _ := applyRotation(tt.rotation, tt.start, 0); have != tt.wantDial {
				t.Errorf("applyRotation(%v, %v) = %v, wantDial %v", tt.start, tt.rotation, have, tt.wantDial)
			}
		})
	}
}

func TestApplyRotationPart2(t *testing.T) {
	tests := []rotationTest{
		{
			start: 50,
			rotation: rotation{
				dir:   'L',
				steps: 82,
			},
			wantDialAtZero: 1,
		},
	}
	for i, tt := range tests {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			if _, have := applyRotation(tt.rotation, tt.start, 0); have != tt.wantDialAtZero {
				t.Errorf("applyRotation(%v, %v) = %v, wantDial %v", tt.start, tt.rotation, have, tt.wantDialAtZero)
			}
		})
	}

}
