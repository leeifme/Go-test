package main

import (
	"fmt"

	"leeif.me/learngo/retriever/mork"
	"leeif.me/learngo/retriever/real"
)

type Retriever interface {
	Get(url string) string
}

func downlard(r Retriever) string {
	return r.Get("http://www.leeif.me")
}

func main() {
	var r Retriever
	r = mork.Retriever{"This is a fake leeif.me"}
	r = real.Retriever{}
	fmt.Println(downlard(r))
}
