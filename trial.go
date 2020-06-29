package main

import "fmt"

func main(){
z:=Location([]int{1,3,5,2,6,7},2)
fmt.Println(z)
}

func Location(num []int, sol int) int{
    //var sol int
    for i, _ := range num{
		//if val[i]
        if num[i] == sol{
            return i
        } 
        
	}
	return -1
}