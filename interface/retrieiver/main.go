package main

import (
	"azsy.cn/thomas-fan/go-demo/interface/retrieiver/mock"
	"azsy.cn/thomas-fan/go-demo/interface/retrieiver/real"
	"fmt"
	"time"
)

const url = "http://www.baidu.com"
type Retriever interface {
	Get(url string) string
}

type Poster interface {
	Post(url string, form map[string]string) string
}

type RetrieverPoster interface {
	Retriever
	Poster
}

func download(r Retriever) string {
	return r.Get(url)
}

func post(poster Poster) {
	poster.Post(url, map[string]string{
		"name":  "thomas",
		"skill": "go, php",
	})
}

func session(s RetrieverPoster) string {
	s.Post(url, map[string]string{
		"name": "thomas",
	})
	return s.Get(url)
}

func main() {
	var r Retriever
	r = &mock.Retriever{"this is mock data"}
	inspect(r)
	r = &real.Retriever{
		UserAgent: "iPhone",
		TimeOut:   time.Minute,
	}
	inspect(r)
	if realRetriever, ok := r.(*real.Retriever); ok {
		fmt.Println(realRetriever.TimeOut)
	} else {
		fmt.Println("not real retriever")
	}
	retriever := mock.Retriever{"test"}
	fmt.Println(session(&retriever))


}

func inspect(r Retriever) {
	fmt.Printf("%T %v\n", r, r)
	switch v := r.(type) {
	case *mock.Retriever:
		fmt.Println("Contents: ", v.Contents)
	case *real.Retriever:
		fmt.Println("UserAgent: ", v.UserAgent)
	}
}
