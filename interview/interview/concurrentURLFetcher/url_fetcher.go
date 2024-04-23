// Concurrent URL Fetcher:
// Question: Design a program in Golang that fetches multiple URLs concurrently. The program should accept a list of URLs as input and asynchronously fetch their contents. Once all URLs are fetched, it should print their contents.
// Sample Input:
// urls := []string{"https://example.com", "https://example.org", "https://example.net"}
// Expected Output:
// Contents of https://example.com
// Contents of https://example.org
// Contents of https://example.net
package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"sync"
	"time"
)

func main() {

	urls := []string{"https://www.google.com", "https://www.yahoo.com", "https://www.eenadu.net"}
	var wg sync.WaitGroup
	wg.Add(len(urls))

	dataCh := make(chan interface{})
	errCh := make(chan error)

	for _, url := range urls {
		go func(url string) {
			defer wg.Done()
			data, err := urlFetcher(url)
			if err != nil {
				errCh <- err
				return
			}
			dataCh <- data

		}(url)
	}

	// Use a Goroutine to close channels after all fetch operations are done
	go func() {
		wg.Wait()
		close(dataCh)
		close(errCh)
	}()

	// Collect and handle data and errors
	for data := range dataCh {
		fmt.Println("Data:", data)
	}
	for err := range errCh {
		fmt.Println("Error:", err)
	}
}

func urlFetcher(url string) (interface{}, error) {
	fmt.Println("Incoming URL", url)
	var data interface{}
	// Create an HTTP client with a timeout
	client := &http.Client{
		Timeout: 10 * time.Second, // Adjust the timeout duration as needed
	}
	fmt.Println("STEP1", url)
	// Make the HTTP request
	resp, err := client.Get(url)
	fmt.Println("STEP2", url)
	if err != nil {
		return nil, fmt.Errorf("unable to fetch the data from URL %s: %v", url, err)
	}
	fmt.Println("STEP3", url)
	defer resp.Body.Close()

	// Check the status code for errors
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("non-OK status code received from URL %s: %s", url, resp.Status)
	}
	fmt.Println("STEP4", url)
	body, err1 := ioutil.ReadAll(resp.Body)
	if err1 != nil {
		fmt.Println("Unable to read the body")
		return nil, err
	}
	fmt.Println("STEP5", url)
	fmt.Println("STEP5A", string(body))
	// Unmarshal JSON data with timeout
	timeout := time.After(5 * time.Second) // Adjust the timeout duration as needed
	unmarshalDone := make(chan bool)
	go func() {
		select {
		case <-timeout:
			err = fmt.Errorf("STEP5B: unmarshaling JSON timed out")
		case <-unmarshalDone:
		}
	}()
	err = json.Unmarshal(body, &data)
	unmarshalDone <- true

	if err != nil {
		return nil, fmt.Errorf("unable to unmarshal JSON: %v", err)
	}
	fmt.Println("STEP6", url)
	fmt.Println("Data", data)
	return data, nil
}
