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
			continue
		}
		year, state, month, numberStr := fields[0], fields[1], fields[2], fields[3]
		number, err := strconv.Atoi(numberStr)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error: Invalid number format in first dataset: %s\n", numberStr)
			continue
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
			continue
		}
		year, numberStr := fields[0], fields[1]
		number, err := strconv.Atoi(numberStr)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error: Invalid number format in second dataset: %s\n", numberStr)
			continue
		}
		totalFires[year] = number
	}

	// Validate the first dataset using the second dataset
	for year, stateData := range fireData {
		for state, monthData := range stateData {
			for month, reportedNumber := range monthData {
				totalNumber, ok := totalFires[year]
				if !ok {
					fmt.Fprintf(os.Stderr, "Error: Year %s not found in second dataset\n", year)
					continue
				}
				if reportedNumber > totalNumber {
					fmt.Printf("Invalid data: Year %s, State %s, Month %s - Reported number of fires exceeds total\n", year, state, month)
				}
			}
		}
	}
}
