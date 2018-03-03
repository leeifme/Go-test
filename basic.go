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
	const (
		ss = "def"
		a  = 3
		b  = 4
	)
	var c int
	c = int(math.Sqrt(a*a + b*b))
	fmt.Println(ss, c)
}

func main() {
	fmt.Println("hello world!")
	variableZeroVaule()
	variableInitVaule()
	variableTypeDeduction()
	variableShorter()
	fmt.Println(aa, bb, cc)
}
