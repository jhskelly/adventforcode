package main

import "testing"

func Test_total(t *testing.T) {

	tests := []struct {
		name string
		args []int
		want int
	}{
		// TODO: Add test cases.
		{name: "Some text", args: []int{}, want: 0},
		{name: "Some text", args: []int{1}, want: 1},
		{name: "Some text", args: []int{1, 3}, want: 4},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := total(tt.args); got != tt.want {
				t.Errorf("total() = %v, want %v", got, tt.want)
			}
		})
	}
}
