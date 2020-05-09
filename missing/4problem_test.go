package main

import "testing"

func TestMissing(t *testing.T) {
	tests := []struct {
		in  []int
		out int
	}{
		{[]int{-1, -2, 0, 1, 2, 7, 4, 3, 5}, 6},
		{[]int{124, 10, 98, 121, 222, 970, 124, 323, 545}, 11},
		{[]int{17, 20, -2, 31, 52, 67, 24, 63, 05}, 6},
		{[]int{9, 0, 2, 1, 2, 7, 4, 3, 5}, 6},
		{[]int{-11, 0, -22, 41, 22, 27, 14, 53, 75}, 1},
		{[]int{3, 0, 3, 1, 4}, 2},
		{[]int{1, 2, 0}, 3},
		{[]int{12, 81, 28, 32, 59}, 13},
		{[]int{7, 24, 2, 42, 97}, 3},
	}
	for _, test := range tests {
		if output := Missing(test.in); output != test.out {
			t.Error("Test Failed: {} inputted,  expected, received: ", test.in, test.out, output)
		}
	}
}
