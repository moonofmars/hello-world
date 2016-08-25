// revpolish
package main

import (
	"fmt"
	"strconv"
	"unicode"
	//"unicode/utf8"
	"os"

	"./stack"
)

func main() {
	fmt.Println("Hello World!")
	var s stack.Stack
	/*	s.Push(34)
		s.Push(2)
		println("pop 1: ", s.Pop())
		fmt.Printf("pop 2: %v\n", s.Pop())
		fmt.Printf("pop 3: %v\n", s.Pop())
	*/
	revpstr := "512+4*+3-"
	//	revpsli := revpstr[:]

	for i, val := range revpstr {
		//println(utf8.DecodeRune(byte(val)))
		if unicode.IsDigit(val) {
			digi, _ := strconv.Atoi(string(val)) //绕了两圈，1-rune to string: string(val);2-string to int: Atoi
			fmt.Printf("val=%v\n", digi)
			s.Push(int(digi))
		} else {
			x := s.Pop() //top value
			y := s.Pop() //bottom value
			fmt.Printf("x=%d\ty=%d\n", x, y)
			switch val {
			case '+':
				s.Push(y + x)
			case '-':
				s.Push(y - x)
			case '*':
				s.Push(y * x)
			case '/':
				s.Push(y / x)
			default:
				fmt.Printf("wrong expression, i=%d, value=%c\n", i, val)
			}
		}
	}
	fmt.Printf("calculation result is:%d\n", s.Pop())

	f, _ := os.Open("123.txt")
	fmt.Printf("stack.go: %v\n", f.Fd())
	f1 := os.NewFile(f.Fd(), "test.go")
	fd, _ := f1.Stat()
	fmt.Printf("test.go : %v\n", fd.ModTime())
}
