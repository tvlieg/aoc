package main

import "testing"

func TestConcat(t *testing.T) {
	type args struct {
		a int
		b int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "2x 1 digit",
			args: args{1, 2},
			want: 12,
		},
		{
			name: "2 and 3 digits",
			args: args{12, 345},
			want: 12345,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := concat(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("concat() = %v, want %v", got, tt.want)
			}
		})
	}
}
