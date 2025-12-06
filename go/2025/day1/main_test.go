package main

import "testing"

func Test_dial_rotate(t *testing.T) {
	tests := []struct {
		name string
		d    dial
		s    string
		want int
	}{
		{
			name: "1-1",
			d:    dial{1},
			s:    "L1",
			want: 0,
		},
		{
			name: "1-2",
			d:    dial{1},
			s:    "L2",
			want: 99,
		},
		{
			name: "50-68",
			d:    dial{50},
			s:    "L68",
			want: 82,
		},
		{
			name: "50-68",
			d:    dial{52},
			s:    "R48",
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.d.rotate(tt.s)
			if got != tt.want {
				t.Errorf("Got %d, want %d", got, tt.want)
			}
		})
	}
}
