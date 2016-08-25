// gosort
//github
package main

import (
	"fmt"
)

func sort(ch chan int, slic []int, base int, sli2 []int) {
	lessnum := 0
	repeat := 0
	for i := 0; i < len(slic); i++ {
		switch {
		case i == base: //自己不用比
		case slic[i] < slic[base]:
			lessnum++
		case slic[i] == slic[base]:
			repeat++
		}
	}

	sli2[lessnum] = slic[base]
	for k := 1; k <= repeat; k++ {
		sli2[lessnum+k] = slic[base]
	}
	ch <- lessnum
}

func main() {
	ch := make(chan int)

	sli := []int{1, 12, 2, 8, 4, 5, 8, 19, 13, 12, 2, 8, 4, 5, 8, 19}
	sli2 := make([]int, len(sli))

	for base, _ := range sli {
		go sort(ch, sli, base, sli2)
	}
	for _, _ = range sli {
		<-ch
	}
	fmt.Println("slic 1: ", sli)
	fmt.Println("slic 2: ", sli2)
	fmt.Println("Hello World!")
}
