package main

import "fmt"

type pointerAA struct {
	a int
}

func main() {
	fmt.Println("aaa")

	a := &pointerAA{
		a: 10,
	}

	b := &pointerAA{
		a: 10,
	}

	if a == b {
		fmt.Println("equal")
	} else {
		fmt.Println("not equal")
	}
}
