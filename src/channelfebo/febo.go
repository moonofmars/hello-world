// febo
package main

import "fmt"

func dup3(in <-chan int) (<-chan int, <-chan int, <-chan int) {
	a, b, c := make(chan int, 2), make(chan int, 2), make(chan int, 2)
	go func() {
		for {
			x := <-in
			//fmt.Println("x =:", x)
			a <- x
			fmt.Println("a =:", x)
			b <- x
			c <- x
		}
	}()
	return a, b, c
}

func fib() <-chan int {
	x := make(chan int, 2)
	a, b, out := dup3(x)
	go func() {
		x <- 2
		x <- 3
		tempa := <-a
		fmt.Println("tempa =:", tempa)
		for {
			val1 := <-a
			fmt.Println("val1 =:", val1)
			val2 := <-b
			fmt.Println("val2 =:", val2)
			x <- val1 + val2
		}
	}()
	return out
}

func main() {
	x := fib()
	for i := 0; i < 3; i++ {
		fmt.Println("i =:", i)
		fmt.Println("channel out=:", <-x)
	}
}

// See sdh33b.blogspot.com/2009/12/fibona  i-in-go.html
