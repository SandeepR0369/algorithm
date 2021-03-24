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

/*
type Addresses []Address

func (add Addresses) Less(i, j int) bool {

	if add[i].id < add[j].id {
		return true
	}

	return false
}

func (add Addresses) Len() int {
	return len(add)
}

func (add Addresses) Swap(i, j int) {
	add[i].id, add[j].id = add[j].id, add[i].id
}
*/
func main() {

	var a []Address

	a = []Address{
		Address{firstname: "Sanjay", lastname: "Sahu", id: 21},
		Address{firstname: "Banjay", lastname: "Aahu", id: 2},
		Address{firstname: "Janjay", lastname: "Bahu", id: 310},
		Address{firstname: "Sanjay", lastname: "Dahu", id: 40},
		Address{firstname: "Banjay", lastname: "Dahu", id: 59},
	}

	sort.Slice(a, func(i, j int) bool {
		if a[i].lastname < a[j].lastname {
			return true
		}
		if a[i].lastname > a[j].lastname {
			return false
		}
		return a[i].firstname < a[j].firstname
	})

	fmt.Println(a)
}
