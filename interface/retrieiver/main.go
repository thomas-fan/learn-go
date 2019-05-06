package main

import (
	"azsy.cn/thomas-fan/go-demo/interface/retrieiver/mock"
	"azsy.cn/thomas-fan/go-demo/interface/retrieiver/real"
	"fmt"
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
	r = real.Retriever{}
	fmt.Println(download(r))
}
