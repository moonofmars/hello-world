// simplehttp
package main

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"net"
	"net/http"
	"net/url"
	"os"
)

func main() {
	fmt.Println("Hello World!")
	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, "usage: %s host:port", os.Args[0])
		os.Exit(1)
	}
	serv := os.Args[1]

	//	conn, err := net.Dial("tcp", serv)
	//	checkError(err)

	tcpAddr, err := net.ResolveTCPAddr("tcp4", serv)
	checkError(err)
	ad, err := net.LookupHost("www.baidu.com")
	fmt.Println("tcpaddr", tcpAddr, ad)
	conn, err := net.DialTCP("tcp", nil, tcpAddr)
	checkError(err)

	_, err = conn.Write([]byte("HEAD / HTTP/1.0\r\n\r\n"))
	checkError(err)

	reslt, err := readFully(conn)
	checkError(err)

	fmt.Println(string(reslt))
	fmt.Println("done")

	resp, _ := http.Get("http://example.com/")
	defer resp.Body.Close()
	//io.Copy(os.Stdout, resp.Body)
	//fmt.Fprintln(os.Stdout, "哈哈", resp.Body)//打印出指针&（0x***形式）
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Printf("哈哈%v", string(body))

	resp, err = http.PostForm("http://example.com/posts", url.Values{"title": {
		"article tile national day"}, "content": {"article body 7 days"}})
	//fmt.Println("error is: ", err, "\nresp is: ", resp)
	//.............................................................
	resp, _ = http.Head("http://qq.com:80")
	//fmt.Println("\nhead is: ", resp)
	//.............................................................
	client := &http.Client{}
	reqest, err := http.NewRequest("GET", "http://www.baidu.com", nil) //建立一个请求
	if err != nil {
		fmt.Println("Fatal error ", err.Error())
		os.Exit(0)
	}
	//Add 头协议
	reqest.Header.Add("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,*/*;q=0.8")
	reqest.Header.Add("Accept-Language", "ja,zh-CN;q=0.8,zh;q=0.6")
	reqest.Header.Add("Connection", "keep-alive")
	reqest.Header.Add("Cookie", "设置cookie")
	//reqest.Header.Add("User-Agent", "Gobook Custom User-Agent") //
	reqest.Header.Add("User-Agent", "Mozilla/5.0 (Windows NT 6.1; WOW64; rv:12.0) Gecko/20100101 Firefox/12.0")
	response, err := client.Do(reqest) //提交
	defer response.Body.Close()
	cookies := response.Cookies() //遍历cookies
	for _, cookie := range cookies {
		fmt.Println("cookie:", cookie)
	}

	body, err1 := ioutil.ReadAll(response.Body)
	if err1 != nil {
		// handle error
	}
	fmt.Println(string(body)) //网页源码
	//.............................................................
	client = &http.Client{
	//	CheckRedirect: redirectPolicyFunc,
	}
	os.Exit(0)
}

func checkError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "fatal error: %s", err.Error())
		os.Exit(1)
	}
}

func readFully(conn net.Conn) ([]byte, error) {
	defer conn.Close()
	fmt.Println("\t readFully")
	result := bytes.NewBuffer(nil)
	n := 0
	var buf [512]byte
	for {
		n, err := conn.Read(buf[0:])
		result.Write(buf[0:n])
		if err != nil {
			if err == io.EOF {
				break
			}
			return nil, err
		}
	}
	fmt.Println("\t readFully", n)
	return result.Bytes(), nil
}
