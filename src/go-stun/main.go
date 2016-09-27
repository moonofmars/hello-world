// main
package main

import (
	"fmt"

	"./stun"
)

func main() {
	addr, err := stun.Lookup("stun:stun.l.google.com:19302", "username", "password")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(addr)
	}
}
