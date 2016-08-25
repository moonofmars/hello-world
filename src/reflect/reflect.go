// reflect
package main

import (
	"fmt"
	"reflect"
)

type Person struct {
	Name string "namestr"
	age  int
}

func Set(i interface{}) {
	switch i.(type) {
	case *Person:
		r := reflect.ValueOf(i)
		r.Elem().Field(0).SetString("Albert Einstein")
		fmt.Printf("*person is %v \n", i)
	default:
		fmt.Printf("i is %v \n", k)
	}
}

func main() {
	man := Person{"Obama", 10}
	Set(&man)
	println(man.Name)

	in := 10
	Set(in)

}

/*
func ShowTag(i interface{}) {
	t := reflect.TypeOf(i)
	v := reflect.ValueOf(i) //得到实际的值,地址
	fmt.Printf("value is %v\n", v)
	fmt.Printf("t is %v\n", t)
	fmt.Printf("t.Kind() is %v\n", t.Kind())
	fmt.Printf("tag is %v\n", t.Elem().Field(0).Tag) //使用 Elem() 得到了指针指向的值。
	fmt.Printf("field is %v\n", v.Elem().Field(0))
	switch k := i.(type) {
	default:
		fmt.Printf("i.type is %v\n%v", k, i)
	}
}

func show(i interface{}) {
	switch i.(type) {
	//case *main.Person:
	default:
		t := reflect.TypeOf(i)  //得到类型的元数据
		v := reflect.ValueOf(i) //得到实际的值
		tag := t.Elem().Field(0).Tag
		//name := v.Elem().Field(0).String()
		fmt.Printf("show tag=%v\tvalue=%d\n", tag, v)
	}
}

func main() {
	fmt.Println("Hello World!")
	var p Person
	p.name = "sunny"
	p.age = 12
	ShowTag(&p)
	//	show(p)
}
*/
