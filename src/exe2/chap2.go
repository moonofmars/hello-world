// chap2
package main

import (
	"fmt"
	//	"strings"
	"strconv"
)

func main() {
	fmt.Println("Hello World!")
	sli := []int{2, 3, 6}
	fmt.Printf("%T \n", mapp(intdeal, sli))
	fmt.Printf("mappstr is :%s \n", mappstr(strdeal, "abcde"))
	fmt.Printf("febo is %d\n", febo(7))
	stint := []int{2, 1, 4, 8, 1}
	sort(stint)
	fmt.Printf("sorted is %d\n", stint)

	p := plusTwo()
	fmt.Printf("plus two is %v\n", p(2))
	px := plusX(5)
	fmt.Printf("plus X is %v\n", px(2))

	var s stack
	s.push(23)
	s.push(498)

	fmt.Printf("stack content is: %v\n", s.stro())
	println("stack pop 1 =  ", s.pop())
	fmt.Printf("stack content is: %v\n", s.stro())
}

//栈
type stack struct {
	i    int
	data [10]int
}

func (st *stack) push(k int) {
	if st.i > 9 {
		println("stack full")
		return
	}
	st.data[st.i] = k
	st.i++
	println("stack length is: ", st.i)
}

func (st *stack) pop() int {
	st.i--
	return st.data[st.i]
}

func (s stack) stro() string {
	var str string
	//println("stack length is: ", s.i)
	for i := 0; i < s.i; i++ {
		str = str + "[" + strconv.Itoa(i) + ":" + strconv.Itoa(s.data[i]) + "]"
	}
	return str
}

//
func plusX(x int) func(int) int {
	return func(y int) int {
		return x + y
	}
}

func plusTwo() func(int) int {
	return func(x int) int {
		return x + 2
	}
}

func sort(x []int) {
	for i := 0; i < len(x)-1; i++ {
		for j := i + 1; j < len(x); j++ {
			if x[j] < x[i] {
				x[i], x[j] = x[j], x[i]
			}
		}
	}
}

func mapp(f func(int) int, slicin []int) (slicout []int) {
	sum := 0
	slicout = slicin
	for i, val := range slicin {
		slicout[i] = f(val)
		sum += val
	}
	println("sum is: ", sum)
	return slicout
}

func mappstr(f func(rune) rune, stri string) (stro string) {
	//var sum rune
	//	var tem byte
	slicout := []rune(stri)
	for i, val := range stri {
		fmt.Printf("ele is %c \n", val)
		slicout[i] = f(val) //val是rune型，如果slicout是byte型，就不能赋值  byte=rune is wrong
		//sum += val //sum += val (mismatched types int and rune)
	}
	//println("sum is: ", sum)
	//println("slic out:", slicout)
	return string(slicout)
}

func intdeal(x int) int {
	return x * 2
}

func strdeal(x rune) rune {
	return x + 1
}

//to calculate Feborachi sequence
//用append可以避免slice初始化时长度不定，用i访问越界.或者初始化时x := make([]int , value)
func febo(x int) (sliout []int) {
	sli := []int{1, 1}
	for i := 0; i < x-2; i++ {
		sli = append(sli, sli[i]+sli[i+1])
	}
	sliout = sli
	return sliout
}
