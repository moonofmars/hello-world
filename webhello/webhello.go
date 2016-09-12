// webhello
package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

//是http.HandlerFunc类型的实例
//ResponseWriter用于包装处理HTTP服务端的响应信息,r *http.Request表示的是此次HTTP请求的一个数据结构体，即代表一个客户端
func helloHandler(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Hello, go world ~~")
}

func main() {
	fmt.Println("Hello World!")
	http.HandleFunc("/hello", helloHandler) //分发 请求
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("listen and serve: ", err.Error())
	}
	//http://127.0.0.1:8080/hello
}
