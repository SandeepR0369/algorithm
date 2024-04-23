package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	fmt.Println("Hello")
	f, err := os.Open("test.csv")
	if err != nil {
		fmt.Println("CSV Open error:", err)
	}
	defer f.Close()

	r := csv.NewReader(f)
	fmt.Println("CSV New Reader :", r)

	// Read the header line
	_, err = r.Read()
	if err == io.EOF {
		fmt.Println("CSV file is empty or contains only the header line")
		return
	} else if err != nil {
		fmt.Println("CSV Read header error:", err)
		return
	}

	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Printf("Error of parsing line: %v, Error: %v", record, err)
			//return fmt.Errorf("Error of parsing line: %v, Error: %v", record, err)
		}
		fmt.Println("record length", len(record), record[0], record[1], record[2])
	}
}
