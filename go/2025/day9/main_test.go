package main

import "testing"

func Test_area(t *testing.T) {
	tests := []struct {
		name string
		a    coord
		b    coord
		want int
	}{
		{
			name: "example 1",
			a:    coord{2, 5},
			b:    coord{9, 7},
			want: 24,
		},
		{
			name: "example 2",
			a:    coord{7, 1},
			b:    coord{11, 7},
			want: 35,
		},
		{
			name: "example 3",
			a:    coord{7, 3},
			b:    coord{2, 3},
			want: 6,
		},
		{
			name: "example 0",
			a:    coord{0, 0},
			b:    coord{0, 0},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := area(tt.a, tt.b)
			if got != tt.want {
				t.Errorf("area() = %v, want %v", got, tt.want)
			}
		})
	}
}
