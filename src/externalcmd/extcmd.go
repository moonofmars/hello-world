// extcmd
package main

import (
	//"bytes"
	"fmt"
	//	"log"
	"os/exec"
	//	"strings"
)

func main() {
	fmt.Println("Hello World!")
	cmd := exec.Command("opp/flag.exe", "-n", "opp/1.txt")
	out, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(out))
	/*
		cmd.Stdin = strings.NewReader("some input")
		var out bytes.Buffer
		cmd.Stdout = &out
		fmt.Printf("initial %q\n", out.String())
		err := cmd.Run()
		if err != nil {
			fmt.Printf("error: %v\n", err)
			log.Fatal(err)
		}
		fmt.Printf("output %q\n", out.String()) //in all caps: "SOME INPUT"
	*/
}
