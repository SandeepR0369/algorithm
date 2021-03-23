package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"reflect"
	"strconv"
)

func main() {

	fmt.Println("Hello, playground")
	//}

	//func trend() {

	f, err := os.Open("/Users/sandeepreddybongunuri/Desktop/data.csv")

	if err != nil {
		fmt.Println("cannot open")
	}

	reader := csv.NewReader(f)

	//fmt.Println(f)

	for {
		st := make(map[string]int)
		sli := make([]int, 0)
		data, err := reader.Read()

		if err == io.EOF {
			fmt.Println("End Of File")
			fmt.Println("STKS", st)
			break
		} else {
			fmt.Errorf("Cannot read data")

		}
		//var sli []int
		//	sli = append(sli, data[1:])

		//	st[data[0]] = average(sli)
		//fmt.Println("LEN", len(data))
		fmt.Println("DATA", data)
		fmt.Printf("Question: %s Answer %s\n", data[0], data[1:])

		for i := 1; i < len(data); i++ {
			//	st := make(map[string][]int)
			if s, err := strconv.Atoi(data[i]); err == nil {
				//fmt.Printf("%T, %v", s, s)
				sli = append(sli, s)
			}
			//	sli = append(sli, data[i])
			//sli = append(sli, int(data[i]))
		}
		st[data[0]] = average(sli)
		fmt.Println(reflect.TypeOf(sli))
		fmt.Println("STKS", st)
	}

	//	fmt.Println(data)

	// Making a slice for each stock
	// appending data into that slice
	// getting average of each slice

	// parsing based on date, and difference of opening and closing

	// separating and storing into slice for each stock
	// Need to get number of rows in a file using map[string]int
	/*

		avg := make(map[string]int)


		//for j := 0; j < rows; j++{
			avg["stock"] = average(j)
		//}


		fmt.Println(avg)

	*/

}

func average(data []int) int {

	sum := 0

	for i := 0; i < len(data); i++ {
		sum += data[i]
	}

	return sum / len(data)
}
