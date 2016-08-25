// calc
package main

import (
	//"flag"
	"fmt"
	"strconv"
)

const (
	_      = 1000 * iota //长量计数器iota在const关键字出现时将被重置为0,iota只能在const内部使用
	ADD                  //1000
	SUB                  //2000
	MUL                  //3000
	DIV                  //4000
	MAXPOS = 11
)

var mop = map[int]string{ADD: "+", SUB: "-", MUL: "*", DIV: "/"}

var (
	ok    bool
	value int
)

type Stack struct {
	i    int
	data [MAXPOS]int
}

//构建了一个堆栈
func (s *Stack) Reset()     { s.i = 0 }
func (s *Stack) Len() int   { return s.i }
func (s *Stack) Push(k int) { s.data[s.i] = k; s.i++ }
func (s *Stack) Pop() int   { s.i--; return s.data[s.i] }

var found int
var stak = new(Stack)

func main() {

	//flag.Parse() //对命令行参数的解析。
	list := []int{1, 6, ADD, SUB, MUL, DIV}
	magic := 7
	/*
		magic, ok := strconv.Atoi(flag.Arg(0)) //magic：和
		if ok != nil {
			return
		}
	*/
	f := make([]int, MAXPOS)
	solve(f, list, 0, magic) //算法函数
	fmt.Println("Hello World!")
}
func solve(form, numberrop []int, index, magic int) {
	//	fmt.Println("indx:\t", index, "form:\n", form, "num:\n", numberrop)
	var tmp int
	for i, v := range numberrop {
		//		fmt.Println("i:\t", i, "v:\t", v)
		if v == 0 {
			goto NEXT
		}
		if v < ADD { //是一个数字，保存起来
			tmp = numberrop[i]
			numberrop[i] = 0
		}
		form[index] = v
		value, ok = rpncalc(form[0 : index+1]) //算法是？

		if ok && value == magic {
			if v < ADD {
				numberrop[i] = tmp //重置并继续
			}
			found++
			fmt.Printf("%s = %d #%d\n", rpnstr(form[0:index+1]), value, found) //rpncalc, rpnstr
		}
		if index == MAXPOS-1 {
			if v < ADD {
				numberrop[i] = tmp //重置并继续
			}
			goto NEXT
		}
		solve(form, numberrop, index+1, magic)
		if v < ADD {
			numberrop[i] = tmp //重置并继续
		}
	NEXT:
	}
}

func rpnstr(r []int) (ret string) { //将rpb转换到固定的标记
	s := make([]string, 0) //分配内存
	for k, t := range r {
		switch t {
		case ADD, SUB, MUL, DIV:
			a, s := s[len(s)-1], s[:len(s)-1]
			b, s := s[len(s)-1], s[:len(s)-1]
			if k == len(r)-1 {
				s = append(s, b+mop[t]+a)
			} else {
				s = append(s, "("+b+mop[t]+a+")")
			}
		default:
			s = append(s, strconv.Itoa(t))
		}
	}
	for _, v := range s {
		ret += v
	}
	return
}

func rpncalc(r []int) (int, bool) {
	stak.Reset()
	for _, t := range r {
		fmt.Println("t:\t", t)
		switch t {
		case ADD, SUB, MUL, DIV:
			if stak.Len() < 2 {
				return 0, false
			}
			a := stak.Pop()
			b := stak.Pop()
			if t == ADD {
				stak.Push(b + a)
			}
			if t == SUB {
				//不接受负数
				if b-a < 0 {
					return 0, false
				}
				stak.Push(b - a)
			}
			if t == MUL {
				stak.Push(b * a)
			}
			if t == DIV {
				if a == 0 {
					return 0, false
				}
				if b%a != 0 {
					return 0, false
				}
				stak.Push(b / a)
			}
		default:
			stak.Push(t)
		}
	}
	if stak.Len() == 1 { //只有一个，最后计算结果
		return stak.Pop(), true
	}
	return 0, false
}
