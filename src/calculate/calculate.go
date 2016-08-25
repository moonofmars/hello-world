// calculate
package main

import (
	"flag"
	"fmt"
	"os"
	//"strings"
	//"simplemath"
	//"strconv"
)

var Usage = func() {
	fmt.Println("hello moto")
}

func main() {
	flag.Parse()
	//arg := flag.Args[0]
	fmt.Println("------ Args start ------")
	for i, v := range os.Args {
		fmt.Printf("arg[%d] = (%s).\n", i, v)
	}
	fmt.Println("Hello World! ", flag.Args())
}
