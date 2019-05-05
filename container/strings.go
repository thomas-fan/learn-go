package main

import "fmt"

func main() {
	s := "中文测试"
	for i, ch:= range []rune(s) {
		fmt.Printf("(%d %c) ", i, ch)
	}
}