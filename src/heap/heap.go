// heap
package main

import (
	"container/heap"
	"fmt"
	"sort"
)

type IntHeap []int

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *IntHeap) Push(x interface{}) {
	// Push and Pop use pointer receivers because they modify the slice's length,
	// not just its contents.
	*h = append(*h, x.(int))
	fmt.Println("*h is:%s", *h)
}

func (h *IntHeap) Pop() interface{} {
	old := *h
	fmt.Printf("old:%d\n", old)
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	fmt.Printf("x:%d\n", x)
	return x
}

//pop，push都会自动刷新
func main() {
	fmt.Println("Hello World!")
	h := &IntHeap{2, 3, 7, 88, 5, 6, 8, 9}
	heap.Init(h)
	fmt.Printf("h1: %d\n", (*h))
	heap.Push(h, 1)
	fmt.Printf("h2: %d\n", (*h))
	///*
	heap.Fix(h, 1) //  在修改第i个元素后，调用本函数修复堆，比删除第i个元素后插入新元素更有效率。复杂度O(log(n))，其中n等于h.Len()。
	fmt.Printf("h3: %d\n", (*h))
	fmt.Printf("minimum: %d\n", (*h)[0])
	i := 0
	for i < 9 {
		fmt.Printf("h[%d]:%d\t", i, (*h)[i])
		i++
	}
	fmt.Println("")
	for h.Len() > 0 {
		fmt.Printf("p:%d\t", heap.Pop(h))
	}

	fmt.Println("\n使用sort.Sort排序:")
	h2 := IntHeap{100, 16, 4, 8, 70, 2, 36, 22, 5, 12}
	sort.Sort(h2)
	for _, v := range h2 {
		fmt.Printf("%d ", v)
	}
	//*/
}
