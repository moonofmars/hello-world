// channel
package main

import (
	"fmt"
	"time"
)

var quit chan int // 只开一个信道

func foo(id int) {
	fmt.Println("func thread is activated:\t", id)
	quit <- 0 // ok, finished. 如果quit无缓冲，且无处接收，线程就卡在此处，下一行不执行
	fmt.Println("func done:", id)
}

func main() {
	count := 5
	quit = make(chan int, 3) // 无缓冲

	for i := 0; i < count; i++ {
		go foo(i)
	}
	//<-quit
	time.Sleep(5 * time.Second)
	//for i := 0; i < count; i++ {
	//<-quit
	//}
}

/*无缓冲，实例
package main
import "fmt"

func main() {
	fmt.Println("28th fibonacci is", Fibc(28))

}

// subf processes two input channels in a wrapped goroutine
// and returns resulting new channels for subsequent fibonaccis
func subf(c1, c2 chan int) (chan int, chan int) {
	o1, o2 := make(chan int), make(chan int)
	go func(c1, c2 chan int) {
		a, b := <-c1, <-c2
		o1 <- b
		o2 <- a + b
	}(c1, c2)
	return o1, o2
}

// Fibc returns the nth fibonacci concurrently
func Fibc(n int) int {
	if n == 0 {
		return 0
	}

	a := make(chan int)
	b := make(chan int)

	// first calculation
	c, d := subf(a, b)

	// boom
	a <- 0
	b <- 1

	// find desired nth fib
	for i := 1; i < n; i++ {
		c, d = subf(c, d)
	}
	return <-c
}
//更简洁
package main

import (
    "fmt"
)

func fibonacci(n int, c chan int) {
    x, y := 1, 1
    for i := 0; i < n; i++ {
        c <- x
        x, y = y, x+y
    }
    close(c)
}

func main() {
    c := make(chan int, 10)
    go fibonacci(cap(c), c)
    for i := range c {
        fmt.Println(i)
    }
}

*/
