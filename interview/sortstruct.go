package main

import (
	"fmt"
	"sort"
)

type Address struct {
	firstname string //`json:"FirstName"`
	lastname  string //`json:"LastName"`
	id        int    //`json: id`
}

type Addresses []Address

func (add Addresses) Less(i, j int) bool {
	//fmt.Println("A", i, j, add[i], add[j])
	if add[i].id < add[j].id {
		return true
	}

	return false
}

func (add Addresses) Len() int {
	return len(add)
}

func (add Addresses) Swap(i, j int) {
	//fmt.Println("B", i, j)
	add[i].id, add[j].id = add[j].id, add[i].id
}

func main() {
	//fmt.Println("Hello, playground")

	var a []Address

	a = []Address{
		Address{firstname: "Sanjay", lastname: "Sahu", id: 21},
		Address{firstname: "Banjay", lastname: "Aahu", id: 2},
		Address{firstname: "Janjay", lastname: "Bahu", id: 310},
		Address{firstname: "Sanjay", lastname: "Dahu", id: 40},
		Address{firstname: "Banjay", lastname: "Jahu", id: 59},
	}

	sort.Sort((Addresses)(a))

	fmt.Println(a)

	//fmt.Println(sort(a.firstname))

}

