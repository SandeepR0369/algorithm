package main

import "fmt"

func main() {
	/*umlist := []int{1, 4, 5, 6, 74, 23}
	fmt.Println(umlist)
	fmt.Println(umlist[2])*/
	sum([]int{1, 4, 5, 6, 74, 231})
	reverse("donga badkav")
	reverse1("nee raa")
}

/*
func NuminList(list []int, num int)bool {
	var list []int
	var num int
if
} */

//need a list of numbers
//separate the numbers in the list --based on the index of the list
//
//and add each of them
//print the total sum

func sum(sum []int) int {
	//var numlist []int
	var total int

	//index := len(numlist)
	for _, val := range sum {
		total = total + val
	}
	fmt.Println(total)
	return total

	//	fmt.Println(numlist)
}

/*
func sum(sum []int) int {
	//var numlist []int
	var total int

	//index := len(numlist)
	for _, val := range sum {
		total = total + val
	}
	fmt.Println(total)
	return total

	//	fmt.Println(numlist)
}
*/
func reverse(word string) string {
	var result string
	for i := len(word) - 1; i >= 0; i-- {
		result = result + string(word[i])
	}
	fmt.Println(result)
	return result
}

func reverse1(word string) string {
	var result string
	for i := 0; i < len(word); i++ {
		fmt.Println(i, string(word[i]))
		fmt.Println("R2", result)
		result = string(word[i]) + result
		fmt.Println("R!", result)
	}
	fmt.Println(result)
	return result
}


