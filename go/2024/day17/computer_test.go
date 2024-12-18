package main

import (
	"strings"
	"testing"
)

func TestComputer_compute(t *testing.T) {
	type reg struct {
		a, b, c int
	}
	tests := []struct {
		name       string
		reg        reg
		program    []int
		wantOutput string
		wantA      int
		wantB      int
	}{
		{
			name:    "If register C contains 9, the program 2,6 would set register B to 1.",
			reg:     reg{c: 9},
			program: []int{2, 6},
			wantB:   1,
		},
		{
			name:       "If register A contains 10, the program 5,0,5,1,5,4 would output 0,1,2.",
			reg:        reg{a: 10},
			program:    []int{5, 0, 5, 1, 5, 4},
			wantOutput: "0,1,2",
		},
		{
			name:       "If register A contains 2024, the program 0,1,5,4,3,0 would output 4,2,5,6,7,7,7,7,3,1,0 and leave 0 in register A.",
			reg:        reg{a: 2024},
			program:    []int{0, 1, 5, 4, 3, 0},
			wantOutput: "4,2,5,6,7,7,7,7,3,1,0",
			wantA:      0,
		},
		{
			name:    "If register B contains 29, the program 1,7 would set register B to 26.",
			reg:     reg{b: 29},
			program: []int{1, 7},
			wantB:   26,
		},
		{
			name:    "If register B contains 2024 and register C contains 43690, the program 4,0 would set register B to 44354.",
			reg:     reg{b: 2024, c: 43690},
			program: []int{4, 0},
			wantB:   44354,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &computer{
				a: tt.reg.a,
				b: tt.reg.b,
				c: tt.reg.c,
			}
			c.compute(tt.program)

			gotOutput := strings.Join(c.stdout, ",")
			if tt.wantOutput != gotOutput {
				t.Errorf("Expected output %v, got %v", tt.wantOutput, gotOutput)
			}

			if tt.wantA != 0 && tt.wantA != c.a {
				t.Errorf("Expected register A to be %v, got %v", tt.wantA, c.a)
			}
			if tt.wantB != 0 && tt.wantB != c.b {
				t.Errorf("Expected register A to be %v, got %v", tt.wantB, c.b)
			}
		})
	}
}
