// ch4q17
package main

import (
	"container/list"
	"fmt"
	"strconv"
)

func main() {
	fmt.Println("Hello World!")
	println(mappany(datadeal, "love"))
	da := []int{1, 2, 5}
	println(mappany(datadeal, da))

	intfc := mappinter(interdeal, "good")
	switch intfc.(type) {
	case string:
		println("this is string: ", intfc.(string))
	}

	//container / list example
	ls := list.New()

	for i := 0; i < 5; i++ {
		ls.PushBack(i)
		ls.PushFront(i * 3)
	}
	//ls.Init()

	for e := ls.Front(); e != nil; e = e.Next() {
		//println(e.Value)
		fmt.Print(":", e.Value)
		//fmt.Println(e.Value)
	}
}

func mappany(f func(arg interface{}) string, arg interface{}) string {
	return f(arg)
}

func mappinter(f func(arg interface{}) interface{}, arg interface{}) interface{} {
	return f(arg)
}

func interdeal(arg interface{}) interface{} {
	return arg
}

func datadeal(arg interface{}) string {
	st := ""
	//for _, val := range arg {
	//fmt.Printf("val is %s\n", arg)
	switch arg.(type) {
	case string:
		//println("string: ")
		for _, val := range arg.(string) {
			st += string(val) + string(val)
		}
	case []int:
		//fmt.Printf("int-string is %d\n", arg.([]int))
		sum := 0
		for _, val := range arg.([]int) {
			sum += val
			st += strconv.Itoa(val * 2)
		}
		//st = strconv.Itoa(sum)
		fmt.Printf("sum is %s \n", st)
	default:
		println("unknown")
	}

	//st = append(st, val.(rune))
	//st = rune(val)
	//}
	//return string(val)
	//fmt.Printf("st is %d:\n", st)
	return st
}
