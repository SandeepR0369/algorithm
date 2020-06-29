package main

import (
	"fmt"
	"sort"
)

func main() {
	x := Missing([]int{1, 2, 45, 3, 4, 5})
	fmt.Println(x)
}

//find the first missing positive integer in linear time and constant space
//example, the input [3, 4, -1, 1] should give 2. The input [1, 2, 0] should give 3
//check what is first number and then next number isn't increment it should return increment

func TrialMissing(numbers []int) {
	var j []int
	sort.Ints(numbers)
	for i, val := range numbers {

		if numbers[i] < 0 {
			continue
		} else if numbers[i] >= 0 {
			j = append(j, val)
		}

	}
	for k, v := range j {
		for m := 0; m <= len(j)-1; m++ {
			if m == len(j) {
				if j[m+1]-j[m] == 1 {
					continue
				} else {
					fmt.Println(k, v)
				}
			}

		}
	}
	fmt.Println(j)
	//return val
}


func Missing(numbers []int) int {
	sort.Ints(numbers)
	for i, val := range numbers {
		//for i = 0; i <= len(numbers)-1; i++{
		if i == 0 {
			continue
		}
		if i == len(numbers)-1 {
			fmt.Println("V1",val + 1)
			return val + 1
		} else {
			if numbers[i-1] < 0 {
				fmt.Println("V2")
				continue
			}
			if val-numbers[i-1] == 1 {
				fmt.Println("V3", val, numbers[i-1], i, val-numbers[i-1])
				continue
			} else {
				fmt.Println("V4", numbers[i-1] - val)
				return numbers[i-1] - val
			}
		}

	//	fmt.Println(val)
		//	}
	}
	return 0
}

/*
func getNo(numbers []int ) int {
	sort.Ints(numbers)
	for i,num :=range(numbers){
		if i==0{
			continue
			}
		if i==len(numbers)-1{
		return num+1
		}else{
		if numbers[i-1]<0{
			continue
		}
		if num-numbers[i-1]==1{
			continue
		}else{
			return numbers[i-1]-num
		}
	}
	}
return 0
}
*/


/****** EDO OKATI VaSTHUNDHI*********/

func TrialMissing(numbers []int) {
	var j []int
	sort.Ints(numbers)
	for i, val := range numbers {

		if numbers[i] < 0 {
			continue
		} else if numbers[i] >= 0 {
			j = append(j, val)
		}

	}
	for k, v := range j {
	
		for m := 0; m < len(j); m++ {
			if m == len(j) {
				if j[m+1]-j[m] == 1 {
					continue
				} else {
					fmt.Println(k, v)
				}
			}

		}
	}
	fmt.Println(j)
	//return val
}

/****** EDO OKATI VaSTHUNDHI*********/

func TrialMissing(numbers []int) {
	var j []int
	sort.Ints(numbers)
	for i, val := range numbers {

		if numbers[i] < 0 {
			continue
		} else if numbers[i] >= 0 {
			j = append(j, val)
		}

	}
	for k, v := range j {
	
		for m := 0; m < len(j); m++ {
			if m == len(j) {
				if j[m+1]-j[m] == 1 {
					continue
				} else if j[m+1]-j[m] > 1{
					return j[m]+1
				}
			}

	//	} /// to print next number or missing number add 1 to previous number... Try chey bey
	}
	fmt.Println(j)
	//return val
}

func TrialMissing(numbers []int)int {
	var j []int
	sort.Ints(numbers)
	for i, val := range numbers {

		if numbers[i] < 0 {
			continue
		} else if numbers[i] >= 0 {
			j = append(j, val)
		}

	}
	//fmt.Println("A1")

	for _, _ = range j {

		for m := 0; m <= len(j)-1; m++ {
		//fmt.Println("A2", v)
			//if m != len(j) {
				if j[m+1]-j[m] == 1 {
					//fmt.Println("A3", j[m+1],j[m])
					continue
				} else if j[m+1]-j[m] > 1 {
					//fmt.Println("A4", j[m+1],j[m] )
					return j[m] + 1
				}

				//fmt.Println(k,v)

			//}
		}
		//fmt.Println("J1", j)
		//return val
	}
	return 0
}