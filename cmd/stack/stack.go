package main

import (
	"algo/internal/stack"
	"fmt"
)

func main() {
	stk := &stack.Stack{}
	stk.Init()
	var err error

	var a = []int{1, 2, 3}
	for _, v := range a {
		err = stk.Push(v)
		if err != nil {
			panic(err)
		}
		fmt.Printf("push: %+v\n", v)
	}

	for stk.Length() > 0 {
		top, err := stk.Pop()
		if err != nil {
			panic(err)
		}
		fmt.Printf("pop: %+v\n", top)
	}
}
