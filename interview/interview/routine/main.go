package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"sync"
)

func main() {
	var urls = []string{"https://www.google.com", "https://www.yahoo.com", "https://www.eenadu.net"}
	var wg sync.WaitGroup

	for _, url := range urls {
		wg.Add(1)
		go httpGet(url, &wg)
	}
	wg.Wait()
}

func httpGet(url string, wg *sync.WaitGroup) {
	defer wg.Done()
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(resp)

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
	}
	var response interface{}
	if err := json.Unmarshal(body, &response); err != nil {
		fmt.Printf("Error decoding JSON from URL %s: %s\n", url, err)
		return
	}

	fmt.Printf("Response from %s: %+v\n", url, response)

	// json.Decoder(resp.Body).Decode(&resp)
	// var response interface{}
	// decoder := json.NewDecoder(resp.Body)
	// if err := decoder.Decode(&response); err != nil {
	// 	fmt.Printf("Error decoding JSON from URL %s: %s\n", url, err)
	// 	return
	// }

	// fmt.Printf("Response from %s: %+v\n", url, response)
}
