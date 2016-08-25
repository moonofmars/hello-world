// golangexec
package main

import (
	"fmt"
	//	"strings"
	//"unicode/utf8"
	//	"time"
)

func main() {
	fmt.Println("Hello World!")
	//	for i := 1; i < 11; i++ {
	//		fmt.Printf("%d\n", i)
	//}
	/*	j := 0

		cyc:
			j++
			println(j)
			if j < 5 {
				time.Sleep(5 * time.Millisecond)
				goto cyc
			}
			arr := [...]int{1, 2, 4, 6}
			for k, val := range arr {
				println(k, val)
			}
	*/
	//fallthrough
	/*
		for m := 0; m < 100; m++ {
			switch {
			case m%3 == 0 && m%5 == 0:
				println(m, "FizzBuzz")
			case m%3 == 0:
				println("Fizz")
			case m%5 == 0:
				println("Buzz")
			}
		}
	*/
	/*
		cha := [...]string{"A"}
		fmt.Printf("original string is %v\n", cha)
		chas := cha[:]
		for n := 0; n < 10; n++ {
			chas = append(chas, "A")
			//chas = append(chas, "A")
			fmt.Println("added string is: ", chas)
		}
	*/
	/*
		cha2 := "A"
		for p := 0; p < 5; p++ {
			cha2 += "A"
			println(cha2)
		}
		println(len(cha2))
	*/
	//for cha; len(cha) < 10; cha++ {
	//println(len(chas))
	//}
	/*
		var str string = "abCD eF "
		slic := []byte(str)
		//str = strings.Replace(str, " ", "", -1)
		println("\nlength is", len(str))
		println("utf-8", utf8.RuneCountInString(str))
		for a, b := range str {
			//		println(a, b)
			fmt.Printf("\t%d, %d", a, b)
			if b == 32 {
				println("got blank")
			}

			if a < len(str)/2 {
				//			println("haha .......")
				//		println(slic[a])
				slic[a], slic[len(str)-1-a] = slic[len(str)-1-a], slic[a]
				//	println(slic[a])
			}
		}
		fmt.Printf("\nswitched slic is:%s ", slic)

		println("\nslic cap is", cap(slic))

		n := copy(slic[4:7], "abc")
		println(n)
		//println(slic[5])
		slic = append(slic, 'X', 'Y')
		fmt.Printf("\n ele is %c", slic[6])
		fmt.Printf("\nslice is: %s ", slic)
	*/
	floatdata := []float64{0, 2, 3.44, 2.5601}
	fmt.Printf("float slice is %g", floatdata)
	sum := float64(0)
	for _, s := range floatdata {
		sum += s
	}
	println(sum)
	//	println(sum / float64(len(floatdata)))

	//	rec(7)
	//	ini(7)

	//	for i := 0; i < 5; i++ {
	//	defer fmt.Printf("\t%d ", i)
	//}

	myfunc(1, 2, 3, 99)

}

func myfunc(arg ...interface{}) {
	for j, k := range arg {
		//		println(j, k)
		fmt.Printf("\n%d\t%d", j, k)
	}
}

func rec(i int) {
	if i == 10 {
		return
	}
	fmt.Printf("%d ", i)
	rec(i + 1)
	fmt.Printf("%d ", i)

}

func ini(i int) (k int) {
	println("initial k = ", k)
	return
}
