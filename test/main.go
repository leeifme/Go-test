package main

import "fmt"

// People People
type People struct{}

// ShowA ShowA
func (p *People) ShowA() {
	fmt.Println("showA")
	p.ShowB()
}

// ShowB ShowB
func (p *People) ShowB() {
	fmt.Println("showB")
}

// Teacher Teacher
type Teacher struct {
	People
}

// ShowB Teacher ShowB
func (t *Teacher) ShowB() {
	fmt.Println("teacher showB")
}

func main() {
	t := Teacher{}
	t.ShowA()
}
