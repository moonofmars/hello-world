// osopen
package main

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"
	"time"
)

func main() {
	fmt.Println("Hello World!")
	//fil, err := os.Stat("opp/2.txt")
	//fmt.Println("1.txt file status:", fil.Size(), fil.IsDir(), fil.Mode(), fil.ModTime(), fil.Name(), fil.Sys(), "\terr:", err)

	buf := make([]byte, 1024)
	fmt.Println("time to start 1:", time.Now())
	f, ok := os.Open("opp/2.txt")
	fmt.Println("ok?", ok)
	fmt.Println("time to start 2:", time.Now())
	defer f.Close()
	i := 0
	for {
		n, _ := f.Read(buf)
		if n == 0 {
			break
		}
		os.Stdout.Write(buf[:n])
		fmt.Println("\ncyle:\t", i, "\n")
		i++
	}
	fmt.Println("time to start 3:", time.Now())

	cmd0 := exec.Command("flag")
	_ = cmd0.Run()
	/*
		fk, err2 := exec.LookPath("ls")
		if err2 != nil {
			fmt.Println(err2)
		}
		fmt.Println(fk) //  /bin/ls
	*/
	cmd := exec.Command("opp/flag")
	cmd.Stdin = strings.NewReader("some input")
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		log.Fatal(err)
	}
	//fmt.Printf("in all caps: %q\n", out.String())　　//in all caps: "SOME INPUT"
}
