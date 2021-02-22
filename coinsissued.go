package main

import (
	"fmt"
)
//Calculate minimum coins issued in USD
func main() {
	fmt.Println("Hello, playground")
	inputnumber := 75
	noofcoins := calculatecoins(inputnumber)
	fmt.Println("coins used", noofcoins)
}

func calculatecoins(inputnumber int) int {
	coinsused := 0

	for {

		if inputnumber == 0 {
			break
		}

		for {
			if inputnumber > 25 {
		//	fmt.Println("A", inputnumber)
				inputnumber = inputnumber - 25
				coinsused += 1
			} else {
				break
			}
		}

		for {
			if inputnumber > 10 {
				inputnumber = inputnumber - 10
				coinsused += 1
			} else {
				break
			}
		}

		for {
			if inputnumber > 5 {
				inputnumber = inputnumber - 5
				coinsused += 1
			} else {
				break
			}
		}

		for {
			if inputnumber >= 1 {
				inputnumber = inputnumber - 1
				coinsused += 1
			} else {
				break
			}
		}

	}
	return coinsused

}
