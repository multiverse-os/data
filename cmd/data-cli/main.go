package main

import (
	"fmt"

	data "github.com/multiverse-os/data"
)

func main() {
	fmt.Println("data-cli")
	fmt.Println("A quick test for data module")

	fmt.Printf("Is true IsTrue()? %v\n", data.IsTrue("true"))
}
