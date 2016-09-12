// gosort
package main

import (
	"fmt"
)

//var ch = make(chan int, 4)

func sort(ch chan int, slic []int, base int, sli2 []int) {
	lessnum := 0
	repeat := 0
	//	fmt.Print("base: ", base, "  ")

	for i := 0; i < len(slic); i++ {
		switch {
		case i == base: //自己不用比
		case slic[i] < slic[base]:
			lessnum++
		case slic[i] == slic[base]:
			repeat++
		}
	}
	//	fmt.Print("less:", lessnum, "  ")
	//	fmt.Println("repeat:", repeat)

	sli2[lessnum] = slic[base]
	for k := 1; k <= repeat; k++ {
		sli2[lessnum+k] = slic[base]
	}
	ch <- 1
}

func main() {
	ch := make(chan int)

	sli := []int{3, 12, 2, 8, 7, 99, 7, 4}
	//sli2 := sli //sli的指针赋给了sli2，修改sli2即会修改sli
	//le := len(sli)
	sli2 := make([]int, len(sli))

	for base, _ := range sli {
		//fmt.Println("base:", base)
		go sort(ch, sli, base, sli2)
	}
	for i, _ := range sli {
		if sli[i] == sli2[i] {
			fmt.Println("sli=sli2", i)
		}
		<-ch
	}

	fmt.Println("slic 1: ", sli)
	fmt.Println("slic 2: ", sli2)

	fmt.Println("Hello World!")
}
