package main

import (
	"azsy.cn/thomas-fan/go-demo/interface/retrieiver/mock"
	"azsy.cn/thomas-fan/go-demo/interface/retrieiver/real"
	"fmt"
	"time"
)

type Retriever interface {
	Get(url string) string
}

func download(r Retriever) string {
	return r.Get("http://www.baidu.com")
}
func main() {
	var r Retriever
	r = mock.Retriever{"this is mock data"}
	inspect(r)
	r = &real.Retriever{
		UserAgent: "iPhone",
		TimeOut:   time.Minute,
	}
	inspect(r)
	if realRetriever,ok := r.(*real.Retriever); ok {
		fmt.Println(realRetriever.TimeOut)
	} else {
		fmt.Println("not real retriever")
	}
}

func inspect(r Retriever) {
	fmt.Printf("%T %v\n", r, r)
	switch v := r.(type) {
	case mock.Retriever:
		fmt.Println("Contents: ", v.Contents)
	case *real.Retriever:
		fmt.Println("UserAgent: ", v.UserAgent)
	}
}
