package main

import (
	"fmt"

	humanize "github.com/dustin/go-humanize"
)

func main() {

	//8929fe90cee4b2cb9deb468b51fb34eba64d1bf0
	//https://github.com/dustin/go-humanize/commit/9f541cc9db5d55bce703bd99987c9d5cb8eea45e#diff-a9aad05aaa37fa1ee6fa0e2c53a40480

	fmt.Println("Hellow World!")
	fmt.Printf("That file is %s.\n", humanize.Bytes(82854982)) // That file is 83 MB.
	fmt.Printf("You owe $%s.\n", humanize.Comma(6582491))      // You owe $6,582,491.
}
