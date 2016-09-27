// fetch
package main

import (
	"fmt"
	"io"
	//"io/ioutil"
	"net/http"
	"os"
	"strings"
)

func main() {
	fmt.Println("Hello World!")
	for _, url := range os.Args[1:] {
		if !strings.HasPrefix(url, "http://") {
			url = "http://" + url
		}
		fmt.Println("url: ", url)
		resp, err := http.Get(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
			os.Exit(1)
		}
		//b, err := ioutil.ReadAll(resp.Body)
		//resp.Body.Close() //avoid leaking resources
		_, err = io.Copy(os.Stdout, resp.Body)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: reading %s: %v\n", url, err)
			os.Exit(1)
		}
		//fmt.Printf("%s", b)
		os.Stdout.WriteString("status code is: " + resp.Status)
	}
}
