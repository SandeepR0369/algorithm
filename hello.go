/*package main

import "testing"
//import "fmt"

func TestAdd(t *testing.T) {
	c := Add(13,9)
	if c != 22 {
		t.Errorf("Sum was incorrect, got: %d, want: %d.", c, 22)
	 }
}
*/
package main

import (
	"io"
	"net/http"
)

func hello(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Hello world!")
}

func main() {
	http.HandleFunc("/", hello)
	http.ListenAndServe(":8000", nil)
}

2+3, 2-3, 
