package main

import (
	"fmt"

	"leeif.me/learngo/retriever/mork"
	"leeif.me/learngo/retriever/real"
)

type Retriever interface {
	Get(url string) string
}

type Poster interface {
	Post(url string, from map[string]string) string
}

const url = "http://www.leeif.me"

func downlard(r Retriever) string {
	return r.Get(url)
}

func post(p Poster) {
	p.Post(url,
		map[string]string{
			"name":   "leeifme",
			"course": "golang",
		})
}

type RetrieverPoster interface {
	Retriever
	Poster
}

func session(s RetrieverPoster) string {
	s.Post(url, map[string]string{
		"contents": "fake leeif.me",
	})
	return s.Get(url)
}

func main() {
	var r Retriever
	r = &mork.Retriever{"This is a fake leeif.me"}
	retriever := mork.Retriever{"This is a fake leeif.me"}
	r = real.Retriever{}
	fmt.Println(downlard(r))

	fmt.Println("Try to seesion")
	fmt.Println(session(&retriever))
}
