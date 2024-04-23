package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	// Create a scanner for reading from standard input
	scanner := bufio.NewScanner(os.Stdin)

	// Initialize fireData as a nested map
	fireData := make(map[string]map[string]map[string]int)

	// Read the first dataset
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			break
		}
		fields := strings.Split(line, ",")
		if len(fields) != 4 {
			fmt.Fprintf(os.Stderr, "Error: Invalid line format in first dataset: %s\n", line)
			fmt.Println("False")
			return
		}
		year, state, month, numberStr := fields[0], fields[1], fields[2], fields[3]
		number, err := strconv.Atoi(numberStr)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error: Invalid number format in first dataset: %s\n", numberStr)
			fmt.Println("False")
			return
		}

		// Create a new year map if it doesn't exist
		if fireData[year] == nil {
			fireData[year] = make(map[string]map[string]int)
		}
		// Create a new state map if it doesn't exist
		if fireData[year][state] == nil {
			fireData[year][state] = make(map[string]int)
		}
		// Store the number of fires for the month and state
		fireData[year][state][month] = number
	}

	// Read the second dataset
	var totalFires map[string]int // year -> total number of fires
	totalFires = make(map[string]int)
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			break
		}
		fields := strings.Split(line, ",")
		if len(fields) != 2 {
			fmt.Fprintf(os.Stderr, "Error: Invalid line format in second dataset: %s\n", line)
			fmt.Println("False")
			return
		}
		year, numberStr := fields[0], fields[1]
		number, err := strconv.Atoi(numberStr)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error: Invalid number format in second dataset: %s\n", numberStr)
			fmt.Println("False")
			return
		}
		totalFires[year] = number
	}

	fmt.Println("FireData", fireData)
	fmt.Println("Total Fires", totalFires)

	// Validate the first dataset using the second dataset
	for year, stateData := range fireData {
		// Initialize total number of fires reported for the year
		totalReported := 0

		// Sum up the reported numbers for all states and months within the same year
		for _, monthData := range stateData {
			for _, reportedNumber := range monthData {
				totalReported += reportedNumber
			}
		}

		// Check if the total reported fires for the year exceed the total fires for that year
		totalNumber, ok := totalFires[year]
		if !ok {
			fmt.Fprintf(os.Stderr, "Error: Year %s not found in second dataset\n", year)
			fmt.Println("False")
			return
		}
		if totalReported > totalNumber {
			fmt.Println("False")
			return
		}
	}

	// If the validation passes, print "True"
	fmt.Println("True")

}
