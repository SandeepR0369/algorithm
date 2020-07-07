package main

import(
	
)
//Find the minimum number of coins required to make n cents.
//You can use standard American denominations, that is, 1¢, 5¢, 10¢, and 25¢.
//For example, given n = 16, return 3 since we can make it with a 10¢, a 5¢, and a 1¢.

func main(){
	x := calculatecoins(357)
	fmt.Prinln("Minimum Coins Issued", x)
}

//Need to check if arg is greater than 25 divide by 25, less than 25 or >= 10, less than 10 or >= 5

func calculatecoins(inputnumber int) int {
	coinsused := 0

	for {

		if inputnumber == 0 {
			break
		}

		for {
			if inputnumber > 25 {
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
////////////////////// ANOTHER TRIAL////////////////////////////////////////////////

/*
//func Remainderchecker(a int)int{
func firststep(a int)int{
	var z []int
	out := 0

	if a >= 25 {
		if a % 25 == 0 {
			return a / 25
		} else if a % 25 != 0 {
			return a % 25
		}
	
	} 

	if a < 25 && a >= 10 {
		if a % 10 == 0 {
			return a / 10
		} else if a % 10 != 0 {
			return a % 10
		}	
	} 

	if a < 10 && a >= 5 {
		if a % 5 == 0 {
			return a / 5
		} else if a % 5 != 0 {
			return a % 5
		}	
	} 

	if a < 5 && a >= 1 {
		if a % 1  == 0 {
			return a / 1
		} 	
	} 

	return out
}


//func divisor(out int){
func secondstep(out int){

}

//func minimumcoins(){
func thirdstep(){

}


func onlystep(a int)int{
	var divisor []int
	var remainder int
	out := 0

	if a >= 25 {
		if a % 25 == 0 {
			divisor = append(divisor, a / 25) 
		} else if a % 25 != 0 {//
			if a % 25 >= 10 {
				if (a % 25) % 10 == 0{
					divisor = append(divisor, (a / 25) / 10) 
				}else if (a % 25) % 10 != 0
				remainder =  (a % 25) % 10

			}else a % 25 < 10{

			}
		}//
	
	} 

	if a < 25 && a >= 10 {
		if a % 10 == 0 {
			return a / 10
		} else if a % 10 != 0 {
			return a % 10
		}	
	} 

	if a < 10 && a >= 5 {
		if a % 5 == 0 {
			return a / 5
		} else if a % 5 != 0 {
			return a % 5
		}	
	} 

	if a < 5 && a >= 1 {
		if a % 1  == 0 {
			return a / 1
		} 	
	} 

	return out
}
*/