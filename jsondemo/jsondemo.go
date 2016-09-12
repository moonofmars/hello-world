// jsondemo
package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

func main() {
	fmt.Println("Hello World!")

	dec := json.NewDecoder(os.Stdin)
	enc := json.NewEncoder(os.Stdout)

	for {
		var v map[string]interface{}
		//var v interface{}
		if err := dec.Decode(&v); err != nil {
			log.Println("decode err: ", err)
			return
		}
		fmt.Println("in: ", v)

		for k := range v {
			if k != "Name" {
				//v[k] = nil //?
				delete(v, k)
			}
		}
		fmt.Println("--\t--\t--\t--\t")
		if err := enc.Encode(&v); err != nil {
			log.Println("encode err: ", err)
		}
		fmt.Println("out: ", v)
	}
	//{"Name":"Wednesday","Age":6,"Parents":["Gomez","Morticia"]}
}
