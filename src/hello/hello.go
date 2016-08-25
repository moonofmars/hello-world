package main

import "fmt" // 我们需要使用fmt包中的Println()函数
import "os"

func main() {
	fmt.Println("Hello, world. 你好，世界！\n")
	fmt.Printf("Hello, world\n")
	myArray := [10]int{1, 2, 3, 4, 5}
	mySlice := make([]int, 5, 10)
	fmt.Println(myArray)
	fmt.Printf("TODAY IS %+v\n", myArray)
	fmt.Println("Slice IS ", mySlice)
	fmt.Printf("Slice len IS %d\n", len(mySlice))
	fmt.Printf("Slice cap IS %d\n", cap(mySlice))
	i := 1
	switch i {
	case 0:
		fmt.Printf("\n0\n")
		fmt.Printf("000\n")
	case 1:
		fmt.Printf("1\n")
		fmt.Printf("111\n")
	}

	switch {
	case (0 < i) && (i < 3):
		fmt.Printf("0-3\n")
	}
JLoop:
	for j := 0; j < 5; j++ {
		for i := 0; i < 10; i++ {
			if i > 5 {
				break JLoop
			}
			fmt.Println(i)
		}
	}
	//JLoop:
	for k := 0; k < 10; k++ {
		fmt.Printf("%d", k)
	}
	fmt.Printf("\n")
	myfunc(1, 2, 3, 4, 5, 6)
	//var hex byte = 'c'
	println('c' - 'a' + 10)
	fmt.Println(unhex('c'))

	var array [6]int
	slice := array[0:5]
	sli := slice[:]
	sli2 := slice
	sli[1] = 499
	println(len(slice))
	println(&array)
	println(slice)
	println(sli)
	println(sli2)
	println(sli[1])
	println(slice[1])
	println(sli2[1])
	println(cap(slice))

	println(append(sli, 2, 3, 6, 99, 108, 18, 12, 1, 2, 12, 3))
	println(append(sli, slice...))
}

func myfunc(args ...int) {
	for _, arg := range args {
		fmt.Println(arg)
	}

	if err := os.Chmod("", 0664); err != nil {
		fmt.Println(err)
	}
}

func unhex(c byte) byte {
	switch {
	case '0' <= c && c <= '9':
		return c - '0'
	case 'a' <= c && c <= 'f':
		return c - 'a' + 10
	case 'A' <= c && c <= 'F':
		return c - 'A' + 10
	}
	return 0
}
