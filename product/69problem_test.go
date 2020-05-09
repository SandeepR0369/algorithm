package main 

import "testing"

func TestLargestProduct(t *testing.T){
	 tests := []struct{
		in []int
		out int
	}{
		{[]int{-1, -2, 0, 1, 2, 7, 4, 3, 5}, 1680},
		{[]int{124, 10, 9, 9}, 11160},
		{[]int{17, 20, -2, 31, 52, 67, 24, 05}, 4406563200 },
	//	{[]int{9, 0, 2, 1, 2, 7, 4, 3, 5}, 15121},
	//	{[]int{-11, 0, -22, 41, 22, 27, 14, 75}, 6188351400 },
	//	{[]int{3, 0, 3, 1, 4}, 36},
	//	{[]int{1, 2, 0}, 2},
		{[]int{12, 81, 28, 32, 59}, 4281984},
		{[]int{7, 24, 2, 42, 97}, 684432},
	}
	for _, test := range tests{
		if output := LargestProduct(test.in); output != test.out {
			t.Error("Test Failed: {} inputted, received, expected : ", test.in, test.out, output)
		}
	}
}

