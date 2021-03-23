package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"sync"
	"time"
)

var wg sync.WaitGroup

func main() {

	wg.Add(2)
	fmt.Println("Hello, playground")

	go ResponseSize("https://www.google.com/")
	go ResponseSize("https://www.yahoo.com/")

	time.Sleep(10 * time.Second)

	wg.Wait()

}

func ResponseSize(url string) {

	defer wg.Done()

	fmt.Println("Step1", url)

	resp, err := http.Get(url)

	if err != nil {
		fmt.Println("ERR1", err)
	}

	fmt.Println("Step2", url)

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		fmt.Println("ERR2", err)
	}

	fmt.Println(body)

}
