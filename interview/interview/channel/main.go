// You can edit this code!
// Click here and start typing.
package main

import (
	"fmt"
	"sync"
)

func main() {
	n := 30
	var wg sync.WaitGroup

	a := make(chan int)
	b := make(chan int)
	wg.Add(1)
	go func() {
		defer wg.Done()
		for i := 0; i <= n; i++ {

			if i%2 == 0 {
				a <- i
			} else {
				b <- i
			}
		}
		close(a)
		close(b)
	}()
	wg.Add(2)
	go func() {
		for v := range a {
			fmt.Println("Even", v)
		}
		wg.Done()
	}()

	go func() {
		for u := range b {
			fmt.Println("Odd", u)
		}
		wg.Done()
	}()

	wg.Wait()
}

/*
func number(n int) {
	a := make(chan int)
	b := make(chan int)

	for i := 0; i <= n; i++ {

		if i%2 == 0 {
			a <- i
		} else {
			b <- i
		}
	}

	close(a)
	close(b)

}
*/
