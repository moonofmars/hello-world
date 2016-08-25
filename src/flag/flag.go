package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
)

var Input_flagvar int

func Init() {
	// 设置flag参数 (变量指针，参数名，默认值，帮助信息)
	flag.IntVar(&Input_flagvar, "flagname", 1234, "help message for flagname")
}

var numberFlag = flag.Bool("n", false, "number each line")

func cat(r *bufio.Reader) {
	i := 1
	for {
		buf, e := r.ReadBytes('\n')
		if *numberFlag {
			fmt.Fprintf(os.Stdout, "%5d %s", i, buf)
			i++
		} else {
			fmt.Fprintf(os.Stdout, "%s", buf)
		}
		if e == io.EOF {
			fmt.Fprintf(os.Stdout, "\n")
			//fmt.Println(os.Stdout, "\n")
			break
		}
	}
	return
}

func main() {
	println("start ...")
	Init()
	fmt.Fprintf(os.Stdout, "%d", flag.NArg())
	flag.Parse()
	fmt.Fprintf(os.Stdout, "%d", flag.NArg())
	if flag.NArg() == 0 {
		cat(bufio.NewReader(os.Stdin)) //如果命令行没带参数，等待输入。输入A，输出A
	}
	for i := 0; i < flag.NArg(); i++ {
		f, e := os.Open(flag.Arg(i)) //打开falg.Arg(i)文件
		if e != nil {
			fmt.Fprintf(os.Stderr, "%s: error reading from %s: %s\n", os.Args[0], flag.Arg(i), e.Error())
			continue
		}
		cat(bufio.NewReader(f))
	}
}
