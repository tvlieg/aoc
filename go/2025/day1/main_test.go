package main

import "testing"

func Test_dial_rotate(t *testing.T) {
	tests := []struct {
		name      string
		pos       int
		cmd       string
		wantPos   int
		wantZeros int
	}{
		// Right
		{
			name:      "0 R0 => 0",
			pos:       0,
			cmd:       "R0",
			wantPos:   0,
			wantZeros: 0,
		},
		{
			name:      "0 R1 => 1",
			pos:       0,
			cmd:       "R1",
			wantPos:   1,
			wantZeros: 0,
		},
		{
			name:      "99 R1 => 0",
			pos:       99,
			cmd:       "R1",
			wantPos:   0,
			wantZeros: 1,
		},
		{
			name:      "52 R48 => 0",
			pos:       52,
			cmd:       "R48",
			wantPos:   0,
			wantZeros: 1,
		},
		{
			name:      "50 R1000 => 50",
			pos:       50,
			cmd:       "R1000",
			wantPos:   50,
			wantZeros: 10,
		},
		// Left
		{
			name:      "0 L0 => 0",
			pos:       0,
			cmd:       "L0",
			wantPos:   0,
			wantZeros: 0,
		},
		{
			name:      "0 L1 => 99",
			pos:       0,
			cmd:       "L1",
			wantPos:   99,
			wantZeros: 0,
		},
		{
			name:      "1 L1 => 0",
			pos:       1,
			cmd:       "L1",
			wantPos:   0,
			wantZeros: 1,
		},
		{
			name:      "1 L2 => 99",
			pos:       1,
			cmd:       "L2",
			wantPos:   99,
			wantZeros: 1,
		},
		{
			name:      "50 L68 => 82",
			pos:       50,
			cmd:       "L68",
			wantPos:   82,
			wantZeros: 1,
		},
		{
			name:      "1 L101=> 0",
			pos:       1,
			cmd:       "L101",
			wantPos:   0,
			wantZeros: 2,
		},
		{
			name:      "50 L1000 => 50",
			pos:       50,
			cmd:       "L1000",
			wantPos:   50,
			wantZeros: 10,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := dial{tt.pos, 100}
			gotZeros := d.rotate(tt.cmd)

			if gotZeros != tt.wantZeros {
				t.Errorf("Got zeros %d, want %d", gotZeros, tt.wantZeros)
			}

			if d.pos != tt.wantPos {
				t.Errorf("Got %d, want %d", d.pos, tt.wantPos)
			}
		})
	}
}
