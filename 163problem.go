package main

import (
	"fmt"
)
//1) The expression is given as a list of numbers and operands. For example: [5, 3, '+'] should return 5 + 3 = 8.
//2) For example, [15, 7, 1, 1, '+', '-', '/', 3, '*', 2, 1, 1, '+', '+', '-'] should return 5, 
//since it is equivalent to ((15 / (7 - (1 + 1))) * 3) - (2 + (1 + 1)) = 5.
//3) You can assume the given expression is always valid.