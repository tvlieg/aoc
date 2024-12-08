package main

import (
	"reflect"
	"testing"
)

func Test_getAntinodes(t *testing.T) {
	type args struct {
		a1 coord
		a2 coord
	}
	tests := []struct {
		name string
		args args
		want []coord
	}{
		{
			name: "horizontal first line",
			args: args{coord{1, 0}, coord{2, 0}},
			want: []coord{{0, 0}, {3, 0}},
		},
		{
			name: "vertical first column",
			args: args{coord{0, 1}, coord{0, 2}},
			want: []coord{{0, 0}, {0, 3}},
		},
		{
			name: "diagonal",
			args: args{coord{1, 1}, coord{2, 2}},
			want: []coord{{0, 0}, {3, 3}},
		},
		{
			name: "diagonal 2",
			args: args{coord{2, 1}, coord{1, 2}},
			want: []coord{{3, 0}, {0, 3}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getAntinodes(tt.args.a1, tt.args.a2); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getAntinodes() = %v, want %v", got, tt.want)
			}
		})
	}
}
