package main

import "testing"

func Test_double(t *testing.T) {
	tests := []struct {
		name string
		n    int
		want int
	}{
		{
			name: "0",
			n:    0,
			want: 0,
		},
		{
			name: "1",
			n:    1,
			want: 11,
		},
		{
			name: "9",
			n:    9,
			want: 99,
		},
		{
			name: "10",
			n:    10,
			want: 1010,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := double(tt.n)
			if got != tt.want {
				t.Errorf("double() = %v, want %v", got, tt.want)
			}
		})
	}
}

// Benchmark function for 'double'
func BenchmarkDouble(b *testing.B) {
	const input = 12345

	for b.Loop() {
		double(input)
	}
}

func Test_get(t *testing.T) {
	tests := []struct {
		name       string
		begin, end int
		want       int
	}{
		{
			name:  "11-22 => 33",
			begin: 11,
			end:   22,
			want:  33,
		},
		{
			name:  "95-115 => 33",
			begin: 95,
			end:   115,
			want:  99,
		},
		{
			name:  "1188511880-1188511890 => 1188511885",
			begin: 1188511880,
			end:   1188511890,
			want:  1188511885,
		},
		{
			name:  "998-1012 => 1010",
			begin: 998,
			end:   1012,
			want:  1010,
		},
		{
			name:  "565653-565659 => 0",
			begin: 565653,
			end:   565659,
			want:  0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := get(tt.begin, tt.end)
			if got != tt.want {
				t.Errorf("get() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_getFirstBase(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		n    int
		want int
	}{
		{
			name: "999 => 1010",
			n:    999,
			want: 10,
		},
		{
			name: "100 => 1010",
			n:    100,
			want: 10,
		},
		{
			name: "1 => 11",
			n:    1,
			want: 1,
		},
		{
			name: "10000 => 100100",
			n:    10000,
			want: 100,
		},
		{
			name: "99999 => 100100",
			n:    99999,
			want: 100,
		},
		{
			name: "1010 => 1010",
			n:    1010,
			want: 10,
		},
		{
			name: "1000 => 1010",
			n:    1000,
			want: 10,
		},
		{
			name: "1011 => 1111",
			n:    1011,
			want: 11,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := getFirstBase(tt.n)
			if got != tt.want {
				t.Errorf("getFirstInvalid() = %v, want %v", got, tt.want)
			}
		})
	}
}
