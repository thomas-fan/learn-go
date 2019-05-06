package main

import (
	"azsy.cn/thomas-fan/go-demo/oop/queue"
	"fmt"
	"golang.org/x/tools/container/intsets"
)

func testSparse()  {
	s := intsets.Sparse{}
	s.Insert(1)
	s.Insert(100)
	s.Insert(100)
	fmt.Println(s)
}

func main() {
	q := queue.Queue{1}
	q.Push(2)
	q.Push(3)
	fmt.Println(q.Pop())
	fmt.Println(q.Pop())
	fmt.Println(q.IsEmpty())
	fmt.Println(q.Pop())
	fmt.Println(q.IsEmpty())
	testSparse()

}
