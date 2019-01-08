package a

import (
	"fmt"

	"bitbucket.org/bigwhite/f"
)

func CallA() {
	fmt.Println("call A: master branch")
	fmt.Println("   --> call F:")
	fmt.Printf("\t")
	f.CallF()
	fmt.Println("   --> call F end")
}