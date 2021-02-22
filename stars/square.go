package main

import (
	"fmt"
	"strings"
)
//TYPE 1
func main() {
	var rows int
	fmt.Print("Enter the rows no: ")
	fmt.Scan(&rows)
	
	for i := 1; i <= rows; i++ {
		fmt.Println(strings.Repeat("*", rows))
	}
}

//TYPE 2
/*										
func main() {
	var rows int
	fmt.Print("Enter the rows no:")
	fmt.Scan(&rows)
	
	for i := 1; i <= rows; i++ {
		fmt.Println("*")
		for m := 1; m < rows; m++ {
			fmt.Print("*")
		}
	}
}
*/
//TYPE 3
/*
func main() {
	var rows int
	fmt.Print("Enter the rows no:")
	fmt.Scan(&rows)

	for i := 1; i <= rows; i++ {
		for m := 1; m <= rows; m++ {
			fmt.Print("*")
			if m == rows {
				fmt.Println()
			}
		}
	}
}
*/