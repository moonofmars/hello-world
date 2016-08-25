// list
package main

import (
	"errors"
	"fmt"
)

type Value int //?Declare a type for the value our list will contain;

type Node struct { //declare a type for the each node in our list;
	Value
	prev, next *Node //why pointer?
}

type List struct {
	head, tail *Node
}

//Mimic the interface of container/list.
func (ls *List) Front() *Node {
	return ls.head
}

func (n *Node) Next() *Node {
	return n.next
}

func (ls *List) Push(v Value) *List {
	n := &Node{Value: v} //When pushing, create a new Node with the provided value;
	if ls.head == nil {  //if the list is empty, put the new node at the head;
		ls.head = n
	} else {
		ls.tail.next = n //otherwise put it at the tail;
		n.prev = ls.tail //make sure the new node points back to the previously existing one;
	}
	ls.tail = n //point tail to the newly inserted node.
	return ls
}

var errEmp = errors.New("List is empty")

func (ls *List) Pop() (v Value, err error) {
	if ls.tail == nil {
		err = errEmp
	} else {
		v = ls.tail.Value //otherwise save the last value;
		ls.tail = ls.tail.prev
		if ls.tail == nil {
			ls.head = nil
		}
	}
	return v, err
}

func main() {
	fmt.Println("Hello World!")
	ls := new(List)
	ls.Push(2)
	ls.Push(1)
	ls.Push(9)

	for e := ls.Front(); e != nil; e = e.Next() {
		fmt.Printf("%v\t", e.Value)
	}
}
