// echo
package main

import (
	"bufio"
	"fmt"
	"net"
)

func main() {
	fmt.Println("Hello World!")
	l, err := net.Listen("tcp", "127.0.0.1:8888")
	if err != nil {
		fmt.Printf("fail to listen: %s\n", err.Error()) //Error()是string类型
	}
	for {
		if c, err := l.Accept(); err == nil {
			fmt.Println("echo start")
			Echo(c)
		}
	}
}

func Echo(c net.Conn) {
	defer c.Close()
	line, err := bufio.NewReader(c).ReadString('\n')
	if err != nil {
		fmt.Printf("fail to read: %s\n", err.Error())
		return
	}
	_, err = c.Write([]byte(line))
	if err != nil {
		fmt.Printf("fail to write: %s\n", err.Error())
		return
	}
}
