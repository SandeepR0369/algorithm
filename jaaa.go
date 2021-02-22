package main

import (
	"fmt"
)

type producer interface  {
	send() string

}

type data struct{
	name string
} 

func main() {

	var d producer

	dat := data{
		name: "Heell",
	}

	d = data(dat)

	fmt.Println(d)

}

func (d data) send()string{
	x := "Hello World"

	return x
}