package main

import "testing"

func TestPerfect(t *testing.T) {
	tests := []struct {
		in  int
		out int
	}{
		{126, 1261},
		{1242, 12421},
		{16, 163},
		{24, 244},
		{53, 532},
		{3031, 30313},
	}
	for _, test := range tests {
		if output := Perfect(test.in); output != test.out {
			t.Error("Test Failed: {} inputted,  expected, received: ", test.in, test.out, output)
		}
	}
}

func TestSliceJoin(t *testing.T) {
	tests := []struct {
		in  []int
		out int
	}{
		{[]int{0, 1, 2, 7, 4, 3, 5}, 0127435},
		{[]int{124, 10, 98, 121, 2225}, 12410981212225},
		{[]int{17, 20, 31, 52, 67, 24, 63, 5}, 172031526724635},
		{[]int{9, 0, 2, 1, 2, 7, 4, 3, 5}, 902127435},
		{[]int{41, 22, 27, 14, 53, 75}, 412227145375},
		{[]int{3, 0, 3, 1, 4}, 30314},
		{[]int{12, 81, 28, 32, 59}, 1281283259},
		{[]int{7, 24, 2, 42, 97}, 72424297},
	}
	for _, test := range tests {
		if output := SliceJoin(test.in); output != test.out {
			t.Error("Test Failed: {} inputted,  expected, received: ", test.in, test.out, output)
		}
	}
}
