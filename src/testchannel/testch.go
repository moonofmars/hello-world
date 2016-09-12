// testch
package main

import (
	"fmt"
)

func main() {
	ch := make(chan int)
	for i := 0; i < 100; i++ {
		select {
		case ch <- 0:
		case ch <- 1:
		default:
		}
		fmt.Println("waiting: ")
		k := <-ch
		fmt.Println("k: ", k)
	}
	fmt.Println("Hello World!")
}
