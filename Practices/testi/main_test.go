package main

import "testing"

func TestAdd(t *testing.T) {

	tests := []struct {
		a   int
		b   int
		out int
	}{
		{2, 3, 5},
		{2, 7, 9},
		{8, 3, 11},
	}

	for _, test := range tests {
		if output := add(test.a, test.b); output != test.out {
			t.Error("Test Failed: {} inputted,  expected, received: ", test.a, test.b, test.out, output)
		}
	}
}
