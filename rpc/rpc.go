// rpc
package main

import (
	"fmt"
	"log"
	"net"
	//	"net/http"
	"net/rpc"
	"time"
)

type Args struct {
	A, B int
}

type Quotient struct {
	Quo, Rem int
}

type Arith int

func (t *Arith) Multiply(args *Args, reply *int) error {
	*reply = args.A * args.B
	return nil
}

func (t *Arith) Divide(args *Args, quo *Quotient) error {
	quo.Quo = args.A / args.B
	quo.Rem = args.A % args.B
	return nil
}

func main() {
	fmt.Println("Hello World!")

	//arith := new(Arith)
	//rpc.Register(arith)
	//rpc.HandleHTTP()

	newServer := rpc.NewServer()
	newServer.Register(new(Arith))
	ls, _ := net.Listen("tcp", "127.0.0.1:1234") // any available address

	go newServer.Accept(ls)

	newServer.HandleHTTP("/foo", "/bar")
	time.Sleep(2 * time.Second)
	address, _ := net.ResolveTCPAddr("tcp", "127.0.0.1:1234")
	conn, _ := net.DialTCP("tcp", nil, address)
	defer conn.Close()
	client := rpc.NewClient(conn)
	defer client.Close()
	/*
		l, e := net.Listen("tcp", ":1234")
		if e != nil {
			log.Fatal("listen error:: ", e)
		}

		go http.Serve(l, nil) //?
		time.Sleep(5 * time.Second)
		//开始拨号
		client, _ := rpc.DialHTTP("tcp", "127.0.0.1"+":1234")
	*/
	args := &Args{70, 8}
	reply := 0
	//同步
	err := client.Call("Arith.Multiply", args, &reply) //
	// 我们rpc请求的时候，调用就是这个方法，传入方法名，参数，获取返回等
	//func (client *Client) Call(serviceMethod string, args interface{}, reply interface{}) error {
	// Call里面调用了client.Go，然后返回一个chan，之后阻塞等待，这是基本的同步调用
	if err != nil {
		log.Fatal("arith error: ", err)
	}
	log.Println("同步：", reply)
	//异步
	quotient := new(Quotient)
	divcall := client.Go("Arith.Divide", args, &quotient, nil)
	replyCall := <-divcall.Done
	log.Println("异步：", replyCall)

}
