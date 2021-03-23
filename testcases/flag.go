package main

import "flag"
import "fmt"

func main() {
    var routingKey string
    flag.StringVar(&routingKey, "routingKey", "demo2018", "Routing Key To Be parsed")
    flag.Parse()
    fmt.Println("Routing Key:", routingKey)
}
