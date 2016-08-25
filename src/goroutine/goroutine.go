// goroutine
package main

import (
	"fmt"
	"time"
)

var c chan int

// 显式地关闭信道
//close(ch)被关闭的信道会禁止数据流入, 是只读的。我们仍然可以从关闭的信道中取出数据，但是不能再写入数据了
func ready(st string, Nsec int) {
	time.Sleep(time.Duration(Nsec) * time.Second)
	fmt.Println(st, "is ready and time is:", time.Now())
	c <- Nsec
	c <- 2
	//_ = <-c
}

func main() {
	c = make(chan int)
	fmt.Println("Hello World!")
	go ready("Tea", 2)
	go ready("Cake", 1)
	fmt.Println("I'm waiting.., time is: ", time.Now())
	<-c
	//time.Sleep(5 * time.Second)

	//<-c
	/*	i := 0
		Loop:
			for {
				select { //select 语句实现了一种监听模式
				case <-c:
					i++
					if i > 1 {
						break Loop
					}
				}
			}
	*/
	fmt.Println("done.., time is: ", time.Now())
}
