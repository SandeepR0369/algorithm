package main

import "fmt"

//Should return an array with count of elements greater to the right side of the current element by taking input array
//Example Input [8, 5, 11, -1, 3, 4, 2] and Output [5, 4, 4, 0, 1, 1, 0]

func main() {
	fmt.Println(RightSmallerThan([]int{8, 5, 11, -1, 3, 4, 2}))

}

func RightSmallerThan(array []int) []int {
	//var j int
	var small []int

	for i := 0; i <= len(array)-1; i++ {
		//fmt.Println("II", array[i])

		for _, v := range array {
			if v != array[i] && v < array[i] {
				//fmt.Println("V", v)
				small = append(small, v)
			}
			//small = append(small, v)
		}
		
	//	fmt.Println("SS", small)

	}

	// Write your code here.
	return nil
}
