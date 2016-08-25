// cat
package main

import (
	"fmt"
	//"io/ioutil"
	"bufio"
	"io"
	"os"
	"strconv"
)

var firstName, lastName, par1, par2, par3 string

func main() {
	/*
		fmt.Println("Pls. input your full name:")
		fmt.Scanln(&firstName, &lastName)
		fmt.Printf("Hi, %s %s\n", firstName, lastName)
	*/
	fmt.Println("Pls. input command and file name to read:")
	fmt.Scanln(&par1, &par2, &par3)
	if par1 == "cat" {
		switch par2 {
		case "-n":
			cat(par2, par3)
		default:
			cat(par2)
		}
	}

	fmt.Println("Hello World!")
	//	cat("1.txt")
}

//http://wangzhezhe.github.io/blog/2015/10/27/file-operation/
//文件读取
func cat(arg ...string) string {
	fi, err := os.Open(arg[0])
	defer fi.Close()
	switch len(arg) {
	case 1:
	case 2:
		if arg[0] != "-n" {
			return "2nd parameter is not '-n'"
		} else {
			fi, err = os.Open(arg[1])
			defer fi.Close()
		}
	default:
		return "two many parameters"
	}

	if err != nil {
		panic(err)
	}

	fw, err0 := os.Create("catoutput.txt")
	if err0 != nil {
		panic(err0)
	}
	fwbuf := bufio.NewWriter(fw)
	//	fwbuf.WriteString("fuck")
	defer fwbuf.Flush()

	linum := 0
	st := ""
	fmt.Printf("st is:%s", st)
	//fd, err := ioutil.ReadAll(fi)

	fd := bufio.NewReader(fi)

	for {
		//读出内容保存为string 每次读到以'\n'为标记的位置
		li, err := fd.ReadString('\n')
		if arg[0] == "-n" {
			st += strconv.Itoa(linum) + string(li)
		} else {
			st += string(li) + string(li)
		}
		fwbuf.WriteString(strconv.Itoa(linum) + "\t" + string(li))
		linum++
		if err == io.EOF || linum > 1e4 {
			break
		}
	}
	//st := string(fd)
	fmt.Printf("fmt.println fd is:\n%s\n", st)
	//println("println fd is:\n", st)
	return st
}

//使用bufio采用带有缓存的方式进行读写，比如通过info:=bufio.NewReader(f)将实现了
//io.Reader的接口的实例加载上来之后，就可以使用info.ReadLine（）来每次实现一整行的读取，直到err信息为io.EOF时，读取结束

/*
http://david-je.iteye.com/blog/1988940
Go代码  收藏代码
package main

import(
    "fmt"
    "os"
    "flag"
    "io"
    "io/ioutil"
    "bufio"
    "time"
)

func read1(path string)string{
    fi,err := os.Open(path)
    if err != nil{
        panic(err)
    }
    defer fi.Close()

    chunks := make([]byte,1024,1024)
    buf := make([]byte,1024)
    for{
        n,err := fi.Read(buf)
        if err != nil && err != io.EOF{panic(err)}
        if 0 ==n {break}
        chunks=append(chunks,buf[:n]...)
        // fmt.Println(string(buf[:n]))
    }
    return string(chunks)
}

func read2(path string)string{
    fi,err := os.Open(path)
    if err != nil{panic(err)}
    defer fi.Close()
    r := bufio.NewReader(fi)

    chunks := make([]byte,1024,1024)

    buf := make([]byte,1024)
    for{
        n,err := r.Read(buf)
        if err != nil && err != io.EOF{panic(err)}
        if 0 ==n {break}
        chunks=append(chunks,buf[:n]...)
        // fmt.Println(string(buf[:n]))
    }
    return string(chunks)
}

func read3(path string)string{
    fi,err := os.Open(path)
    if err != nil{panic(err)}
    defer fi.Close()
    fd,err := ioutil.ReadAll(fi)
    // fmt.Println(string(fd))
    return string(fd)
}

func main(){

    flag.Parse()
    file := flag.Arg(0)
    f,err := ioutil.ReadFile(file)
    if err != nil{
        fmt.Printf("%s\n",err)
        panic(err)
    }
    fmt.Println(string(f))
    start := time.Now()
    read1(file)
    t1 := time.Now()
    fmt.Printf("Cost time %v\n",t1.Sub(start))
    read2(file)
    t2 := time.Now()
    fmt.Printf("Cost time %v\n",t2.Sub(t1))
    read3(file)
    t3 := time.Now()
    fmt.Printf("Cost time %v\n",t3.Sub(t2))

}
*/
