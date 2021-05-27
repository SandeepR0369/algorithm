package main

import "fmt"

func main() {

	fmt.Println(Nearestpositiveindex([]int{4, 1, 3, 5, 6}, 0))
}
//https://play.golang.org/p/O-eEE-XNewS

//Given an array of numbers and an index i, return the index of the nearest larger number of the number at index i, where distance is measured in array indices.
//For example, given [4, 1, 3, 5, 6] and index 0, you should return 3.
//If two distances to larger numbers are the equal, then return any one of them. If the array at i doesn't have a nearest larger integer, then return null.

func Nearestpositiveindex(sli []int, target int) int {

	var sli2 []int
	for i := 0; i < len(sli); i++ {
		if i == target {
			sli2 = append(sli2, 0)
			continue
		}

		if sli[i]-sli[target] > 0 || sli[i]-sli[target] <= 0 {
			sli2 = append(sli2, sli[i]-sli[target])

		}

	}

	return leastpositive(sli2)

}

func leastpositive(sli3 []int) int {
	for i := 0; i < len(sli3); i++ {
		if sli3[i] > 0 {
			return i
		}
	}
	return 0
}

