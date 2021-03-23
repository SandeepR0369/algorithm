//Write code to check a String is palindrome or not?

//Read more: https://javarevisited.blogspot.com/2011/06/top-programming-interview-questions.html#ixzz6DNvlhOv6

package main

import "fmt"

var ac string

func main() {
	fmt.Println("hello jaffa")
	a := 9
	b := 17
	c := add(a, b)
	fmt.Println(c)
}

//var s  = "malayalam"

func add(a int, b int) int {
	//c := fmt.Println("Sum of Two numbers:", a+b)
	return a + b
}

func stringcount(ac string) string {
	//ac := []string{}
	res1 := strings.Split(strconv.Itoa(ac), "")
	for k,v := range res1{
		fmt.Println(k,v)
	}

}
