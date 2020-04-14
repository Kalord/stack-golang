package main

import (
	"Stack/stack"
	"fmt"
)

func main() {
	var stack stack.Stack

	stack.Push("One")
	stack.Push("Two")
	stack.Push("Three")

	fmt.Println("Size of the stack:", stack.Len())

	fmt.Println(stack.Pop())
	fmt.Println(stack.Pop())
	fmt.Println(stack.Pop())
}
