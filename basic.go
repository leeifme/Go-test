package main

import (
	"fmt"
	"math"
)

var (
	aa = 1
	bb = true
	cc = "def"
)

func variableZeroVaule() {
	var a int
	var s string
	fmt.Printf("%d %q\n", a, s)
}

func variableInitVaule() {
	var a, b int = 3, 6
	var s string = "abc"
	fmt.Println(a, b, s)
}

func variableTypeDeduction() {
	var a, b, c, d = 1, 2, true, "def"
	fmt.Println(a, b, c, d)
}

func variableShorter() {
	a, b, c, d := 1, 2, true, "def"
	b = 5
	fmt.Println(a, b, c, d)
}

func consts() {
	// const可以作为各种数据类型使用
	const (
		ss = "def"
		a  = 3
		b  = 4
	)
	var c int
	c = int(math.Sqrt(a*a + b*b))
	fmt.Println(ss, c)
}

func enums() {
	// 枚举
	const (
		golang = iota
		python
		java
		javascript
	)
	fmt.Println(golang, python, java, javascript)

	const (
		b = 1 << (10 * iota)
		kb
		mb
		gb
		tb
		pb
	)
	fmt.Println(b, kb, mb, gb, tb, pb)
}

func main() {
	fmt.Println("hello world!")
	variableZeroVaule()
	variableInitVaule()
	variableTypeDeduction()
	variableShorter()
	fmt.Println(aa, bb, cc)

	consts()
	enums()
}
