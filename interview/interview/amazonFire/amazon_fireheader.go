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

	// Initialize totalFires to store the total number of fires for each year
	totalFires := make(map[string]int)

	// Initialize a flag to identify the current dataset
	var isFirstDataset bool = true

	// Read the datasets
	for scanner.Scan() {
		line := scanner.Text()

		// Check for an empty line to switch datasets
		if strings.TrimSpace(line) == "" {
			isFirstDataset = false
			continue
		}

		//fmt.Println("Processing line:", line) // Debugging line

		// Skip the header lines
		if strings.HasPrefix(line, "year, state, month, number") || strings.HasPrefix(line, "year, number") {
			continue
		}

		// Process the first dataset
		if isFirstDataset {
			fields := strings.Split(line, ",")
			if len(fields) != 4 {
				fmt.Fprintf(os.Stderr, "Error: Invalid line format in first dataset: %s\n", line)
				fmt.Println("False")
				return
			}
			year, state, month, numberStr := fields[0], fields[1], fields[2], fields[3]
			number, err := strconv.Atoi(strings.TrimSpace(numberStr)) // Trim spaces
			if err != nil {
				fmt.Fprintf(os.Stderr, "Error: Invalid number format in first dataset: %s\n", numberStr)
				fmt.Println("False")
				return
			}
			if fireData[year] == nil {
				fireData[year] = make(map[string]map[string]int)
			}
			if fireData[year][state] == nil {
				fireData[year][state] = make(map[string]int)
			}
			fireData[year][state][month] = number
		} else { // Process the second dataset
			fields := strings.Split(line, ",")
			if len(fields) != 2 {
				fmt.Fprintf(os.Stderr, "Error: Invalid line format in second dataset: %s\n", line)
				fmt.Println("False")
				return
			}
			year, numberStr := fields[0], fields[1]
			number, err := strconv.Atoi(strings.TrimSpace(numberStr)) // Trim spaces
			if err != nil {
				fmt.Fprintf(os.Stderr, "Error: Invalid number format in second dataset: %s\n", numberStr)
				fmt.Println("False")
				return
			}
			// Store the total number of fires for the year
			totalFires[year] = number
		}
	}

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
		if totalReported > totalNumber || totalReported < totalNumber {
			fmt.Println("False")
			return
		}
	}

	// If the validation passes, print "True"
	fmt.Println("True")
}
\

fkr